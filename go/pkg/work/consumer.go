package work

import (
	"errors"
	"sync"
	"time"

	"github.com/yanyiwu/simplelog"
)

var (
	ErrAlreadyExist   = errors.New("work: cosumer with the same kind already exist")
	ErrGracefulFailed = errors.New("work: failed to shutdown gracefully")
)

const (
	maxGracefulCheck    = 60
	consumeWaitDuration = time.Hour
	gracefulCheckPeriod = time.Second
)

type Consumer interface {
	Kind() string
	Consume(job Job)
}

type ConsumerSwitch struct {
	MaxProcessing int
	processing    int
	consumers     map[string]Consumer
	queue         *Queue
	end           chan bool
	mu            sync.RWMutex
}

func NewConsumerSwitch(q *Queue) *ConsumerSwitch {
	return &ConsumerSwitch{
		MaxProcessing: 8,
		consumers:     make(map[string]Consumer),
		queue:         q,
		end:           make(chan bool),
	}
}

func (cs *ConsumerSwitch) AddConsumer(con Consumer) error {
	cs.mu.Lock()
	defer cs.mu.Unlock()

	if _, ok := cs.consumers[con.Kind()]; ok {
		return ErrAlreadyExist
	}
	cs.consumers[con.Kind()] = con
	return nil
}

func (cs *ConsumerSwitch) getConsumer(kind string) Consumer {
	cs.mu.RLock()
	defer cs.mu.RUnlock()

	if con, ok := cs.consumers[kind]; ok {
		return con
	}
	return nil
}

func (cs *ConsumerSwitch) GetProcessing() int {
	cs.mu.RLock()
	defer cs.mu.RUnlock()

	return cs.processing
}

func (cs *ConsumerSwitch) increaseProcessing() {
	cs.mu.Lock()
	defer cs.mu.Unlock()

	cs.processing++
}

func (cs *ConsumerSwitch) decreaseProcessing() {
	cs.mu.Lock()
	defer cs.mu.Unlock()

	cs.processing--
}

func (cs *ConsumerSwitch) Start() {
	go cs.switching()
}

func (cs *ConsumerSwitch) End() error {
	select {
	case cs.end <- true:
	default:
	}

	try := 0
	t := time.NewTicker(gracefulCheckPeriod)
	for range t.C {
		if try == maxGracefulCheck {
			return ErrGracefulFailed
		}

		if cs.GetProcessing() == 0 {
			return nil
		}

		try++
	}
	panic("world is weird")
}

func (cs *ConsumerSwitch) switching() {
	for {
		select {
		case <-cs.end:
			simplelog.Info("consumer switching ended")
			return
		default:
			if cs.GetProcessing() >= cs.MaxProcessing {
				time.Sleep(consumeWaitDuration)
				continue
			}

			job, err := cs.queue.popFromWorkQueue()
			if err != nil {
				simplelog.Error("error while getting job from work queue | err: %v", err)
				continue
			}

			con := cs.getConsumer(job.Kind)

			if con == nil {
				simplelog.Error("couldn't find appropriate consumer for job repushing to work queue | kind: %s", job.Kind)
				err = cs.queue.PushToWorkQueue(job)
				if err != nil {
					simplelog.Error("error while repushing the job | err: %v", err)
				}
				continue
			}

			job.TakenAt = time.Now().UTC()

			err = cs.queue.addToProcessingSet(job)
			if err != nil {
				simplelog.Error("redis error while adding job to processing set | err: %v", err)
				continue
			}

			cs.increaseProcessing()
			job.cs = cs
			go con.Consume(job)

			simplelog.Info("gived the job to consumer | job: %v", job)
		}
	}
}
