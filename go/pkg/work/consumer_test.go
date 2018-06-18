package work_test

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/sunho/gorani-reader-server/go/pkg/util"
	"github.com/sunho/gorani-reader-server/go/pkg/work"
)

var job = work.Job{
	Kind:    "test",
	Payload: "asdf",
	Timeout: 10,
}

type testConsumer struct {
	a *assert.Assertions
}

func (t testConsumer) Consume(j work.Job) {
	t.a.Equal(job.Kind, j.Kind)
	t.a.Equal(job.Payload, j.Payload)
	t.a.Equal(job.Timeout, j.Timeout)
	time.Sleep(50 * time.Millisecond)
	j.Complete()
}

func (testConsumer) Kind() string {
	return "test"
}

func TestConsumer(t *testing.T) {
	gorn := util.SetupTestGorani()
	gorn.Redis.FlushDB()
	a := assert.New(t)

	gorn.WorkQueue.AddConsumer(testConsumer{a})
	gorn.WorkQueue.StartConsuming()

	a.Equal(0, gorn.WorkQueue.GetProcessing())

	err := gorn.WorkQueue.PushToWorkQueue(job)
	a.Nil(err)

	time.Sleep(10 * time.Millisecond)

	a.Equal(1, gorn.WorkQueue.GetProcessing())

	time.Sleep(100 * time.Millisecond)

	a.Equal(0, gorn.WorkQueue.GetProcessing())
}

type testConsumer2 struct {
	a *assert.Assertions
}

func (t testConsumer2) Consume(j work.Job) {
	t.a.Equal(false, true)
}

func (testConsumer2) Kind() string {
	return "test2"
}

type testConsumer3 struct {
	a *assert.Assertions
}

func (t testConsumer3) Consume(j work.Job) {
	time.Sleep(10 * time.Second)
}

func (testConsumer3) Kind() string {
	return "test3"
}

func TestCosumingMany(t *testing.T) {
	gorn := util.SetupTestGorani()
	gorn.Redis.FlushDB()
	a := assert.New(t)

	gorn.WorkQueue.AddConsumer(testConsumer2{a})
	gorn.WorkQueue.AddConsumer(testConsumer3{a})
	gorn.WorkQueue.StartConsuming()

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

	a.Equal(8, gorn.WorkQueue.GetProcessing())
}

type testConsumer4 struct {
	a *assert.Assertions
}

func (t testConsumer4) Consume(j work.Job) {
	time.Sleep(50 * time.Millisecond)
	j.Fail()
}

func (testConsumer4) Kind() string {
	return "test"
}

func TestConsumerFail(t *testing.T) {
	gorn := util.SetupTestGorani()
	gorn.Redis.FlushDB()
	a := assert.New(t)

	gorn.WorkQueue.AddConsumer(testConsumer4{a})
	gorn.WorkQueue.StartConsuming()

	a.Equal(0, gorn.WorkQueue.GetProcessing())

	err := gorn.WorkQueue.PushToWorkQueue(job)
	a.Nil(err)

	time.Sleep(10 * time.Millisecond)

	a.Equal(1, gorn.WorkQueue.GetProcessing())

	time.Sleep(100 * time.Millisecond)

	a.Equal(0, gorn.WorkQueue.GetProcessing())
}
