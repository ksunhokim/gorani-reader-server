package work

import (
	"os"
	"testing"
	"time"

	"github.com/go-redis/redis"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func setupQueueForTest() *Queue {
	isCI := os.Getenv("ISCI")

	redisaddr := ""
	if isCI == "true" {
		redisaddr = "redis:6379"
	} else {
		redisaddr = "127.0.0.1:6379"
	}

	cli := redis.NewClient(&redis.Options{
		Addr:     redisaddr,
		PoolSize: 2,
	})

	return NewQueue(cli)
}

func TestQueue(t *testing.T) {
	a := assert.New(t)

	q := setupQueueForTest()
	q.client.FlushDB()
	job := Job{
		Kind:    "asdf",
		Payload: "asdf",
		TakenAt: time.Now().UTC(),
	}

	err := q.PushToWorkQueue(job)
	a.Nil(err)

	job2, err := q.popFromWorkQueue()
	a.Nil(err)
	a.Equal(job, job2)
}

func TestProcessing(t *testing.T) {
	a := assert.New(t)

	q := setupQueueForTest()
	q.client.FlushDB()
	job := Job{
		Kind:    "asdf",
		Payload: "asdf",
		TakenAt: time.Now().UTC(),
	}

	err := q.addToProcessingSet(job)
	a.Nil(err)

	err = q.removeFromProcessingSet(job)
	a.Nil(err)
}

func TestSubscribe(t *testing.T) {
	a := assert.New(t)

	q := setupQueueForTest()
	q.client.FlushDB()
	c, _ := q.subscribeToEventChannel()
	res := Result{
		Uuid:    uuid.New(),
		Kind:    "hello",
		Payload: "world",
		Success: true,
	}

	time.Sleep(500 * time.Millisecond)
	err := q.publishToEventChannel(res)
	a.Nil(err)

	select {
	case res2 := <-c:
		a.Equal(res, res2)
	case <-time.After(2 * time.Second):
		a.Fail("Timeout")
	}
}
