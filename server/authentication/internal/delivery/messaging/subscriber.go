package messaging

import (
	amqp "github.com/rabbitmq/amqp091-go"
	"github.com/sirupsen/logrus"
)

type Subscriber struct {
	RabbitMQ *amqp.Connection
	Log *logrus.Logger
}

func NewSubscriber(rabbitmq *amqp.Connection, log *logrus.Logger) *Subscriber {
	return &Subscriber{
		RabbitMQ: rabbitmq,
		Log: log,
	}
}

func (s *Subscriber) Receive() {
	
}