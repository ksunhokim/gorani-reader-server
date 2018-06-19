package work_test

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/sunho/gorani-reader-server/go/pkg/util"
	"github.com/sunho/gorani-reader-server/go/pkg/work"
)

type testConsumer struct {
	a *assert.Assertions
}

var job = work.Job{
	Kind:    "test",
	Payload: "asdf",
	Timeout: 10,
}

func (t testConsumer) Consume(j work.Job) {
	t.a.Equal(job.Kind, j.Kind)
	t.a.Equal(job.Payload, j.Payload)
	t.a.Equal(job.Timeout, j.Timeout)
	time.Sleep(200 * time.Millisecond)
	j.Complete(work.Result{Success: true})
}

func (testConsumer) Kind() string {
	return "test"
}

func TestConsumer(t *testing.T) {
	gorn := util.SetupTestGorani()
	gorn.Redis.FlushDB()
	a := assert.New(t)

	cs := work.NewConsumerSwitch(gorn.WorkQueue)
	cs.AddConsumer(testConsumer{a})
	cs.Start()

	a.Equal(0, cs.GetProcessing())

	err := gorn.WorkQueue.PushToWorkQueue(job)
	a.Nil(err)

	time.Sleep(100 * time.Millisecond)

	a.Equal(1, cs.GetProcessing())

	time.Sleep(300 * time.Millisecond)

	a.Equal(0, cs.GetProcessing())

	err = cs.End()
	a.Nil(err)
}

type testConsumer2 struct {
	a *assert.Assertions
}

func (t testConsumer2) Consume(j work.Job) {
	t.a.Equal(false, true)
	j.Complete(work.Result{Success: true})
}

func (testConsumer2) Kind() string {
	return "test2"
}

type testConsumer3 struct {
	a *assert.Assertions
}

func (t testConsumer3) Consume(j work.Job) {
	time.Sleep(200 * time.Millisecond)
	j.Complete(work.Result{Success: true})
}

func (testConsumer3) Kind() string {
	return "test3"
}

func TestCosumingMany(t *testing.T) {
	gorn := util.SetupTestGorani()
	gorn.Redis.FlushDB()
	a := assert.New(t)

	cs := work.NewConsumerSwitch(gorn.WorkQueue)
	cs.AddConsumer(testConsumer2{a})
	cs.AddConsumer(testConsumer3{a})
	cs.Start()

	job3 := work.Job{
		Kind: "test3",
	}

	for i := 0; i < 8; i++ {
		err := gorn.WorkQueue.PushToWorkQueue(job3)
		a.Nil(err)
	}

	job2 := work.Job{
		Kind: "test2",
	}

	err := gorn.WorkQueue.PushToWorkQueue(job2)
	a.Nil(err)

	time.Sleep(100 * time.Millisecond)

	a.Equal(8, cs.GetProcessing())

	err = cs.End()
	a.Nil(err)

	a.Equal(0, cs.GetProcessing())
}

type testConsumer4 struct {
	a *assert.Assertions
}

func (t testConsumer4) Consume(j work.Job) {
	time.Sleep(100 * time.Millisecond)
	j.Complete(work.Result{Success: false})
}

func (testConsumer4) Kind() string {
	return "test"
}

func TestConsumerFail(t *testing.T) {
	gorn := util.SetupTestGorani()
	gorn.Redis.FlushDB()
	a := assert.New(t)

	cs := work.NewConsumerSwitch(gorn.WorkQueue)
	cs.AddConsumer(testConsumer4{a})
	cs.Start()

	a.Equal(0, cs.GetProcessing())

	err := gorn.WorkQueue.PushToWorkQueue(job)
	a.Nil(err)

	time.Sleep(50 * time.Millisecond)

	a.Equal(1, cs.GetProcessing())

	time.Sleep(200 * time.Millisecond)

	a.Equal(0, cs.GetProcessing())

	err = cs.End()
	a.Nil(err)
}
