package http

import (
	"github.com/NetSinx/yconnect-shop/server/category/internal/delivery/http"
	"github.com/NetSinx/yconnect-shop/server/category/internal/delivery/http/middleware"
	"github.com/labstack/echo/v4"
)

type APIRoutes struct {
	App                *echo.Echo
	AppGroup           *echo.Group
	CSRFMiddleware     *middleware.CSRFMiddleware
	CategoryController *http.CategoryController
}

func (a *APIRoutes) NewAPIRoutes() {
	a.GuestAPIRoutes()
	a.AuthAdminAPIRoutes()
}

func (a *APIRoutes) GuestAPIRoutes() {
	a.AppGroup = a.App.Group("/api")
	a.AppGroup.GET("/category", a.CategoryController.ListCategory)
	a.AppGroup.GET("/category/:slug", a.CategoryController.GetCategoryBySlug)
}

func (a *APIRoutes) AuthAdminAPIRoutes() {
	a.AppGroup = a.App.Group("/admin")
	a.AppGroup.Use(a.CSRFMiddleware.NewCSRFMiddleware)
	a.AppGroup.POST("/category", a.CategoryController.CreateCategory)
	a.AppGroup.PUT("/category/:slug", a.CategoryController.UpdateCategory)
	a.AppGroup.DELETE("/category/:slug", a.CategoryController.DeleteCategory)
}