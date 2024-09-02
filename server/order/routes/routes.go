package routes

import (
	// "net/http"
	"github.com/NetSinx/yconnect-shop/server/order/config"
	"github.com/NetSinx/yconnect-shop/server/order/controller"
	"github.com/NetSinx/yconnect-shop/server/order/repository"
	"github.com/NetSinx/yconnect-shop/server/order/service"
	"github.com/labstack/echo/v4"
	// "github.com/labstack/echo/v4/middleware"
)

func RoutesAPI() *echo.Echo {
	orderRepository := repository.OrderRepo(config.ConnectDB())
	orderService := service.OrderServ(orderRepository)
	orderController := controller.OrderContrllr(orderService)

	router := echo.New()
	// router.Use(middleware.CSRFWithConfig(middleware.CSRFConfig{
	// 	TokenLookup: "header:xsrf",
	// 	CookieName: "xsrf",
	// 	CookiePath: "/",
	// 	CookieMaxAge: 60,
	// 	CookieHTTPOnly: true,
	// 	CookieSameSite: http.SameSiteStrictMode,
	// 	CookieSecure: true,
	// }))
	router.GET("/order/:username", orderController.ListOrder)
	router.POST("/order", orderController.AddOrder)
	router.DELETE("/order/:username/:id", orderController.DeleteOrder)

	return router
}