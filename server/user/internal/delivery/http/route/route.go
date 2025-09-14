package route

import (
	httpController "github.com/NetSinx/yconnect-shop/server/user/internal/delivery/http"
	"github.com/labstack/echo/v4/middleware"
	"github.com/labstack/echo/v4"
	"net/http"
)

type APIRoutes struct {
	AppGroup       *echo.Group
	UserController *httpController.UserController
}

func NewApiRoutes(apiRoutes *APIRoutes) {
	apiGroup := apiRoutes.AppGroup
	apiGroup.Use(middleware.CSRFWithConfig(middleware.CSRFConfig{
		CookieName: "csrf_token",
		TokenLookup: "header:X-CSRF-Token",
		ContextKey: "csrf_token",
		CookiePath: "/",
		CookieHTTPOnly: true,
		CookieSecure: true,
		CookieMaxAge: 600,
		CookieSameSite: http.SameSiteStrictMode,
	}))
	apiGroup.GET("/user/:id", apiRoutes.UserController.GetUserByID)
	apiGroup.PUT("/user/:id", apiRoutes.UserController.UpdateUser)
	apiGroup.DELETE("/user/:id", apiRoutes.UserController.DeleteUser)
}