package messaging

import (
	"context"

	"github.com/NetSinx/yconnect-shop/server/user/internal/helpers"
	"github.com/NetSinx/yconnect-shop/server/user/internal/model"
	amqp "github.com/rabbitmq/amqp091-go"
	"github.com/sirupsen/logrus"
)

type Publisher struct {
	RabbitMQ *amqp.Connection
	Log      *logrus.Logger
}

func NewPublisher(rabbitmq *amqp.Connection, log *logrus.Logger) *Publisher {
	return &Publisher{
		RabbitMQ: rabbitmq,
		Log: log,
	}
}

func (p *Publisher) Send(ctx context.Context, message *model.DeleteUserEvent) {
	ch, err := p.RabbitMQ.Channel()
	helpers.PanicError(p.Log, err, "failed to open a channel")
	defer ch.Close()
}
