package config

import (
	"github.com/NetSinx/yconnect-shop/server/user/internal/delivery/http"
	"github.com/NetSinx/yconnect-shop/server/user/internal/delivery/http/route"
	// subscribeMsg "github.com/NetSinx/yconnect-shop/server/user/internal/delivery/messaging"
	publishMsg "github.com/NetSinx/yconnect-shop/server/user/internal/gateway/messaging"
	"github.com/NetSinx/yconnect-shop/server/user/internal/repository"
	"github.com/NetSinx/yconnect-shop/server/user/internal/usecase"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	amqp "github.com/rabbitmq/amqp091-go"
	"github.com/redis/go-redis/v9"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"gorm.io/gorm"
)

type AppConfig struct {
	Config      *viper.Viper
	DB          *gorm.DB
	Log         *logrus.Logger
	Validator   *validator.Validate
	App         *echo.Echo
	RedisClient *redis.Client
	RabbitMQ    *amqp.Connection
}

func BootstrapApp(config *AppConfig) {
	publisher := publishMsg.NewPublisher(config.RabbitMQ, config.Log)

	repository := repository.NewUserRepository(config.Log)
	usecase := usecase.NewUserUseCase(config.Config, config.DB, config.Log, config.Validator, config.RedisClient, repository, publisher)
	controller := http.NewUserController(config.Log, usecase)

	// subscriber := subscribeMsg.NewSubscriber(config.RabbitMQ, config.Log, config.DB, usecase)
	// subscriber.Receive()

	route.NewApiRoutes(&route.APIRoutes{
		App:            config.App,
		UserController: controller,
	})
}
