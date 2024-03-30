package routes

import (
	"net/http"
	"github.com/NetSinx/yconnect-shop/server/product/app/config"
	"github.com/NetSinx/yconnect-shop/server/product/controller"
	"github.com/NetSinx/yconnect-shop/server/product/repository"
	"github.com/NetSinx/yconnect-shop/server/product/service"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func ApiRoutes() *echo.Echo {
	productRepository := repository.ProductRepository(config.DB)
	productService := service.ProductService(productRepository)
	productController := controller.ProductController(productService)

	router := echo.New()
	router.Use(middleware.CSRFWithConfig(middleware.CSRFConfig{
		TokenLookup: "cookie:xsrf",
		CookiePath: "/",
		CookieHTTPOnly: true,
		CookieSameSite: http.SameSiteStrictMode,
		CookieSecure: true,
	}))

	router.GET("/product", productController.ListProduct)
	router.GET("/product/:id", productController.GetProduct)
	router.GET("/product/category/:id", productController.GetProductByCategory)
	router.POST("/product", productController.CreateProduct)
	router.PUT("/product/:id", productController.UpdateProduct)
	router.DELETE("/product/:id", productController.DeleteProduct)
	router.GET("/product/seller/:id", productController.GetProductBySeller)

	return router
}