package work_test

import (
	"testing"

	"github.com/sunho/gorani-reader-server/go/pkg/util"
	"github.com/sunho/gorani-reader-server/go/pkg/work"
)

func TestGarbageCollector(t *testing.T) {
	gorn := util.SetupTestGorani()
	cs := work.NewConsumerSwitch(gorn.WorkQueue)

}
