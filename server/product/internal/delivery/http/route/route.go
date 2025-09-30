package route

import (
	httpController "github.com/NetSinx/yconnect-shop/server/product/internal/delivery/http"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"net/http"
)

type RouteConfig struct {
	App               *echo.Echo
	ProductController *httpController.ProductController
}

func NewAPIRoutes(routeConfig *RouteConfig) {
	routeConfig.GuestRoutes()
	routeConfig.AdminRoutes()
}

func (r *RouteConfig) GuestRoutes() {
	apiGroup := r.App.Group("/api")
	apiGroup.GET("/products", r.ProductController.GetAllProduct)
	apiGroup.GET("/products/:slug", r.ProductController.GetProductBySlug)
	apiGroup.GET("/products/:slug/category", r.ProductController.GetCategoryProduct)
	apiGroup.GET("/products/category/:slug", r.ProductController.GetProductByCategory)
	apiGroup.Static("/images", "../public")
}

func (r *RouteConfig) AdminRoutes() {
	adminGroup := r.App.Group("/api/admin")
	adminGroup.Use(middleware.CSRFWithConfig(middleware.CSRFConfig{
		TokenLookup:    "header:X-CSRF-Token",
		CookieName:     "csrf_token",
		CookiePath:     "/",
		CookieHTTPOnly: true,
		CookieSecure:   true,
		CookieSameSite: http.SameSiteStrictMode,
	}))
	adminGroup.POST("/products", r.ProductController.CreateProduct)
	adminGroup.PUT("/products/:slug", r.ProductController.UpdateProduct)
	adminGroup.DELETE("/products/:slug", r.ProductController.DeleteProduct)
}
