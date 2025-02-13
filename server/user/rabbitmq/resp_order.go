package rabbitmq

import (
	"encoding/json"
	"github.com/NetSinx/yconnect-shop/server/order/model/entity"
	"github.com/NetSinx/yconnect-shop/server/user/utils"
	amqp "github.com/rabbitmq/amqp091-go"
)

func RequestGetOrderByUsername(username string) entity.Order {
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	utils.FailOnError("Failed to connect to RabbitMQ", err)
	defer conn.Close()

	ch, err := conn.Channel()
	utils.FailOnError("Failed to open a channel", err)
	defer ch.Close()

	q, err := ch.QueueDeclare(
		"resp_order",
		false,
		false,
		false,
		false,
		nil,
	)
	utils.FailOnError("Failed to declare queue", err)

	msgs, err := ch.Consume(
		q.Name,
		"",
		true,
		false,
		false,
		false,
		nil,
	)
	utils.FailOnError("Failed to consume messages", err)

	for msg := range msgs {
		var respOrder entity.Order
		json.Unmarshal(msg.Body, &respOrder)
		return respOrder
	}

	return entity.Order{}
}