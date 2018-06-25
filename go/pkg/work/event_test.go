package work_test

import (
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/sunho/gorani-reader-server/go/pkg/util"
	"github.com/sunho/gorani-reader-server/go/pkg/work"
)

type testConsumer6 struct {
}

func (t testConsumer6) Consume(j work.Job) {
	time.Sleep(200 * time.Millisecond)
	j.Complete(work.Result{Payload: "gogo", Success: true})
}

func (testConsumer6) Kind() string {
	return "test"
}

func TestWaitForResult(t *testing.T) {
	a := assert.New(t)
	gorn := util.SetupTestGorani()
	gorn.Redis.FlushDB()

	eh := work.NewEventHub(gorn.WorkQueue)
	eh.Start()

	time.Sleep(500 * time.Millisecond)

	cs := work.NewConsumerHub(gorn.WorkQueue)
	cs.AddConsumer(testConsumer6{})
	cs.Start()

	job := work.Job{
		Kind:    "test",
		Uuid:    uuid.New(),
		Timeout: time.Second,
	}
	err := gorn.WorkQueue.PushToWorkQueue(job)
	a.Nil(err)

	res, err := eh.WaitForResult(job)
	a.Nil(err)
	a.Equal("gogo", res.Payload)

	eh.End()
	cs.End()
}
