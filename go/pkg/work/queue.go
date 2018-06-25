package work

import (
	"encoding/json"
	"time"

	"github.com/go-redis/redis"
	"github.com/google/uuid"
	"github.com/yanyiwu/simplelog"
)

const (
	redisQueueKey      = "gorani_work_queue"
	redisProcessingKey = "gorani_processing_set"
	redisEventChannel  = "gorani_event_channel"
	brpopTimeout       = time.Second
)

type Result struct {
	Uuid    uuid.UUID
	Kind    string
	Payload string
	Success bool
}

type Job struct {
	Kind    string
	Uuid    uuid.UUID
	Payload string
	TakenAt time.Time
	Timeout time.Duration
	ch      *ConsumerHub
}

func (j Job) Deadline() time.Time {
	return j.TakenAt.Add(j.Timeout)
}

func (j Job) Complete(res Result) {
	j.ch.decreaseProcessing()
	if !res.Success {
		simplelog.Info("job failed | job: %v", j)
	} else {
		err := j.ch.queue.removeFromProcessingSet(j)
		if err != nil {
			simplelog.Error("error while completing the job | err: %v", err)
			return
		}
		simplelog.Info("job suceeded | job: %v res: %v", j, res)
	}

	res.Kind = j.Kind
	res.Uuid = j.Uuid
	err := j.ch.queue.publishToEventChannel(res)
	if err != nil {
		simplelog.Error("error while publishing to event channel | err: %v", err)
	}
}

func NewJob() Job {
	return Job{
		Uuid: uuid.New(),
	}
}

type Queue struct {
	client *redis.Client
}

func NewQueue(client *redis.Client) *Queue {
	return &Queue{
		client: client,
	}
}

func (q *Queue) PushToWorkQueue(job Job) error {
	buf, err := json.Marshal(job)
	if err != nil {
		return err
	}

	_, err = q.client.LPush(redisQueueKey, buf).Result()
	if err != nil {
		return err
	}

	return nil
}

func (q *Queue) popFromWorkQueue() (job Job, err error) {
	strs, err := q.client.BRPop(brpopTimeout, redisQueueKey).Result()
	if err != nil {
		return
	}

	str := strs[1]
	err = json.Unmarshal([]byte(str), &job)
	if err != nil {
		return
	}

	return
}

func (q *Queue) addToProcessingSet(job Job) error {
	buf, err := json.Marshal(job)
	if err != nil {
		return err
	}

	_, err = q.client.SAdd(redisProcessingKey, buf).Result()
	if err != nil {
		return err
	}

	return nil
}

func (q *Queue) removeFromProcessingSet(job Job) error {
	buf, err := json.Marshal(job)
	if err != nil {
		return err
	}

	_, err = q.client.SRem(redisProcessingKey, buf).Result()
	if err != nil {
		return err
	}

	return nil
}

func (q *Queue) getAllFromProcessingSet() (jobs []Job, err error) {
	result := q.client.SMembers(redisProcessingKey)
	err = result.Err()
	if err != nil {
		simplelog.Error("redis error while getting processing set: %v", err)
		return
	}

	vals := result.Val()
	for _, val := range vals {
		job := Job{}

		err = json.Unmarshal([]byte(val), &job)
		if err != nil {
			simplelog.Error("json unmarshal failed | val: %s, err: %v", val, err)
			err = nil
			continue
		}

		jobs = append(jobs, job)
	}

	return
}

func (q *Queue) publishToEventChannel(res Result) error {
	buf, err := json.Marshal(res)
	if err != nil {
		return err
	}

	err = q.client.Publish(redisEventChannel, buf).Err()
	return err
}

func (q *Queue) subscribeToEventChannel() (c <-chan Result, done chan bool) {
	pubsub := q.client.Subscribe(redisEventChannel)
	c, done = eventChannelAdapter(pubsub)
	return
}

func eventChannelAdapter(pubsub *redis.PubSub) (<-chan Result, chan bool) {
	old := pubsub.Channel()
	new := make(chan Result)
	done := make(chan bool)

	go func() {
		for {
			select {
			case <-done:
				err := pubsub.Close()
				if err != nil {
					simplelog.Error("error while closing pubsub | err: %v", err)
				}
				return
			case msg := <-old:
				buf := msg.Payload
				res := Result{}
				err := json.Unmarshal([]byte(buf), &res)
				if err != nil {
					simplelog.Error("error while json unmarshal | res: %v err: %v", buf, err)
					break
				}
				new <- res
			}
		}
	}()

	return new, done
}
