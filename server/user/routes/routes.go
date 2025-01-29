package routes

import (
	"net/http"
	"github.com/NetSinx/yconnect-shop/server/user/config"
	"github.com/NetSinx/yconnect-shop/server/user/controller"
	"github.com/NetSinx/yconnect-shop/server/user/repository"
	"github.com/NetSinx/yconnect-shop/server/user/service"
	"github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	authMiddleware "github.com/NetSinx/yconnect-shop/server/user/middleware"
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
		CookieMaxAge: 60,
		CookieSecure: true,
	}))
	router.Use(middleware.KeyAuthWithConfig(middleware.KeyAuthConfig{
		KeyLookup: "header:api-token",
		Validator: func(auth string, c echo.Context) (bool, error) {
			return auth == "dfkgjdgj#753846873248358645*&#%^*$54%hgdf", nil
		},
	}))
	router.GET("/gencsrf", func(c echo.Context) error {
		csrfToken := c.Get("csrf")

		return c.JSON(200, map[string]interface{}{
			"csrf_token": csrfToken,
		})
	})
	router.POST("/user/sign-up", userController.RegisterUser)
	router.POST("/user/sign-in", userController.LoginUser)

	authRoute := router.Group("/auth")
	authRoute.Use(echojwt.WithConfig(echojwt.Config{
		SigningKey: []byte("yasinnetsinx15"),
		SigningMethod: "HS512",
	}))
	authRoute.Use(authMiddleware.JWTAuthMiddleware)
	authRoute.POST("/user/send-otp", userController.SendOTP)
	authRoute.POST("/user/email-verify", userController.VerifyEmail)
	authRoute.POST("/user/set-timezone", userController.SetTimezone)
	authRoute.GET("/user/userinfo", userController.GetUserInfo)
	authRoute.GET("/user/:username", userController.GetUser)
	authRoute.GET("/user/logout", userController.UserLogout)
	
	adminRoute := router.Group("/admin/")
	adminRoute.GET("/user", userController.ListUsers)
	adminRoute.PUT("/user/:username", userController.UpdateUser)
	adminRoute.DELETE("/user/:username", userController.DeleteUser)
	adminRoute.POST("/user/set-timezone", userController.SetTimezone)
	adminRoute.GET("/user/userinfo", userController.GetUserInfo)
	adminRoute.GET("/user/:username", userController.GetUser)
	adminRoute.GET("/user/logout", userController.UserLogout)
	
	return router
}