package work

import (
	"io/ioutil"
	"testing"
	"time"

	"github.com/go-redis/redis"
	"github.com/stretchr/testify/assert"
	yaml "gopkg.in/yaml.v2"
)

func setupQueueForTest() *Queue {
	buf, err := ioutil.ReadFile("../../config_test.yaml")
	if err != nil {
		panic(err)
	}

	config := struct {
		RedisUrl string `yaml:"redis_url"`
	}{}
	err = yaml.Unmarshal(buf, &config)
	if err != nil {
		panic(err)
	}

	cli := redis.NewClient(&redis.Options{
		Addr: "127.0.0.1:6379",
	})

	return New(cli)
}

func TestQueue(t *testing.T) {
	a := assert.New(t)

	q := setupQueueForTest()
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
