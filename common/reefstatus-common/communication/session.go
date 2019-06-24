package communication

// Session interface
type Session interface {
	Publish(message string, data string) error
	PublishData(message string, data []byte) error
	Subscribe(message string) (chan string, error)
	QueueSubscribe(message string, queue string) (chan string, error)
	Close()
}

// NewSession creates a new session
func NewSession(address, token string) (Session, error) {
	return newNatsSession(address, token)
}
