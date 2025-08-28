package config

import (
	"github.com/NetSinx/yconnect-shop/server/category/internal/delivery/http"
	"github.com/NetSinx/yconnect-shop/server/category/internal/delivery/http/route"
	"github.com/NetSinx/yconnect-shop/server/category/internal/gateway/messaging"
	"github.com/NetSinx/yconnect-shop/server/category/internal/repository"
	"github.com/NetSinx/yconnect-shop/server/category/internal/usecase"
	"github.com/go-playground/validator/v10"
	amqp "github.com/rabbitmq/amqp091-go"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"gorm.io/gorm"
)

type AppBootstrap struct {
	DB        *gorm.DB
	Config    *viper.Viper
	Log       *logrus.Logger
	Validator *validator.Validate
	RabbitMQ  *amqp.Connection
}

func (a *AppBootstrap) NewAppBootstrap(db *gorm.DB, config *viper.Viper, log *logrus.Logger, validator *validator.Validate, rabbitMQ *amqp.Connection) {
	publisher := messaging.NewPublisher(rabbitMQ, log)
	
	repository := repository.NewCategoryRepository(log)
	useCase := usecase.NewCategoryUseCase(db, log, validator, repository, publisher)
	controller := http.NewCategoryController(useCase, log)

	apiRoutes := 
}
