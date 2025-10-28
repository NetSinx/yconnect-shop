package rabbitmq

import (
	"encoding/json"
	"log"
	"github.com/NetSinx/yconnect-shop/server/cart/config"
	"github.com/NetSinx/yconnect-shop/server/cart/model/entity"
	"github.com/NetSinx/yconnect-shop/server/cart/repository"
	"github.com/NetSinx/yconnect-shop/server/cart/utils"
	amqp "github.com/rabbitmq/amqp091-go"
)

func ResponseGetCartByUserID() {
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	utils.FailOnError("Failed to connect to RabbitMQ", err)
	defer conn.Close()

	ch, err := conn.Channel()
	utils.FailOnError("Failed to open a channel", err)
	defer ch.Close()

	q, err := ch.QueueDeclare(
		"req_cart",
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

	go func()  {
		for msg := range msgs {
			log.Printf("Received a message: %s", msg.Body)

			var id uint

			cartRepo := repository.CartRepository(config.DB)

			if err := json.Unmarshal(msg.Body, &id); err != nil {
				continue
			}

			cart, _ := cartRepo.GetCartByUser([]entity.Cart{}, id)
			
			body, err := json.Marshal(cart)
			if err != nil {
				continue
			}

			err = ch.Publish(
				"",
				"resp_cart",
				false,
				false,
				amqp.Publishing{
					ContentType: "application/json",
					Body: body,
				},
			)
			if err != nil {
				continue
			}
			log.Printf("Sent a response: %v", cart)
		}
	}()

	log.Printf("[x] Waiting messages... Hold CTRL+C to exit")
	<-forever
}
