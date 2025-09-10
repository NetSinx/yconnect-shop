package messaging

import (
	"context"
	"encoding/json"
	"time"
	"github.com/NetSinx/yconnect-shop/server/authentication/internal/helpers"
	"github.com/NetSinx/yconnect-shop/server/authentication/internal/model"
	amqp "github.com/rabbitmq/amqp091-go"
	"github.com/sirupsen/logrus"
)

type Publisher struct {
	RabbitMQ *amqp.Connection
	Log      *logrus.Logger
}

func NewPublisher(rabbitmq *amqp.Connection, log *logrus.Logger) *Publisher {
	return &Publisher{
		RabbitMQ: rabbitmq,
		Log: log,
	}
}

func (p *Publisher) Send(ctx context.Context, message *model.RegisterUserEvent) {
	ch, err := p.RabbitMQ.Channel()
	helpers.PanicError(p.Log, err, "failed to open a channel")
	defer ch.Close()

	exchange := "user_auth_events"
	err = ch.ExchangeDeclare(
		exchange,
		"direct",
		true,
		false,
		false,
		false,
		nil,
	)
	helpers.FatalError(p.Log, err, "failed to declare an exchange")

	c, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	msgByte, err := json.Marshal(message)
	helpers.FatalError(p.Log, err, "failed to marshaling message")

	err = ch.PublishWithContext(
		c,
		exchange,
		"user_register",
		false,
		false,
		amqp.Publishing{
			ContentType: "application/json",
			Body: msgByte,
		},
	)
	helpers.FatalError(p.Log, err, "failed to publish a message body")
}