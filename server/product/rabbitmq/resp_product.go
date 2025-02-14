package rabbitmq

import (
	"encoding/json"
	"log"
	"github.com/NetSinx/yconnect-shop/server/product/config"
	"github.com/NetSinx/yconnect-shop/server/product/model/entity"
	"github.com/NetSinx/yconnect-shop/server/product/repository"
	"github.com/NetSinx/yconnect-shop/server/product/utils"
	amqp "github.com/rabbitmq/amqp091-go"
)

func ResponseProductByID() {
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	utils.FailOnError("Failed to connect to RabbitMQ", err)
	defer conn.Close()

	ch, err := conn.Channel()
	utils.FailOnError("Failed to open a channel", err)
	defer ch.Close()

	q, err := ch.QueueDeclare(
		"req_product",
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

	var forever chan struct{}

	go func ()  {
		for msg := range msgs {
			log.Printf("Received a message: %s", msg.Body)

			var product_id string

			productRepo := repository.ProductRepository(config.DB)

			if err := json.Unmarshal(msg.Body, &product_id); err != nil {
				log.Print(err)
			}

			product, err := productRepo.GetProductByID(entity.Product{}, product_id)
			if err != nil {
				log.Printf("Product not found")
			}
			
			body, err := json.Marshal(product)
			if err != nil {
				log.Print(err)
			}

			err = ch.Publish(
				"",
				"resp_product",
				false,
				false,
				amqp.Publishing{
					ContentType: "application/json",
					Body: body,
				},
			)
			if err != nil {
				log.Print(err)
			}
			log.Printf("Sent a message: %v", product)
		}
	}()

	log.Printf("[x] Waiting messages... Hold CTRL+C to exit")
	<-forever
}