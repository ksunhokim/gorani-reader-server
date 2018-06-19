package work

import (
	"time"

	"github.com/yanyiwu/simplelog"
)

const (
	gcDuration = time.Second
)

type GarbageCollector struct {
	queue *Queue
	end   chan bool
}

func NewGarbageCollector(q *Queue) *GarbageCollector {
	return &GarbageCollector{
		queue: q,
		end:   make(chan bool),
	}
}

func (gc *GarbageCollector) Start() {
	go gc.garbageCollecting()
}

func (gc *GarbageCollector) End() {
	gc.end <- true
}

func (gc *GarbageCollector) garbageCollecting() {
	t := time.NewTicker(gcDuration)
	for {
		select {
		case <-gc.end:
			simplelog.Info("garbageCollecting ended")
			return
		case <-t.C:
			simplelog.Info("start work garbage collecting")
			jobs, err := gc.queue.getAllFromProcessingSet()
			if err != nil {
				simplelog.Error("error while getting jobs from processing set | err: %v", err)
				continue
			}

			for _, job := range jobs {
				if time.Now().UTC().Before(job.Deadline()) {
					continue
				}

				simplelog.Info("job missed the deadline sending back to the queue | job: %v", job)

				err = gc.queue.removeFromProcessingSet(job)
				if err != nil {
					simplelog.Error("redis error while removing job from processing set | job: %v err: %v", job, err)
					continue
				}

				err = gc.queue.PushToWorkQueue(job)
				if err != nil {
					simplelog.Error("error while pushing job to the queue the job will be discarded | job: %v", job)
				}
			}
		}
	}
}
