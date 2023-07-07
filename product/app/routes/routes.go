package routes

import (
	"net/http"
	"github.com/NetSinx/yconnect-shop/product/app/config"
	authMiddleware "github.com/NetSinx/yconnect-shop/product/app/middleware"
	"github.com/NetSinx/yconnect-shop/product/controller"
	"github.com/NetSinx/yconnect-shop/product/repository"
	"github.com/NetSinx/yconnect-shop/product/service"
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
			AllowMethods: []string{http.MethodGet},
		}),
	)
	router.GET("/products", productController.ListProduct)
	router.GET("/products/:slug", productController.GetProduct)
	router.GET("/products/category/:id", productController.GetProductByCategory)

	router.Group("/api", middleware.CORSWithConfig(middleware.CORSConfig{
			AllowOrigins: []string{"http://localhost:4200"},
			AllowMethods: []string{http.MethodPost, http.MethodPut, http.MethodDelete},
		}), 
		authMiddleware.AuthMiddleware,
	)
	router.POST("/api/products", productController.CreateProduct)
	router.PUT("/api/products/:slug", productController.UpdateProduct)
	router.DELETE("/api/products/:slug", productController.DeleteProduct)
	router.GET("/api/products/user/:id", productController.GetProductByUser)

	return router
}