package rabbitmq

import (
	"encoding/json"
	"log"
	"github.com/NetSinx/yconnect-shop/server/order/utils"
	amqp "github.com/rabbitmq/amqp091-go"
)

func GetProductByID(product_id int) {
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	utils.FailOnError("failed to connect to RabbitMQ", err)
	defer conn.Close()

	ch, err := conn.Channel()
	utils.FailOnError("failed to open a channel", err)
	defer ch.Close()

	q, err := ch.QueueDeclare(
		"req_product", // name
		false,         // durable
		false,         // delete when unused
		false,         // exclusive
		false,         // no-wait
		nil,           // arguments
	)
	utils.FailOnError("failed to declare a queue", err)

	body, err := json.Marshal(product_id)
	utils.FailOnError("failed to marshaling message body", err)

	err = ch.Publish(
		"",    
		q.Name,
		false,
		false,
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        body,
		},
	)
	utils.FailOnError("failed to publish a message", err)
	log.Printf("Sent a message: %d", product_id)
}
