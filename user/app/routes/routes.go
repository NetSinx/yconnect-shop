package routes

import (
	"net/http"
	"github.com/NetSinx/yconnect-shop/user/app/config"
	auth "github.com/NetSinx/yconnect-shop/user/app/middleware"
	"github.com/NetSinx/yconnect-shop/user/controller"
	"github.com/NetSinx/yconnect-shop/user/repository"
	"github.com/NetSinx/yconnect-shop/user/service"
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
			AllowMethods: []string{http.MethodGet, http.MethodPost, http.MethodPut, http.MethodDelete},
		}),
	)
	router.POST("/users/sign-up", userController.RegisterUser)
	router.POST("/users/sign-in", userController.LoginUser)
	
	routerAuth := router.Group("/api", auth.AuthMiddleware)
	routerAuth.GET("/users", userController.ListUsers)
	routerAuth.GET("/users/:id", userController.GetUser)

	return router
}