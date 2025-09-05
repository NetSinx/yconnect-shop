package messaging

import (
	"encoding/json"

	"github.com/NetSinx/yconnect-shop/server/authentication/internal/helpers"
	"github.com/NetSinx/yconnect-shop/server/authentication/internal/model"
	"github.com/NetSinx/yconnect-shop/server/authentication/internal/usecase"
	amqp "github.com/rabbitmq/amqp091-go"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type Subscriber struct {
	Connection  *amqp.Connection
	Log         *logrus.Logger
	DB          *gorm.DB
	AuthUseCase *usecase.AuthUseCase
}

func NewSubscriber(connection *amqp.Connection, log *logrus.Logger, db *gorm.DB, authUseCase *usecase.AuthUseCase) *Subscriber {
	return &Subscriber{
		Connection:  connection,
		Log:         log,
		DB:          db,
		AuthUseCase: authUseCase,
	}
}

func (s *Subscriber) Receive() {
	ch, err := s.Connection.Channel()
	helpers.PanicError(s.Log, err, "failed to open a channel")
	defer ch.Close()

	exchange := "user_service_events"
	err = ch.ExchangeDeclare(
		exchange,
		"topic",
		true,
		false,
		false,
		false,
		nil,
	)
	helpers.FatalError(s.Log, err, "failed to declare an exchange")

	q, err := ch.QueueDeclare(
		"user_auth_mirror",
		true,
		false,
		false,
		false,
		nil,
	)
	helpers.FatalError(s.Log, err, "failed to declare a queue")

	routingKey := []string{"user.created", "user.updated", "user.deleted"}
	for _, k := range routingKey {
		err = ch.QueueBind(
			q.Name,
			k,
			exchange,
			false,
			nil,
		)
		helpers.FatalError(s.Log, err, "failed to bind a queue")
	}

	msgs, err := ch.Consume(
		q.Name,
		"",
		false,
		false,
		false,
		false,
		nil,
	)
	helpers.FatalError(s.Log, err, "failed to register a consumer")

	go func() {
		for d := range msgs {
			var userEvent *model.UserEvent
			if err := json.Unmarshal(d.Body, userEvent); err != nil {
				helpers.FatalError(s.Log, err, "error unmarshalling message body")
			}

			switch d.RoutingKey {
			case "user.created":
				if err := s.AuthUseCase.Create(userEvent); err != nil {
					s.Log.WithError(err).Error("error creating user data miror")
				}
			case "user.updated":
				if err := s.AuthUseCase.Update(userEvent); err != nil {
					s.Log.WithError(err).Error("error updating user data miror")
				}
			case "user.deleted":
				if err := s.AuthUseCase.Delete(userEvent); err != nil {
					s.Log.WithError(err).Error("error deleting user data miror")
				}
			default:
			}
		}
	}()
}
