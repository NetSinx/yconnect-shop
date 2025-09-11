package messaging

import (
	"context"
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

	exchange := "user_auth_events"
	err = ch.ExchangeDeclare(
		exchange,
		"direct",
		true,
		false,
		false,
		false,
		nil,
	)
	helpers.FatalError(s.Log, err, "failed to declare an exchange")

	q, err := ch.QueueDeclare(
		"user_deleted_queue",
		true,
		false,
		false,
		false,
		nil,
	)
	helpers.FatalError(s.Log, err, "failed to declare a queue")

	err = ch.QueueBind(
		q.Name,
		"user_deleted",
		exchange,
		false,
		nil,
	)
	helpers.FatalError(s.Log, err, "failed to bind a queue")

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
			s.Log.Infof("Receive a message from user service: %s", d.Body)

			ctx := context.Background()
			var deleteUserEvent *model.DeleteUserEvent
			if err := json.Unmarshal(d.Body, &deleteUserEvent); err != nil {
				s.Log.WithError(err).Error(s.Log, err, "error unmarshaling message body")
				continue
			}

			if err := s.AuthUseCase.DeleteUserAuthentication(ctx, deleteUserEvent); err != nil {
				s.Log.WithError(err).Error("error deleting user authentication")
			}
		}
	}()
	
	s.Log.Info("Waiting for messages from user service...")
}
