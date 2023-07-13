package routes

import (
	"net/http"
	"github.com/NetSinx/yconnect-shop/cart/config"
	"github.com/NetSinx/yconnect-shop/cart/controller"
	auth "github.com/NetSinx/yconnect-shop/cart/middleware"
	"github.com/NetSinx/yconnect-shop/cart/repository"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func ApiRoutes() *echo.Echo {
	cartRepository := repository.CartRepository(config.DB)
	cartController := controller.CartController(cartRepository)

	router := echo.New()
	router.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"http://localhost:4200"},
		AllowMethods: []string{http.MethodGet, http.MethodPost, http.MethodPut, http.MethodDelete},
	}),
	auth.AuthMiddleware(),
	)

	router.GET("/carts", cartController.ListCart)
	router.POST("/carts", cartController.CreateCart)
	router.PUT("/carts/:id", cartController.UpdateCart)
	router.DELETE("/carts/:id", cartController.DeleteCart)
	router.GET("/carts/id/:id", cartController.GetCartById)
	router.GET("/carts/slug/:slug", cartController.GetCartBySlug)

	return router
}