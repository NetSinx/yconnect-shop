package http

import (
	"net/http"
	"github.com/NetSinx/yconnect-shop/server/authentication/handler/dto"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func APIRoutes(e *echo.Echo, authHandler authHandler) {
	apiGroup := e.Group("/api")
	apiGroup.Use(middleware.CSRFWithConfig(middleware.CSRFConfig{
		TokenLookup: "cookie:csrf_token",
		CookieName: "csrf_token",
		CookiePath: "/",
		CookieHTTPOnly: true,
		CookieSecure: true,
		ErrorHandler: func(err error, c echo.Context) error {
			return echo.NewHTTPError(http.StatusBadRequest, dto.MessageResp{
				Message: "csrf token not available",
			})
		},
	}))
	apiGroup.GET("/csrf-token", func(c echo.Context) error {
		return c.JSON(200, dto.MessageResp{
			Message: "CSRF token berhasil di-generate",
		})
	})
	apiGroup.POST("/login", authHandler.LoginUser)
	apiGroup.POST("/logout", authHandler.UserLogout)
	apiGroup.GET("/refresh_token", authHandler.RefreshToken)
}