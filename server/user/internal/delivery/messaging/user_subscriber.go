package messaging

import (
	"context"
	"encoding/json"
	"github.com/NetSinx/yconnect-shop/server/user/internal/helpers"
	"github.com/NetSinx/yconnect-shop/server/user/internal/model"
	"github.com/NetSinx/yconnect-shop/server/user/internal/usecase"
	amqp "github.com/rabbitmq/amqp091-go"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type Subscriber struct {
	RabbitMQ    *amqp.Connection
	Log         *logrus.Logger
	DB          *gorm.DB
	UserUseCase *usecase.UserUseCase
}

func NewSubscriber(rabbitmq *amqp.Connection, Log *logrus.Logger, db *gorm.DB, userUseCase *usecase.UserUseCase) *Subscriber {
	return &Subscriber{
		RabbitMQ:    rabbitmq,
		Log:         Log,
		DB:          db,
		UserUseCase: userUseCase,
	}
}

func (s *Subscriber) Receive() {
	ch, err := s.RabbitMQ.Channel()
	helpers.PanicError(s.Log, err, "failed to open a channel")
	defer ch.Close()

	exchange := "user_data_events"
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
		"user_data_queue",
		true,
		false,
		false,
		false,
		nil,
	)
	helpers.FatalError(s.Log, err, "failed to declare a queue")

	err = ch.QueueBind(
		q.Name,
		"user_register",
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
	helpers.FatalError(s.Log, err, "failed to consume messages")

	go func() {
		for d := range msgs {
			s.Log.Infof("Receive a message from authentication service: %s", d.Body)

			var ctx context.Context
			var userEvent *model.RegisterUserEvent
			if err := json.Unmarshal(d.Body, &userEvent); err != nil {
				s.Log.WithError(err).Error("error unmarshaling message body")
				continue
			}

			if err := s.UserUseCase.RegisterUser(ctx, userEvent); err != nil {
				s.Log.WithError(err).Error("error registering user")
			}
		}
	}()

	s.Log.Info("Waiting for messages from authentication service...")
}
