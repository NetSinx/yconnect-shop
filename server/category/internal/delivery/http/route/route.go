package route

import (
	"github.com/NetSinx/yconnect-shop/server/category/internal/delivery/http"
	"github.com/NetSinx/yconnect-shop/server/category/internal/delivery/http/middleware"
	"github.com/labstack/echo/v4"
)

type APIRoutes struct {
	AppGroup           *echo.Group
	CategoryController *http.CategoryController
}

func NewAPIRoutes(apiRoutes *APIRoutes) {
	apiRoutes.guestAPIRoutes()
	apiRoutes.authAdminAPIRoutes()
}

func (a *APIRoutes) guestAPIRoutes() {
	guestGroup := a.AppGroup
	guestGroup.GET("/category", a.CategoryController.ListCategory)
	guestGroup.GET("/category/:slug", a.CategoryController.GetCategoryBySlug)
}

func (a *APIRoutes) authAdminAPIRoutes() {
	adminGroup := a.AppGroup.Group("/admin")
	adminGroup.Use(middleware.CSRFMiddleware)
	adminGroup.POST("/category", a.CategoryController.CreateCategory)
	adminGroup.PUT("/category/:slug", a.CategoryController.UpdateCategory)
	adminGroup.DELETE("/category/:slug", a.CategoryController.DeleteCategory)
}
