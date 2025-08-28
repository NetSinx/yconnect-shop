package messaging

import (
	"context"
	"encoding/json"
	"time"

	"github.com/NetSinx/yconnect-shop/server/category/internal/helpers"
	amqp "github.com/rabbitmq/amqp091-go"
	"github.com/sirupsen/logrus"
)

type Publisher struct {
	Connection *amqp.Connection
	Log *logrus.Logger
}

func NewPublisher(connection *amqp.Connection, log *logrus.Logger) *Publisher {
	return &Publisher{
		Connection: connection,
		Log: log,
	}
}

func (p *Publisher) Send(routingKey string, message any) {
	ch, err := p.Connection.Channel()
	helpers.PanicError(p.Log, err, "Failed to open a channel")
	defer ch.Close()

	exchange := "category_events"
	err = ch.ExchangeDeclare(
		exchange,
		"topic",
		true,
		false,
		false,
		false,
		nil,
	)
	helpers.FatalError(p.Log, err, "Failed to declare an exchange")

	body, _ := json.Marshal(message)
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err = ch.PublishWithContext(ctx,
		exchange,
		routingKey,
		false,
		false,
		amqp.Publishing{
			ContentType: "application/json",
			Body:        body,
		})
	helpers.FatalError(p.Log, err, "Failed to publish a message")
}
