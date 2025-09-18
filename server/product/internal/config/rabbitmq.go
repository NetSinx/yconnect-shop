package config

import (
	"fmt"
	"github.com/NetSinx/yconnect-shop/server/product/internal/helpers"
	amqp "github.com/rabbitmq/amqp091-go"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func NewRabbitMQ(config *viper.Viper, log *logrus.Logger) *amqp.Connection {
	if !config.GetBool("rabbitmq.enabled") {
		log.Info("RabbitMQ service is disabled")
		return nil
	}

	username := config.GetString("rabbitmq.username")
	password := config.GetString("rabbitmq.password")
	port := config.GetInt("rabbitmq.port")
	host := config.GetString("rabbitmq.host")
	
	url := fmt.Sprintf("amqp://%s:%s@%s:%d/", username, password, host, port)

	connection, err := amqp.Dial(url)
	helpers.PanicError(log, err, "failed to connect rabbitmq")

	return connection
}