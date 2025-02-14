package rabbitmq

import (
	"encoding/json"
	"log"
	"github.com/NetSinx/yconnect-shop/server/order/utils"
	amqp "github.com/rabbitmq/amqp091-go"
)

func GetUserID(username string) {
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	utils.FailOnError("failed to connect to RabbitMQ", err)
	defer conn.Close()

	ch, err := conn.Channel()
	utils.FailOnError("failed to open a channel", err)
	defer ch.Close()

	q, err := ch.QueueDeclare(
		"req_username",
		false,
		false,
		false,
		false,
		nil,
	)
	utils.FailOnError("failed to declare a queue", err)
	
	body, err := json.Marshal(username)
	utils.FailOnError("failed to marshaling message body", err)

	err = ch.Publish(
		"",     // exchange
		q.Name, // routing key
		false,  // mandatory
		false,  // immediate
		amqp.Publishing {
			ContentType: "text/plain",
			Body:        body,
		},
	)
	utils.FailOnError("failed to publish message", err)
	log.Printf("Sent a message: %s", username)
}
