package routes

import (
	"net/http"
	"github.com/NetSinx/yconnect-shop/server/order/config"
	"github.com/NetSinx/yconnect-shop/server/order/controller"
	"github.com/NetSinx/yconnect-shop/server/order/repository"
	"github.com/NetSinx/yconnect-shop/server/order/service"
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	// authMiddleware "github.com/NetSinx/yconnect-shop/server/order/middleware"
)

func RoutesAPI() *echo.Echo {
	orderRepository := repository.OrderRepo(config.DB)
	orderService := service.OrderServ(orderRepository)
	orderController := controller.OrderContrllr(orderService)

	router := echo.New()
	router.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"http://localhost:4200"},
		AllowMethods: []string{"GET", "POST", "PUT", "DELETE"},
	}))
	router.Use(middleware.CSRFWithConfig(middleware.CSRFConfig{
		TokenLookup: "header:xsrf",
		CookieName: "xsrf",
		CookiePath: "/",
		CookieMaxAge: 60,
		CookieHTTPOnly: true,
		CookieSameSite: http.SameSiteStrictMode,
		CookieSecure: true,
	}))
	router.Use(echojwt.WithConfig(echojwt.Config{
		SigningKey: []byte("yasinnetsinx15"),
		SigningMethod: "HS512",
	}))
	// router.Use(authMiddleware.JWTAuthMiddleware)
	router.GET("/order", orderController.GetOrder)
	router.POST("/order", orderController.AddOrder)
	router.DELETE("/order", orderController.DeleteOrder)

	return router
}