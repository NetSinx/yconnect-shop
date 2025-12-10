package messaging

import (
	"context"
	"encoding/json"
	"github.com/NetSinx/yconnect-shop/server/product/internal/helpers"
	"github.com/NetSinx/yconnect-shop/server/product/internal/model"
	"github.com/NetSinx/yconnect-shop/server/product/internal/usecase"
	amqp "github.com/rabbitmq/amqp091-go"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type Subscriber struct {
	Connection     *amqp.Connection
	Log            *logrus.Logger
	DB             *gorm.DB
	ProductUseCase *usecase.ProductUseCase
}

func NewSubscriber(connection *amqp.Connection, log *logrus.Logger, db *gorm.DB, productUseCase *usecase.ProductUseCase) *Subscriber {
	return &Subscriber{
		Connection:     connection,
		Log:            log,
		DB:             db,
		ProductUseCase: productUseCase,
	}
}

func (s *Subscriber) Receive() {
	ch, err := s.Connection.Channel()
	helpers.PanicError(s.Log, err, "failed to open a channel")

	exchange := "category_events"
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
		"category_queue",
		true,
		false,
		false,
		false,
		nil,
	)
	helpers.FatalError(s.Log, err, "failed to declare a queue")

	routingKeys := []string{"category.created", "category.updated", "category.deleted"}
	for _, key := range routingKeys {
		err = ch.QueueBind(
			q.Name,
			key,
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
			ctx := context.Background()
			categoryEvent := new(model.CategoryEvent)
			if err := json.Unmarshal(d.Body, categoryEvent); err != nil {
				s.Log.WithError(err).Error("error unmarshaling data")
				break
			}

			switch d.RoutingKey {
			case "category.created":
				s.ProductUseCase.CreateCategoryMirror(ctx, categoryEvent)
			case "category.updated":
				s.ProductUseCase.UpdateCategoryMirror(ctx, categoryEvent)
			case "category.deleted":
				s.ProductUseCase.DeleteCategoryMirror(ctx, categoryEvent)
			}
		}
	}()

	s.Log.Info("Waiting for messages from user service...")
}
