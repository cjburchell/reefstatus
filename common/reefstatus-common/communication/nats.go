package communication

import (
	"github.com/cjburchell/go-uatu"
	"github.com/nats-io/go-nats"
	"github.com/pkg/errors"
)

type natsSession struct {
	nc            *nats.Conn
	subscriptions map[string]*nats.Subscription
}

func newNatsSession(address, token string) (Session, error) {
	nc, err := nats.Connect(address, nats.Token(token),
		nats.DisconnectHandler(func(nc *nats.Conn) {
			log.Printf("Got disconnected!\n")
		}),
		nats.ReconnectHandler(func(nc *nats.Conn) {
			log.Printf("Got reconnected to %v!\n", nc.ConnectedUrl())
		}),
		nats.ClosedHandler(func(nc *nats.Conn) {
			log.Printf("Connection closed. Reason: %q\n", nc.LastError())
		}))

	if err != nil {
		return nil, errors.WithStack(err)
	}

	var session natsSession
	session.nc = nc
	session.subscriptions = make(map[string]*nats.Subscription)
	return &session, nil
}

func (session natsSession) Publish(message string, data string) error {
	log.Debugf("Publish message %s, data: %s", message, data)
	err := session.nc.Publish(message, []byte(data))
	if err != nil {
		return errors.WithStack(err)
	}

	return nil
}

func (session natsSession) PublishData(message string, data []byte) error {
	log.Debugf("Publish data %s, data: %s", message, data)
	err := session.nc.Publish(message, data)
	if err != nil {
		return errors.WithStack(err)
	}

	return nil
}

func (session *natsSession) Subscribe(message string) (chan string, error) {
	log.Debugf("Subscribe to %s", message)
	ch := make(chan string)
	sub, err := session.nc.Subscribe(message, func(msg *nats.Msg) {
		log.Debugf("Received Message %s", message)
		ch <- string(msg.Data)
	})
	if err != nil {
		return nil, errors.WithStack(err)
	}
	session.subscriptions[message] = sub
	return ch, nil
}

func (session *natsSession) QueueSubscribe(message string, queue string) (chan string, error) {
	log.Debugf("Queue Subscribe to %s, queue: %s", message, queue)
	ch := make(chan string)
	sub, err := session.nc.QueueSubscribe(message, queue, func(msg *nats.Msg) {
		log.Debugf("Received Message %s from queue: %s", message, queue)
		ch <- string(msg.Data)
	})
	if err != nil {
		return nil, errors.WithStack(err)
	}
	session.subscriptions[message] = sub
	return ch, nil
}

func (session natsSession) Close() {
	session.nc.Close()
	for _, sub := range session.subscriptions {
		sub.Unsubscribe()
	}
}
