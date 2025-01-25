package routes

import (
	"net/http"
	"github.com/NetSinx/yconnect-shop/server/user/config"
	"github.com/NetSinx/yconnect-shop/server/user/controller"
	"github.com/NetSinx/yconnect-shop/server/user/repository"
	"github.com/NetSinx/yconnect-shop/server/user/service"
	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/NetSinx/yconnect-shop/server/user/utils"
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
		ErrorHandler: func(c echo.Context, err error) error {
			authToken, _ := c.Cookie("Authorization")
			token, err := jwt.Parse(authToken.Value, func(t *jwt.Token) (interface{}, error) {
				return []byte("yasinnetsinx15"), nil
			})
			if err != nil {
				return echo.NewHTTPError(http.StatusUnauthorized, map[string]string{
					"message": "Your token is invalid",
				})
			}

			claims := token.Claims.(*utils.CustomClaims)
			if claims.Username == "" && claims.Role == "" {
				return echo.NewHTTPError(http.StatusUnauthorized, map[string]string{
					"message": "Your claims is invalid",
				})
			}

			return nil
		},
	}))
	authRoute.POST("/user/send-otp", userController.SendOTP)
	authRoute.POST("/user/email-verify", userController.VerifyEmail)
	authRoute.POST("/user/set-timezone", userController.SetTimezone)
	authRoute.GET("/user/userinfo", userController.GetUserInfo)
	authRoute.GET("/user/:username", userController.GetUser)
	authRoute.GET("/user/logout", userController.UserLogout)
	
	authRoute.GET("/user", userController.ListUsers)
	authRoute.PUT("/user/:username", userController.UpdateUser)
	authRoute.DELETE("/user/:username", userController.DeleteUser)
	
	return router
}