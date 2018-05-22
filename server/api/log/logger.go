package log

type M map[string]interface{}

type Logger interface {
	Log(tag string, data interface{})
}

type BothLogger struct {
	logger1 Logger
	logger2 Logger
}

func (b *BothLogger) Log(tag string, data interface{}) {
	b.logger1.Log(tag, data)
	b.logger2.Log(tag, data)
}

func NewBothLogger(logger1 Logger, logger2 Logger) *BothLogger {
	return &BothLogger{logger1: logger1, logger2: logger2}
}
