package work

import (
	"encoding/json"
	"errors"
	"sync"
	"time"

	"github.com/go-redis/redis"
	"github.com/yanyiwu/simplelog"
)

var (
	ErrAlreadyExist   = errors.New("work: cosumer with the same kind already exist")
	ErrGracefulFailed = errors.New("work: failed to shutdown gracefully")
)

const (
	redisQueueKey       = "gorani_work_queue"
	redisProcessingKey  = "gorani_processing_set"
	maxGracefulCheck    = 60
	brpopTimeout        = time.Hour
	consumeWaitDuration = time.Hour
	gracefulCheckPeriod = time.Second
	gcDuration          = 10 * time.Minute
)

type Job struct {
	Kind    string
	Payload string
	TakenAt time.Time
	Timeout time.Duration
	queue   *Queue
}

func (j Job) Deadline() time.Time {
	return j.TakenAt.Add(j.Timeout)
}

func (j Job) Complete() {
	j.queue.decreaseProcessing()
	err := j.queue.removeFromProcessingSet(j)
	if err != nil {
		simplelog.Error("error while completing the job | err: %v", err)
		return
	}

	simplelog.Info("job completed | job: %v", j)
}

func (j Job) Fail() {
	j.queue.decreaseProcessing()
	// beleive garbage collecting will send back the job to work queue
	simplelog.Info("job failed | job: %v", j)
}

type Queue struct {
	MaxProcessing int

	mu         sync.RWMutex
	processing int
	client     *redis.Client
	consumers  map[string]Consumer
	consumeEnd chan bool
	garbageEnd chan bool
}

func New(client *redis.Client) *Queue {
	return &Queue{
		MaxProcessing: 8,
		client:        client,
		consumers:     make(map[string]Consumer),
		consumeEnd:    make(chan bool),
		garbageEnd:    make(chan bool),
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

func (q *Queue) GetProcessing() int {
	q.mu.RLock()
	defer q.mu.RUnlock()

	return q.processing
}

func (q *Queue) increaseProcessing() {
	q.mu.Lock()
	defer q.mu.Unlock()

	q.processing++
}

func (q *Queue) decreaseProcessing() {
	q.mu.Lock()
	defer q.mu.Unlock()

	q.processing--
}

// gracefully end
func (q *Queue) End() error {
	select {
	case q.consumeEnd <- true:
	default:
	}

	select {
	case q.garbageEnd <- true:
	default:
	}

	try := 0

	t := time.NewTicker(gracefulCheckPeriod)
	for range t.C {
		if try == maxGracefulCheck {
			return ErrGracefulFailed
		}

		if q.GetProcessing() == 0 {
			return nil
		}

		try++
	}

	panic("world is weird")
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
