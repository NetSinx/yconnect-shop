package http

import (
	"time"
	"github.com/NetSinx/yconnect-shop/server/authentication/middleware"
	"github.com/labstack/echo/v4"
)

func APIRoutes(e *echo.Echo, authHandler authHandler) {
	apiGroup := e.Group("/api")
	
	csrfManager := middleware.NewCSRFManager(60 * time.Second)
	
	apiGroup.GET("/csrf-token", middleware.GetCSRFTokenHandler(csrfManager))
	apiGroup.POST("/login", authHandler.LoginUser)
	apiGroup.POST("/logout", authHandler.UserLogout)
	apiGroup.GET("/refresh_token", authHandler.RefreshToken)
}