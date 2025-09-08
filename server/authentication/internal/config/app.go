package config

import (
	"github.com/NetSinx/yconnect-shop/server/authentication/internal/delivery/http"
	"github.com/NetSinx/yconnect-shop/server/authentication/internal/delivery/http/route"
	subscribeMsg "github.com/NetSinx/yconnect-shop/server/authentication/internal/delivery/messaging"
	publishMsg "github.com/NetSinx/yconnect-shop/server/authentication/internal/gateway/messaging"
	"github.com/NetSinx/yconnect-shop/server/authentication/internal/helpers"
	"github.com/NetSinx/yconnect-shop/server/authentication/internal/repository"
	"github.com/NetSinx/yconnect-shop/server/authentication/internal/usecase"
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
	tokenUtil := helpers.NewTokenUtil("rahasiadeh", config.RedisClient)
	publisher := publishMsg.NewPublisher(config.RabbitMQ, config.Log)
	
	repository := repository.NewAuthRepository(config.Log)
	usecase := usecase.NewAuthUseCase(config.Config, config.DB, config.Log, config.Validator, publisher, config.RedisClient, repository, tokenUtil)
	controller := http.NewAuthController(config.Log, usecase)


	subscriber := subscribeMsg.NewSubscriber(config.RabbitMQ, config.Log, config.DB, usecase)
	subscriber.Receive()

	route.NewAPIRoutes(&route.APIRoutes{
		AppGroup:       config.App.Group("/api"),
		AuthController: controller,
	})
}
