package routes

import (
	"net/http"
	"github.com/NetSinx/yconnect-shop/server/product/config"
	"github.com/NetSinx/yconnect-shop/server/product/controller"
	"github.com/NetSinx/yconnect-shop/server/product/model/domain"
	"github.com/NetSinx/yconnect-shop/server/product/repository"
	"github.com/NetSinx/yconnect-shop/server/product/service"
	"github.com/NetSinx/yconnect-shop/server/authentication/utils"
	"github.com/golang-jwt/jwt/v5"
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func ApiRoutes() *echo.Echo {
	productRepository := repository.ProductRepository(config.ConnectDB())
	productService := service.ProductService(productRepository)
	productController := controller.ProductController(productService)

	router := echo.New()
	apiGroup := router.Group("/api")
	apiGroup.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"http://localhost:4200"},
		AllowMethods: []string{"GET", "POST", "PUT", "DELETE"},
	}))
	apiGroup.Use(middleware.CSRFWithConfig(middleware.CSRFConfig{
		TokenLookup: "cookie:csrf_token",
		CookieName: "csrf_token",
		CookiePath: "/",
		CookieHTTPOnly: true,
		CookieSameSite: http.SameSiteStrictMode,
		CookieMaxAge: 60,
		CookieSecure: true,
		ErrorHandler: func(err error, c echo.Context) error {
			return echo.NewHTTPError(http.StatusBadRequest, domain.MessageResp{
				Message: "csrf token not available",
			})
		},
	}))
	apiGroup.GET("/product", productController.ListProduct)
	apiGroup.GET("/product/id/:id", productController.GetProductByID)
	apiGroup.GET("/product/slug/:slug", productController.GetProductBySlug)
	apiGroup.GET("/product/:slug/category", productController.GetCategoryProduct)

	adminGroup := apiGroup.Group("/admin")
	adminGroup.Use(echojwt.WithConfig(echojwt.Config{
		SigningKey: []byte(utils.AdminJwtKey),
		SigningMethod: "HS512",
		NewClaimsFunc: func(c echo.Context) jwt.Claims {
			return new(utils.CustomClaims)
		},
		ErrorHandler: func(c echo.Context, err error) error {
			return echo.ErrUnauthorized
		},
	}))
	adminGroup.Use(utils.CheckAdminRole)
	adminGroup.POST("/product", productController.CreateProduct)
	adminGroup.PUT("/product/:slug", productController.UpdateProduct)
	adminGroup.DELETE("/product/:slug", productController.DeleteProduct)

	return router
}