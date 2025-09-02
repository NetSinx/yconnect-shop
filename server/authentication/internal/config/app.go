package config

import (
	"github.com/NetSinx/yconnect-shop/server/authentication/internal/delivery/http"
	"github.com/NetSinx/yconnect-shop/server/authentication/internal/delivery/http/route"
	"github.com/NetSinx/yconnect-shop/server/authentication/internal/helpers"
	"github.com/NetSinx/yconnect-shop/server/authentication/internal/repository"
	"github.com/NetSinx/yconnect-shop/server/authentication/internal/usecase"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
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
	TokenUtil   *helpers.TokenUtil
}

func BootstrapApp(config *AppConfig) {
	repository := repository.NewAuthRepository(config.Log)
	usecase := usecase.NewAuthUseCase(config.DB, config.Log, config.Validator, config.RedisClient, repository, config.TokenUtil)
	controller := http.NewAuthController(config.Log, usecase)

	route.NewAPIRoutes(&route.APIRoutes{
		AppGroup: config.App.Group("/api"),
		AuthController: controller,
		RedisClient: config.RedisClient,
	})
}
