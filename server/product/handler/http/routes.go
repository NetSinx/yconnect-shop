package http

import (
	"net/http"
	"github.com/NetSinx/yconnect-shop/server/product/handler/dto"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func ApiRoutes(e *echo.Echo, productHandler *productHandler) {
	apiGroup := e.Group("/api")
	apiGroup.GET("/products", productHandler.ListProduct)
	apiGroup.GET("/products/id/:id", productHandler.GetProductByID)
	apiGroup.GET("/products/slug/:slug", productHandler.GetProductBySlug)
	apiGroup.GET("/products/:slug/category", productHandler.GetCategoryProduct)
	apiGroup.GET("/products/category/:slug", productHandler.GetProductByCategory)

	adminGroup := apiGroup.Group("/admin")
	adminGroup.Use(middleware.CSRFWithConfig(middleware.CSRFConfig{
		TokenLookup: "cookie:csrf_token",
		CookieName: "csrf_token",
		CookiePath: "/",
		CookieHTTPOnly: true,
		CookieSecure: true,
		CookieSameSite: http.SameSiteStrictMode,
		ErrorHandler: func(err error, c echo.Context) error {
			return echo.NewHTTPError(http.StatusBadRequest, dto.MessageResp{
				Message: "csrf token not available",
			})
		},
	}))
	adminGroup.POST("/products", productHandler.CreateProduct)
	adminGroup.PUT("/products/:slug", productHandler.UpdateProduct)
	adminGroup.DELETE("/products/:slug", productHandler.DeleteProduct)
}