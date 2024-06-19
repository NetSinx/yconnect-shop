package routes

import (
	"net/http"

	"github.com/NetSinx/yconnect-shop/server/cart/config"
	"github.com/NetSinx/yconnect-shop/server/cart/controller"
	"github.com/NetSinx/yconnect-shop/server/cart/repository"
	"github.com/NetSinx/yconnect-shop/server/cart/service"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func ApiRoutes() *echo.Echo {
	cartRepository := repository.CartRepository(config.DB)
	cartService := service.CartService(cartRepository)
	cartController := controller.CartController(cartService)

	router := echo.New()
	router.Use(middleware.CSRFWithConfig(middleware.CSRFConfig{
		TokenLookup: "header:xsrf",
		CookiePath: "/",
		CookieName: "xsrf",
		CookieHTTPOnly: true,
		CookieSameSite: http.SameSiteStrictMode,
		CookieSecure: true,
		CookieMaxAge: 60,
	}))

	router.GET("/cart/gencsrf", func(c echo.Context) error {
		csrf := c.Get("csrf")

		return c.JSON(http.StatusOK, map[string]interface{}{
			"csrf_token": csrf,
		})
	})
	router.GET("/cart", cartController.ListCart)
	router.POST("/cart/:id", cartController.AddToCart)
	router.PUT("/cart/:id", cartController.UpdateCart)
	router.DELETE("/cart/:id", cartController.DeleteProductInCart)
	router.GET("/cart/:id", cartController.GetCart)
	router.GET("/cart/user/:id", cartController.GetCartByUser)

	return router
}