package route

import (
	"github.com/NetSinx/yconnect-shop/server/authentication/internal/delivery/http"
	"github.com/NetSinx/yconnect-shop/server/authentication/internal/delivery/http/middleware"
	"github.com/labstack/echo/v4"
	"github.com/redis/go-redis/v9"
)

type APIRoutes struct {
	AppGroup       *echo.Group
	AuthController *http.AuthController
	RedisClient    *redis.Client
}

func NewAPIRoutes(apiRoutes *APIRoutes) {
	apiGroup := apiRoutes.AppGroup
	apiGroup.Use(middleware.CSRFMiddleware(&middleware.CSRFConfig{
		RedisClient: apiRoutes.RedisClient,
	}))
	apiGroup.GET("/auth/csrf-token", apiRoutes.AuthController.GetCSRFToken)
	apiGroup.POST("/auth/login", apiRoutes.AuthController.LoginUser)
	apiGroup.POST("/auth/verify", apiRoutes.AuthController.Verify)
	apiGroup.POST("/auth/logout", apiRoutes.AuthController.LogoutUser)
}
