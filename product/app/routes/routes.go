package routes

import (
	"net/http"
	"github.com/NetSinx/yconnect-shop/product/app/config"
	"github.com/NetSinx/yconnect-shop/product/controller"
	"github.com/NetSinx/yconnect-shop/product/repository"
	"github.com/NetSinx/yconnect-shop/product/service"
	auth "github.com/NetSinx/yconnect-shop/product/app/middleware"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func ApiRoutes() *echo.Echo {
	productRepository := repository.ProductRepository(config.DB)
	productService := service.ProductService(productRepository)
	productController := controller.ProductController(productService)

	router := echo.New()
	router.Use(middleware.CORSWithConfig(middleware.CORSConfig{
			AllowOrigins: []string{"http://localhost:4200"},
			AllowMethods: []string{http.MethodGet, http.MethodPost, http.MethodPut, http.MethodDelete},
		}),
	)
	router.GET("/products", productController.ListProduct)
	router.GET("/products/:slug", productController.GetProduct)
	router.GET("/products/category/:id", productController.GetProductByCategory)

	routerAuth := router.Group("/api", auth.AuthMiddleware)
	routerAuth.POST("/products", productController.CreateProduct)
	routerAuth.PUT("/products/:slug", productController.UpdateProduct)
	routerAuth.DELETE("/products/:slug", productController.DeleteProduct)
	routerAuth.GET("/products/user/:id", productController.GetProductByUser)

	return router
}