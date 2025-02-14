package rabbitmq

import (
	"encoding/json"

	"github.com/NetSinx/yconnect-shop/server/order/utils"
	amqp "github.com/rabbitmq/amqp091-go"
)

func ConsumeUserID() (string, error) {
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	utils.FailOnError("failed to connect to RabbitMQ", err)
	defer conn.Close()

	ch, err := conn.Channel()
	utils.FailOnError("failed to open a channel", err)
	defer ch.Close()

	q, err := ch.QueueDeclare(
		"resp_user_id",
		false,
		false,
		false,
		false,
		nil,
	)
	utils.FailOnError("failed to declare a queue", err)

	msgs, err := ch.Consume(
		q.Name,
		"",
		true,
		false,
		false,
		false,
		nil,
	)
	utils.FailOnError("failed to consume message", err)

	for msg := range msgs {
		var respUsername string
		json.Unmarshal(msg.Body, &respUsername)
		return respUsername, nil
	}

	return "", nil
}
