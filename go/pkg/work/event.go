package work

import (
	"errors"
	"sync"
	"time"

	"github.com/google/uuid"
	"github.com/yanyiwu/simplelog"
)

var (
	ErrTimeOut = errors.New("work: timeout")
)

type KindListener interface {
	Listen(Result)
	Kind() string
}

type kindUuid struct {
	kind string
	id   uuid.UUID
}

func newKindUuid(kind string, id uuid.UUID) kindUuid {
	return kindUuid{
		kind: kind,
		id:   id,
	}
}

type jobListener struct {
	eh  *EventHub
	key kindUuid
	c   chan Result
}

func (j *jobListener) close() {
	delete(j.eh.jobListeners, j.key)
}

func (j *jobListener) Listen(res Result) {
	j.c <- res
}

func newJobListner(eh *EventHub, job Job) *jobListener {
	return &jobListener{
		eh:  eh,
		key: newKindUuid(job.Kind, job.Uuid),
		c:   make(chan Result),
	}
}

type EventHub struct {
	queue         *Queue
	kindListeners map[string]KindListener
	jobListeners  map[kindUuid]*jobListener
	end           chan bool
	mu            sync.RWMutex
}

func NewEventHub(queue *Queue) *EventHub {
	return &EventHub{
		queue:         queue,
		kindListeners: make(map[string]KindListener),
		jobListeners:  make(map[kindUuid]*jobListener),
		end:           make(chan bool),
	}
}

func (eh *EventHub) WaitForResult(job Job) (Result, error) {
	jl := newJobListner(eh, job)
	eh.jobListeners[newKindUuid(job.Kind, job.Uuid)] = jl
	defer jl.close()

	select {
	case res := <-jl.c:
		return res, nil
	case <-time.After(job.Timeout):
		return Result{}, ErrTimeOut
	}
}

// one listener per one kind
func (eh *EventHub) AddKindListener(lis KindListener) {
	eh.mu.Lock()
	defer eh.mu.Unlock()
	eh.kindListeners[lis.Kind()] = lis
}

func (eh *EventHub) DeleteKindListener(lis KindListener) {
	eh.mu.Lock()
	defer eh.mu.Unlock()
	delete(eh.kindListeners, lis.Kind())
}

func (eh *EventHub) End() {
	close(eh.end)
}

func (eh *EventHub) Start() {
	go func() {
		c, done := eh.queue.subscribeToEventChannel()
		defer func() {
			close(done)
		}()

		for {
			select {
			case <-eh.end:
				simplelog.Info("event hub endded")
				return

			case res := <-c:
				kinduuid := newKindUuid(res.Kind, res.Uuid)
				if lis, ok := eh.jobListeners[kinduuid]; ok {
					simplelog.Info("gave result to jobListener | result: %v", res)
					lis.Listen(res)
				}

				if lis, ok := eh.kindListeners[res.Kind]; ok {
					simplelog.Info("gave result to kindListener | result: %v", res)
					lis.Listen(res)
				}
			}
		}
	}()
}
