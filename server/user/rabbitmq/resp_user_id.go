package rabbitmq

import (
	"encoding/json"
	"log"
	"github.com/NetSinx/yconnect-shop/server/user/config"
	"github.com/NetSinx/yconnect-shop/server/user/model/entity"
	"github.com/NetSinx/yconnect-shop/server/user/repository"
	"github.com/NetSinx/yconnect-shop/server/user/utils"
	amqp "github.com/rabbitmq/amqp091-go"
)

func ResponseGetUsernameByID() {
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	utils.FailOnError("Failed to connect to RabbitMQ", err)
	defer conn.Close()

	ch, err := conn.Channel()
	utils.FailOnError("Failed to open a channel", err)
	defer ch.Close()

	q, err := ch.QueueDeclare(
		"req_username",
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

			var username string

			userRepo := repository.UserRepository(config.DB)

			if err := json.Unmarshal(msg.Body, &username); err != nil {
				continue
			}

			user, _ := userRepo.GetUser(entity.User{}, username)
			
			body, err := json.Marshal(user.Id)
			if err != nil {
				continue
			}

			err = ch.Publish(
				"",
				"resp_user_id",
				false,
				false,
				amqp.Publishing{
					ContentType: "text/plain",
					Body: body,
				},
			)
			if err != nil {
				continue
			}
		}
	}()

	log.Printf("[x] Waiting messages... Hold CTRL+C to exit")
	<-forever
}