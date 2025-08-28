package route

import (
	"github.com/NetSinx/yconnect-shop/server/category/internal/delivery/http"
	"github.com/NetSinx/yconnect-shop/server/category/internal/delivery/http/middleware"
	"github.com/labstack/echo/v4"
)

type APIRoutes struct {
	App                *echo.Echo
	CategoryController *http.CategoryController
}

func NewAPIRoutes(apiRoutes *APIRoutes) {
	apiRoutes.guestAPIRoutes()
	apiRoutes.authAdminAPIRoutes()
}

func (a *APIRoutes) guestAPIRoutes() {
	guestGroup := a.App.Group("/api")
	guestGroup.GET("/category", a.CategoryController.ListCategory)
	guestGroup.GET("/category/:slug", a.CategoryController.GetCategoryBySlug)
}

func (a *APIRoutes) authAdminAPIRoutes() {
	adminGroup := a.App.Group("/admin")
	adminGroup.Use(middleware.CSRFMiddleware)
	adminGroup.POST("/category", a.CategoryController.CreateCategory)
	adminGroup.PUT("/category/:slug", a.CategoryController.UpdateCategory)
	adminGroup.DELETE("/category/:slug", a.CategoryController.DeleteCategory)
}