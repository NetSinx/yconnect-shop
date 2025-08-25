package http

import (
	"net/http"
	"github.com/NetSinx/yconnect-shop/server/category/handler/dto"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func ApiRoutes(e *echo.Echo, categoryHandler categoryHandler) {
	apiGroup := e.Group("/api")
	apiGroup.GET("/category", categoryHandler.ListCategory)
	apiGroup.GET("/category/id/:id", categoryHandler.GetCategoryById)
	apiGroup.GET("/category/slug/:slug", categoryHandler.GetCategoryBySlug)

	adminGroup := apiGroup.Group("/admin")
	adminGroup.Use(middleware.CSRFWithConfig(middleware.CSRFConfig{
		TokenLookup: "cookie:csrf_token",
		CookieName: "csrf_token",
		CookiePath: "/",
		CookieHTTPOnly: true,
		CookieSecure: true,
		ErrorHandler: func(err error, c echo.Context) error {
			return echo.NewHTTPError(http.StatusBadRequest, dto.MessageResp{
				Message: "csrf token not available",
			})
		},
	}))
	adminGroup.POST("/category", categoryHandler.CreateCategory)
	adminGroup.PUT("/category/:slug", categoryHandler.UpdateCategory)
	adminGroup.DELETE("/category/:slug", categoryHandler.DeleteCategory)
}