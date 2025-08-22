package rabbitmq

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"time"
	"github.com/NetSinx/yconnect-shop/server/category/errs"
	amqp "github.com/rabbitmq/amqp091-go"
)

var (
	RoutingCKCreated = "category.created"
	RoutingCKUpdated = "category.updated"
	RoutingCKDeleted = "category.deleted"
)

func Publisher(routingKey string, message string) error {
	amqpURL := fmt.Sprintf("amqp://%s:%s@%s:%s", 
		os.Getenv("RABBITMQ_USER"),
		os.Getenv("RABBITMQ_PASSWORD"),
		os.Getenv("RABBITMQ_HOST"),
		os.Getenv("RABBITMQ_PORT"),
	)
	conn, err := amqp.Dial(amqpURL)
	errs.LogOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	ch, err := conn.Channel()
	errs.LogOnError(err, "Failed to open a channel")
	defer ch.Close()

	exchange := "category.events"
	err = ch.ExchangeDeclare(
		exchange,
		"topic",
		true,
		false,
		false,
		false,
		nil,
	)
	errs.LogOnError(err, "Failed to declare an exchange")

	body, _ := json.Marshal(message)
	ctx, cancel := context.WithTimeout(context.Background(), 5 * time.Second)
	defer cancel()

	return ch.PublishWithContext(ctx,
		exchange,
		routingKey,
		false,
		false,
		amqp.Publishing{
		ContentType: "application/json",
		Body:        body,
	})
}