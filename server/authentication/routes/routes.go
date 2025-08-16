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
	apiRouter := router.Group("/api")
	apiRouter.Use(middleware.CSRFWithConfig(middleware.CSRFConfig{
		TokenLookup: "cookie:_csrf",
		CookiePath: "/",
		CookieHTTPOnly: true,
		CookieSameSite: http.SameSiteStrictMode,
		CookieMaxAge: 60,
		CookieSecure: true,
		ErrorHandler: func(err error, c echo.Context) error {
			return echo.NewHTTPError(http.StatusBadRequest, domain.MessageResp{
				Message: "csrf token not available",
			})
		},
	}))
	apiRouter.GET("/gencsrf", func(c echo.Context) error {
		return c.JSON(200, map[string]interface{}{
			"message": "CSRF token berhasil di-generate",
		})
	})
	apiRouter.POST("/login", authController.LoginUser)
	apiRouter.POST("/logout", authController.UserLogout)
	apiRouter.GET("/refresh_token", authController.RefreshToken)

	return router
}