package work

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

type testConsumer5 struct {
}

func (c testConsumer5) Consume(j Job) {
	time.Sleep(time.Second)
}

func (testConsumer5) Kind() string {
	return "test"
}

func TestGarbageCollector(t *testing.T) {
	q := setupQueueForTest()
	q.client.FlushDB()
	a := assert.New(t)

	gc := NewGarbageCollector(q)
	gc.Start()

	cs := NewConsumerHub(q)
	cs.AddConsumer(testConsumer5{})
	cs.Start()

	job := Job{
		Kind:    "test",
		Timeout: 500 * time.Millisecond,
	}

	err := q.PushToWorkQueue(job)
	a.Nil(err)

	cs.end <- true

	time.Sleep(time.Second)

	jobs, err := q.getAllFromProcessingSet()
	a.Nil(err)
	a.Equal(0, len(jobs))
}
