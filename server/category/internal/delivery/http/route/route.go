package route

import (
	"net/http"

	httpController "github.com/NetSinx/yconnect-shop/server/category/internal/delivery/http"
	"github.com/NetSinx/yconnect-shop/server/category/internal/delivery/http/middleware"
	"github.com/labstack/echo/v4"
	echoMiddleware "github.com/labstack/echo/v4/middleware"
)

type APIRoutes struct {
	AppGroup           *echo.Group
	CategoryController *httpController.CategoryController
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
	adminGroup.Use(echoMiddleware.CSRFWithConfig(echoMiddleware.CSRFConfig{
		CookieName: "csrf_token",
		TokenLookup: "cookie:csrf_token",
		ContextKey: "csrf_token",
		CookiePath: "/",
		CookieHTTPOnly: true,
		CookieSecure: true,
		CookieSameSite: http.SameSiteStrictMode,
	}))
	adminGroup.Use(middleware.AuthorizationMiddleware)
	adminGroup.POST("/category", a.CategoryController.CreateCategory)
	adminGroup.PUT("/category/:slug", a.CategoryController.UpdateCategory)
	adminGroup.DELETE("/category/:slug", a.CategoryController.DeleteCategory)
}
