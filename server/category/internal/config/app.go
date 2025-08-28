package config

import (
	"github.com/NetSinx/yconnect-shop/server/category/internal/delivery/http"
	"github.com/NetSinx/yconnect-shop/server/category/internal/delivery/http/route"
	"github.com/NetSinx/yconnect-shop/server/category/internal/gateway/messaging"
	"github.com/NetSinx/yconnect-shop/server/category/internal/repository"
	"github.com/NetSinx/yconnect-shop/server/category/internal/usecase"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	amqp "github.com/rabbitmq/amqp091-go"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"gorm.io/gorm"
)

type AppBootstrap struct {
	DB        *gorm.DB
	App       *echo.Echo
	Config    *viper.Viper
	Log       *logrus.Logger
	Validator *validator.Validate
	RabbitMQ  *amqp.Connection
}

func NewAppBootstrap(appBootstrap *AppBootstrap) {
	publisher := messaging.NewPublisher(appBootstrap.RabbitMQ, appBootstrap.Log)

	repository := repository.NewCategoryRepository(appBootstrap.Log)
	useCase := usecase.NewCategoryUseCase(appBootstrap.DB, appBootstrap.Log, appBootstrap.Validator, repository, publisher)
	controller := http.NewCategoryController(useCase, appBootstrap.Log)

	route.NewAPIRoutes(&route.APIRoutes{
		App: appBootstrap.App,
		CategoryController: controller,
	})
}