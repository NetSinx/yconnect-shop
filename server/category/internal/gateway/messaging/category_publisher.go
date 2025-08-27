package messaging

import (
	"context"
	"encoding/json"
	"time"
	"github.com/NetSinx/yconnect-shop/server/category/internal/helpers"
	amqp "github.com/rabbitmq/amqp091-go"
)

type Publisher struct {
	Connection *amqp.Connection
	Helpers    *helpers.Helpers
}

func NewPublisher(connection *amqp.Connection, helpers *helpers.Helpers) *Publisher {
	return &Publisher{
		Connection: connection,
		Helpers:    helpers,
	}
}

func (p *Publisher) Send(routingKey string, message any) {
	ch, err := p.Connection.Channel()
	p.Helpers.PanicError(err, "Failed to open a channel")
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
	p.Helpers.FatalError(err, "Failed to declare an exchange")

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
	p.Helpers.FatalError(err, "Failed to publish a message")
}
