package config

import (
	"github.com/NetSinx/yconnect-shop/server/product/internal/delivery/http"
	"github.com/NetSinx/yconnect-shop/server/product/internal/delivery/http/route"
	"github.com/NetSinx/yconnect-shop/server/product/internal/gateway/messaging"
	"github.com/NetSinx/yconnect-shop/server/product/internal/repository"
	"github.com/NetSinx/yconnect-shop/server/product/internal/usecase"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	amqp "github.com/rabbitmq/amqp091-go"
	"github.com/redis/go-redis/v9"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"gorm.io/gorm"
)

type AppBootstrap struct {
	DB          *gorm.DB
	App         *echo.Echo
	Config      *viper.Viper
	Log         *logrus.Logger
	RedisClient *redis.Client
	Validator   *validator.Validate
	RabbitMQ    *amqp.Connection
}

func NewAppBootstrap(appBootstrap *AppBootstrap) {
	publisher := messaging.NewPublisher(appBootstrap.RabbitMQ, appBootstrap.Log)

	repository := repository.NewProductRepository(appBootstrap.Log)
	useCase := usecase.NewProductUseCase(appBootstrap.Config, appBootstrap.DB, appBootstrap.Log, appBootstrap.Validator, appBootstrap.RedisClient, publisher, repository)
	controller := http.NewProductController(appBootstrap.Log, useCase)

	route.NewAPIRoutes(&route.RouteConfig{
		App:               appBootstrap.App,
		ProductController: controller,
	})
}
