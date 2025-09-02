package route

import (
	"github.com/NetSinx/yconnect-shop/server/authentication/internal/delivery/http"
	"github.com/NetSinx/yconnect-shop/server/authentication/internal/delivery/http/middleware"
	"github.com/labstack/echo/v4"
)

type APIRoutes struct {
	AppGroup       *echo.Group
	AuthController *http.AuthController
}

func NewAPIRoutes(apiRoutes *APIRoutes) {
	apiGroup := apiRoutes.AppGroup
	apiGroup.Use(middleware.CSRFMiddleware)
	apiGroup.GET("/auth/csrf-token", apiRoutes.AuthController.GetCSRFToken)
	apiGroup.POST("/auth/login", apiRoutes.AuthController.LoginUser)
	apiGroup.POST("/auth/verify", apiRoutes.AuthController.Verify)
	apiGroup.POST("/auth/logout", apiRoutes.AuthController.LogoutUser)
}