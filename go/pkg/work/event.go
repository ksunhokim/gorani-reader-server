package work

import (
	"sync"

	"github.com/google/uuid"
)

type KindListener interface {
	Listen(Result)
	Kind() string
}

type JobListener interface {
	Listen(Result)
	Kind() string
	Uuid() uuid.UUID
}

type EventHub struct {
	queue         *Queue
	kindListeners map[string]KindListener
	jobListeners  map[kindUuid]JobListener
	end           chan bool
	mu            sync.RWMutex
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

func (e *EventHub) AddKindListener(lis KindListener) {
	e.mu.Lock()
	defer e.mu.Unlock()
	e.kindListeners[lis.Kind()] = lis
}

func (e *EventHub) DeleteListener(lis KindListener) {
	e.mu.Lock()
	defer e.mu.Unlock()
	delete(e.kindListeners, lis.Kind())
}

func (e *EventHub) AddJobListener(lis JobListener) {
	e.mu.Lock()
	defer e.mu.Unlock()
	kinduuid := newKindUuid(lis.Kind(), lis.Uuid())
	e.jobListeners[kinduuid] = lis
}

func (e *EventHub) DeleteJobListener(lis JobListener) {
	e.mu.Lock()
	defer e.mu.Unlock()
	kinduuid := newKindUuid(lis.Kind(), lis.Uuid())
	delete(e.jobListeners, kinduuid)
}

func (e *EventHub) End() {
	e.end <- true
}

func (e *EventHub) Start() {
	c, done := e.queue.subscribeToEventChannel()
	defer func() {
		done <- true
	}()

	for {
		select {
		case <-e.end:
			return

		case res := <-c:
			kinduuid := newKindUuid(res.Kind, res.Uuid)
			if lis, ok := e.jobListeners[kinduuid]; ok {
				lis.Listen(res)
			}

			if lis, ok := e.kindListeners[res.Kind]; ok {
				lis.Listen(res)
			}
		}
	}
}
