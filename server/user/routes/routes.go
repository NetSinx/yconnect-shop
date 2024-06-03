package routes

import (
	"net/http"
	"github.com/NetSinx/yconnect-shop/server/user/config"
	"github.com/NetSinx/yconnect-shop/server/user/controller"
	"github.com/NetSinx/yconnect-shop/server/user/repository"
	"github.com/NetSinx/yconnect-shop/server/user/service"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func ApiRoutes() *echo.Echo {
	userRepository := repository.UserRepository(config.ConnectDB())
	userService := service.UserService(userRepository)
	userController := controller.UserController(userService)

	router := echo.New()
	router.Use(middleware.CSRFWithConfig(middleware.CSRFConfig{
		TokenLookup: "header:xsrf",
		CookieName: "xsrf",
		CookiePath: "/",
		CookieHTTPOnly: true,
		CookieSameSite: http.SameSiteStrictMode,
		CookieMaxAge: 120,
		CookieSecure: true,
	}))

	router.GET("/user/gencsrf", func(c echo.Context) error {
		csrfToken := c.Get("csrf")

		return c.JSON(200, map[string]interface{}{
			"csrf_token": csrfToken,
		})
	})
	router.POST("/user/sign-up", userController.RegisterUser)
	router.POST("/user/sign-in", userController.LoginUser)
	router.POST("/user/send-otp", userController.SendOTP)
	router.POST("/user/email-verify", userController.VerifyEmail)
	router.GET("/user", userController.ListUsers)
	router.GET("/user/:username", userController.GetUser)
	router.PUT("/user/:username", userController.UpdateUser)
	router.DELETE("/user/:username", userController.DeleteUser)

	return router
}