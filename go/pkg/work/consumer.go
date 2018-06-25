package work

import (
	"errors"
	"log"
	"strings"
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
	consumeWaitDuration = time.Second
	gracefulCheckPeriod = time.Second
)

type Consumer interface {
	Kind() string
	Consume(job Job)
}

type ConsumerHub struct {
	MaxProcessing int
	processing    int
	consumers     map[string]Consumer
	queue         *Queue
	end           chan bool
	mu            sync.RWMutex
}

func NewConsumerHub(q *Queue) *ConsumerHub {
	return &ConsumerHub{
		MaxProcessing: 8,
		consumers:     make(map[string]Consumer),
		queue:         q,
		end:           make(chan bool),
	}
}

func (ch *ConsumerHub) AddConsumer(con Consumer) error {
	ch.mu.Lock()
	defer ch.mu.Unlock()

	if _, ok := ch.consumers[con.Kind()]; ok {
		return ErrAlreadyExist
	}
	ch.consumers[con.Kind()] = con
	return nil
}

func (ch *ConsumerHub) getConsumer(kind string) Consumer {
	ch.mu.RLock()
	defer ch.mu.RUnlock()

	if con, ok := ch.consumers[kind]; ok {
		return con
	}
	return nil
}

func (ch *ConsumerHub) GetProcessing() int {
	ch.mu.RLock()
	defer ch.mu.RUnlock()

	return ch.processing
}

func (ch *ConsumerHub) increaseProcessing() {
	ch.mu.Lock()
	defer ch.mu.Unlock()

	ch.processing++
}

func (ch *ConsumerHub) decreaseProcessing() {
	ch.mu.Lock()
	defer ch.mu.Unlock()

	ch.processing--
}

func (ch *ConsumerHub) End() error {
	log.Println("trying to stop consuming")
	close(ch.end)

	try := 0
	t := time.NewTicker(gracefulCheckPeriod)
	log.Println("gracefully shutdowning consumers")
	for range t.C {
		if try == maxGracefulCheck {
			return ErrGracefulFailed
		}

		if ch.GetProcessing() == 0 {
			return nil
		}

		try++
	}
	panic("world is weird")
}

func (ch *ConsumerHub) Start() {
	go func() {
		for {
			select {
			case <-ch.end:
				simplelog.Info("consumer hub ended")
				return
			default:
				if ch.GetProcessing() >= ch.MaxProcessing {
					time.Sleep(consumeWaitDuration)
					continue
				}

				job, err := ch.queue.popFromWorkQueue()
				if err != nil {
					if !strings.Contains(err.Error(), "nil") {
						simplelog.Error("error while getting job from work queue | err: %v", err)
					}
					continue
				}

				con := ch.getConsumer(job.Kind)

				if con == nil {
					simplelog.Error("couldn't find appropriate consumer for job repushing to work queue | kind: %s", job.Kind)
					err = ch.queue.PushToWorkQueue(job)
					if err != nil {
						simplelog.Error("error while repushing the job | err: %v", err)
					}
					continue
				}

				job.TakenAt = time.Now().UTC()

				err = ch.queue.addToProcessingSet(job)
				if err != nil {
					simplelog.Error("redis error while adding job to processing set | err: %v", err)
					continue
				}

				ch.increaseProcessing()
				job.ch = ch
				go con.Consume(job)

				simplelog.Info("gave the job to consumer | job: %v", job)
			}
		}
	}()
}
