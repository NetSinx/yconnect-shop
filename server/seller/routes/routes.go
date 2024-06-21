package routes

import (
	"github.com/NetSinx/yconnect-shop/server/seller/config"
	"github.com/NetSinx/yconnect-shop/server/seller/controller"
	"github.com/NetSinx/yconnect-shop/server/seller/repository"
	"github.com/NetSinx/yconnect-shop/server/seller/service"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"net/http"
)

func APIRoutes() *echo.Echo {
	sellerRepository := repository.SellerRepository(config.ConfigDB())
	sellerService := service.SellerService(sellerRepository)
	sellerController := controller.SellerController(sellerService)

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
	router.GET("/seller", sellerController.ListSeller)
	router.POST("/seller/:username", sellerController.RegisterSeller)
	router.PUT("/seller/:username", sellerController.UpdateSeller)
	router.DELETE("/seller/:username", sellerController.DeleteSeller)
	router.GET("/seller/:username", sellerController.GetSeller)

	return router
}