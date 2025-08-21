package http

import (
	"github.com/NetSinx/yconnect-shop/server/authentication/middleware"
	"github.com/labstack/echo/v4"
)

func APIRoutes(e *echo.Echo, authHandler authHandler) {
	apiGroup := e.Group("/api")
	apiGroup.Use(middleware.CSRFMiddleware)
	apiGroup.GET("/gencsrf", func(c echo.Context) error {
		return c.JSON(200, map[string]any{
			"message": "CSRF token berhasil di-generate",
		})
	})
	apiGroup.POST("/login", authHandler.LoginUser)
	apiGroup.POST("/logout", authHandler.UserLogout)
	apiGroup.GET("/refresh_token", authHandler.RefreshToken)
}