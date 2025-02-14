package rabbitmq

import (
	"encoding/json"
	"github.com/NetSinx/yconnect-shop/server/order/utils"
	amqp "github.com/rabbitmq/amqp091-go"
	productEntity "github.com/NetSinx/yconnect-shop/server/product/model/entity"
)

func ConsumeProduct() (productEntity.Product, error) {
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	utils.FailOnError("failed to connect to RabbitMQ", err)
	defer conn.Close()

	ch, err := conn.Channel()
	utils.FailOnError("failed to open a channel", err)
	defer ch.Close()

	q, err := ch.QueueDeclare(
		"resp_product",
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
		var respProduct productEntity.Product
		json.Unmarshal(msg.Body, &respProduct)
		return respProduct, nil
	}
	
	return productEntity.Product{}, nil
}
