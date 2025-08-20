package http

import (
	"net/http"
	"github.com/NetSinx/yconnect-shop/server/product/handler/dto"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func ApiRoutes(e *echo.Echo, productHandler productHandler) {
	apiGroup := e.Group("/api")
	apiGroup.GET("/product", productHandler.ListProduct)
	apiGroup.GET("/product/id/:id", productHandler.GetProductByID)
	apiGroup.GET("/product/slug/:slug", productHandler.GetProductBySlug)
	apiGroup.GET("/product/:slug/category", productHandler.GetCategoryProduct)

	adminGroup := apiGroup.Group("/admin")
	adminGroup.Use(middleware.CSRFWithConfig(middleware.CSRFConfig{
		TokenLookup: "cookie:csrf_token",
		CookieName: "csrf_token",
		CookiePath: "/",
		CookieHTTPOnly: true,
		CookieMaxAge: 30,
		CookieSecure: true,
		ErrorHandler: func(err error, c echo.Context) error {
			return echo.NewHTTPError(http.StatusBadRequest, dto.MessageResp{
				Message: "csrf token not available",
			})
		},
	}))
	adminGroup.POST("/product", productHandler.CreateProduct)
	adminGroup.PUT("/product/:slug", productHandler.UpdateProduct)
	adminGroup.DELETE("/product/:slug", productHandler.DeleteProduct)
}