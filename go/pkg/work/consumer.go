package work

import (
	"time"

	"github.com/yanyiwu/simplelog"
)

type Consumer interface {
	Kind() string
	Consume(job Job)
}

func (q *Queue) AddConsumer(con Consumer) error {
	q.mu.Lock()
	defer q.mu.Unlock()

	if _, ok := q.consumers[con.Kind()]; ok {
		return ErrAlreadyExist
	}
	q.consumers[con.Kind()] = con
	return nil
}

func (q *Queue) getConsumer(kind string) Consumer {
	q.mu.RLock()
	defer q.mu.RUnlock()

	if con, ok := q.consumers[kind]; ok {
		return con
	}
	return nil
}

func (q *Queue) StartConsuming() {
	go q.consuming()
}

func (q *Queue) consuming() {
	for {
		select {
		case <-q.consumeEnd:
			simplelog.Info("consuming ended")
			return
		default:
			if q.GetProcessing() >= q.MaxProcessing {
				time.Sleep(consumeWaitDuration)
				continue
			}

			job, err := q.popFromWorkQueue()
			if err != nil {
				simplelog.Error("error while getting job from work queue | err: %v", err)
				continue
			}

			con := q.getConsumer(job.Kind)

			if con == nil {
				simplelog.Error("couldn't find appropriate consumer for job repushing to work queue | kind: %s", job.Kind)
				err = q.PushToWorkQueue(job)
				if err != nil {
					simplelog.Error("error while repushing the job | err: %v", err)
				}
				continue
			}

			job.TakenAt = time.Now().UTC()

			err = q.addToProcessingSet(job)
			if err != nil {
				simplelog.Error("redis error while adding job to processing set | err: %v", err)
				continue
			}

			q.increaseProcessing()
			job.queue = q
			go con.Consume(job)

			simplelog.Info("gived the job to consumer | job: %v", job)
		}
	}
}
