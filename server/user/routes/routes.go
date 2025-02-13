package routes

import (
	"fmt"
	"net/http"

	"github.com/NetSinx/yconnect-shop/server/user/config"
	"github.com/NetSinx/yconnect-shop/server/user/controller"
	"github.com/NetSinx/yconnect-shop/server/user/model/domain"
	"github.com/NetSinx/yconnect-shop/server/user/repository"
	"github.com/NetSinx/yconnect-shop/server/user/service"
	"github.com/NetSinx/yconnect-shop/server/user/utils"
	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func ApiRoutes() *echo.Echo {
	userRepository := repository.UserRepository(config.DB)
	userService := service.UserService(userRepository)
	userController := controller.UserController(userService)

	router := echo.New()
	router.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"http://localhost:4200"},
		AllowMethods: []string{"GET", "POST", "PUT", "DELETE"},
	}))
	router.Use(middleware.CSRFWithConfig(middleware.CSRFConfig{
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
	router.GET("/gencsrf", func(c echo.Context) error {
		return c.JSON(200, map[string]interface{}{
			"message": "CSRF token berhasil di-generate",
		})
	})
	router.GET("/user", userController.ListUsers, 
		echojwt.WithConfig(echojwt.Config{
			SigningKey: []byte("yasinnetsinx15"),
			SigningMethod: "HS512",
			TokenLookup: "cookie:user_session",
			ParseTokenFunc: func(c echo.Context, auth string) (interface{}, error) {
				token, err := jwt.ParseWithClaims(auth, &utils.CustomClaims{}, func(t *jwt.Token) (interface{}, error) {
					return []byte("yasinnetsinx15"), nil
				})
				if err != nil {
					return nil, err
				}
				if !token.Valid {
					return nil, fmt.Errorf("your token is invalid")
				}

				claims := token.Claims.(*utils.CustomClaims)
				if claims.Username != "netsinx_15" && claims.Role != "admin" {
					return nil, fmt.Errorf("your claims is invalid")
				}

				return token, nil
			},
		}),
	)
	router.POST("/user/sign-up", userController.RegisterUser)
	router.POST("/user/sign-in", userController.LoginUser)
	router.GET("/user/refresh_token", userController.GetAccessToken)

	authRoute := router.Group("/auth")
	authRoute.Use(echojwt.WithConfig(echojwt.Config{
		SigningKey: []byte("yasinnetsinx15"),
		SigningMethod: "HS512",
		TokenLookup: "cookie:user_session",
		ParseTokenFunc: func(c echo.Context, auth string) (interface{}, error) {
			token, err := jwt.ParseWithClaims(auth, &utils.CustomClaims{}, func(t *jwt.Token) (interface{}, error) {
				return []byte("yasinnetsinx15"), nil
			})
			if err != nil {
				return nil, err
			}
			if !token.Valid {
				return nil, fmt.Errorf("your token is invalid")
			}

			claims := token.Claims.(*utils.CustomClaims)
			if claims.Username == "" && claims.Role == "" {
				return nil, fmt.Errorf("your claims is invalid")
			}

			return token, nil
		},
	}))
	// authRoute.Use(authMiddleware.JWTAuthMiddleware)
	authRoute.POST("/user/send-otp", userController.SendOTP)
	authRoute.POST("/user/email-verify", userController.VerifyEmail)
	authRoute.GET("/user/:username", userController.GetUser)
	authRoute.POST("/user/logout", userController.UserLogout)
	
	return router
}