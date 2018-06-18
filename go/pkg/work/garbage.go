package work

import (
	"encoding/json"
	"time"

	"github.com/yanyiwu/simplelog"
)

func (q *Queue) StartGarbageCollecting() {
	go q.garbageCollecting()
}

func (q *Queue) garbageCollecting() {
	t := time.NewTicker(gcDuration)
	for {
		select {
		case <-q.garbageEnd:
			simplelog.Info("garbageCollecting ended")
			return
		case <-t.C:
			simplelog.Info("start work garbage collecting")

			result := q.client.SMembers(redisProcessingKey)
			err := result.Err()
			if err != nil {
				simplelog.Error("redis error while getting processing set: %v", err)
				continue
			}

			vals := result.Val()
			for _, val := range vals {
				job := Job{}

				err = json.Unmarshal([]byte(val), &job)
				if err != nil {
					simplelog.Error("json unmarshal failed | val: %s, err: %v", val, err)
					continue
				}

				if time.Now().UTC().Before(job.Deadline()) {
					continue
				}

				simplelog.Info("job missed the deadline sending back to the queue | job: %v", job)

				err = q.removeFromProcessingSet(job)
				if err != nil {
					simplelog.Error("redis error while removing job from processing set | job: %v err: %v", job, err)
					continue
				}

				err = q.PushToWorkQueue(job)
				if err != nil {
					simplelog.Error("error while pushing job to the queue the job will be discarded | job: %v", job)
				}
			}
		}
	}
}
