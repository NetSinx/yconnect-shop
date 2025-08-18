package routes

import (
	"net/http"
	"github.com/NetSinx/yconnect-shop/server/authentication/controller"
	"github.com/NetSinx/yconnect-shop/server/authentication/domain"
	"github.com/NetSinx/yconnect-shop/server/authentication/service"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func APIRoutes() *echo.Echo {
	authService := service.AuthServ()
	authController := controller.AuthContrllr(authService)

	router := echo.New()
	apiGroup := router.Group("/api")
	apiGroup.Use(middleware.CSRFWithConfig(middleware.CSRFConfig{
		TokenLookup: "cookie:csrf_token",
		CookieName: "csrf_token",
		CookiePath: "/",
		CookieMaxAge: 30,
		CookieHTTPOnly: true,
		CookieSecure: true,
		ErrorHandler: func(err error, c echo.Context) error {
			return echo.NewHTTPError(http.StatusBadRequest, domain.MessageResp{
				Message: "csrf token not available",
			})
		},
	}))
	apiGroup.GET("/gencsrf", func(c echo.Context) error {
		return c.JSON(200, map[string]any{
			"message": "CSRF token berhasil di-generate",
		})
	})
	apiGroup.POST("/login", authController.LoginUser)
	apiGroup.POST("/logout", authController.UserLogout)
	apiGroup.GET("/refresh_token", authController.RefreshToken)

	return router
}