package communication

import logger "github.com/cjburchell/uatu-go"

// Session interface
type Session interface {
	SubscribeSession
	PublishSession
	Close()
}

type SubscribeSession interface {
	QueueSubscribe(message string, queue string) (chan string, error)
}

type PublishSession interface {
	Publish(message string, data string) error
	PublishData(message string, data []byte) error
}

// NewSession creates a new session
func NewSession(address, token string, log logger.ILog) (Session, error) {
	return newNatsSession(address, token, log)
}
