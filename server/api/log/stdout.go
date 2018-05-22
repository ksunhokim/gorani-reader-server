package log

import (
	"fmt"
)

type StdoutLogger struct {
}

func (s *StdoutLogger) Log(tag string, data interface{}) {
	fmt.Println(tag, ":", data)
}

func NewStdoutLogger() *StdoutLogger {
	return &StdoutLogger{}
}
