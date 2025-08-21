package http

import (
	"github.com/NetSinx/yconnect-shop/server/authentication/middleware"
	"github.com/labstack/echo/v4"
)

func ApiRoutes(e *echo.Echo, categoryHandler categoryHandler) {
	apiGroup := e.Group("/api")
	apiGroup.GET("/category", categoryHandler.ListCategory)
	apiGroup.GET("/category/id/:id", categoryHandler.GetCategoryById)
	apiGroup.GET("/category/slug/:slug", categoryHandler.GetCategoryBySlug)

	adminGroup := apiGroup.Group("/admin")
	adminGroup.Use(middleware.CSRFMiddleware)
	adminGroup.POST("/category", categoryHandler.CreateCategory)
	adminGroup.PUT("/category/:slug", categoryHandler.UpdateCategory)
	adminGroup.DELETE("/category/:slug", categoryHandler.DeleteCategory)
}