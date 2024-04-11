package routes

import (
	"github.com/NetSinx/yconnect-shop/server/seller/config"
	"github.com/NetSinx/yconnect-shop/server/seller/controller"
	"github.com/NetSinx/yconnect-shop/server/seller/repository"
	"github.com/NetSinx/yconnect-shop/server/seller/service"
	"github.com/labstack/echo/v4"
)

func APIRoutes() *echo.Echo {
	sellerRepository := repository.SellerRepository(config.ConfigDB())
	sellerService := service.SellerService(sellerRepository)
	sellerController := controller.SellerController(sellerService)

	router := echo.New()

	router.GET("/seller", sellerController.ListSeller)
	router.POST("/seller/:username", sellerController.RegisterSeller)
	router.PUT("/seller/:username", sellerController.UpdateSeller)
	router.DELETE("/seller/:username", sellerController.DeleteSeller)
	router.GET("/seller/:username", sellerController.GetSeller)

	return router
}