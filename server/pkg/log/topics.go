package log

type Topic string

const (
	TopicConfig  Topic = "config"
	TopicSystem  Topic = "system"
	TopicRequest Topic = "request"
	TopicError   Topic = "error"
)

func (t Topic) Api() Topic {
	return "api." + t
}
