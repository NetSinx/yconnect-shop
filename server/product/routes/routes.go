package routes

import (
	"net/http"
	"github.com/NetSinx/yconnect-shop/server/product/config"
	"github.com/NetSinx/yconnect-shop/server/product/controller"
	authMiddleware "github.com/NetSinx/yconnect-shop/server/product/middleware"
	"github.com/NetSinx/yconnect-shop/server/product/repository"
	"github.com/NetSinx/yconnect-shop/server/product/service"
	"github.com/NetSinx/yconnect-shop/server/product/utils"
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func ApiRoutes() *echo.Echo {
	productRepository := repository.ProductRepository(config.ConnectDB())
	productService := service.ProductService(productRepository)
	productController := controller.ProductController(productService)

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
	router.GET("/product", productController.ListProduct)
	router.GET("/product/:id", productController.GetProductByID)
	router.GET("/product/:slug", productController.GetProductBySlug)
	router.GET("/product/:slug/category", productController.GetCategoryProduct)

	authRoute := router.Group("/auth")
	authRoute.Use(echojwt.WithConfig(echojwt.Config{
		SigningKey: []byte(utils.AdminJwtKey),
		SigningMethod: "HS512",
	}))
	authRoute.Use(authMiddleware.JWTAuthMiddleware)
	authRoute.POST("/product", productController.CreateProduct)
	authRoute.PUT("/product/:slug", productController.UpdateProduct)
	authRoute.DELETE("/product/:slug", productController.DeleteProduct)

	return router
}