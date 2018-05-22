package log

import (
	"fmt"

	"github.com/fluent/fluent-logger-golang/fluent"
)

type FluentLogger struct {
	flu *fluent.Fluent
}

func (f *FluentLogger) Log(tag string, data interface{}) {
	err := f.flu.Post(tag, data)
	if err != nil {
		fmt.Println(err)
	}
}

func NewFluentLogger(host string, port int) (*FluentLogger, error) {
	conf := fluent.Config{
		FluentHost: host,
		FluentPort: port,
	}

	flu, err := fluent.New(conf)
	if err != nil {
		return nil, err
	}

	return &FluentLogger{flu: flu}, nil
}
