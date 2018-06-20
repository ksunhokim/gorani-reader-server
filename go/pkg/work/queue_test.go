package work

import (
	"os"
	"testing"
	"time"

	"github.com/go-redis/redis"
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
		Addr: redisaddr,
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
