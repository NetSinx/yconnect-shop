package messaging

import (
	"github.com/NetSinx/yconnect-shop/server/user/internal/helpers"
	amqp "github.com/rabbitmq/amqp091-go"
	"github.com/sirupsen/logrus"
)

type Subscriber struct {
	Connection *amqp.Connection
	Log *logrus.Logger
}

func NewSubscriber(connection *amqp.Connection, Log *logrus.Logger) *Subscriber {
	return &Subscriber{
		Connection: connection,
		Log: Log,
	}
}

func (s *Subscriber) Receive() {
	ch, err := s.Connection.Channel()
	helpers.PanicError(s.Log, err, "failed to open a channel")
	defer ch.Close()

	exchange := "user_events"
	err = ch.ExchangeDeclare(
		exchange,
		"topic",
		true,
		false,
		false,
		false,
		nil,
	)
	helpers.FatalError(s.Log, err, "failed to declare an exchange")

	q, err := ch.QueueDeclare(
		"user_data",
		true,
		false,
		false,
		false,
		nil,
	)
	helpers.FatalError(s.Log, err, "failed to declare a queue")

	err = ch.QueueBind(
		q.Name,
		"user.register",
		exchange,
		false,
		nil,
	)

	msgs, err := ch.Consume(
		q.Name,
		"",
		false,
		false,
		false,
		false,
		nil,
	)

	go func() {
		for d := range msgs {
			
		}
	}()
}