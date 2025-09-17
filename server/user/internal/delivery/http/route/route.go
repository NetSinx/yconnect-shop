package route

import (
	"net/http"
	httpController "github.com/NetSinx/yconnect-shop/server/user/internal/delivery/http"
	"github.com/NetSinx/yconnect-shop/server/user/internal/delivery/http/middleware"
	"github.com/labstack/echo/v4"
	echoMiddleware "github.com/labstack/echo/v4/middleware"
)

type APIRoutes struct {
	AppGroup       *echo.Group
	UserController *httpController.UserController
}

func NewApiRoutes(apiRoutes *APIRoutes) {
	apiGroup := apiRoutes.AppGroup
	apiGroup.Use(echoMiddleware.CSRFWithConfig(echoMiddleware.CSRFConfig{
		CookieName: "csrf_token",
		TokenLookup: "header:X-CSRF-Token",
		ContextKey: "csrf_token",
		CookiePath: "/",
		CookieHTTPOnly: true,
		CookieSecure: true,
		CookieMaxAge: 600,
		CookieSameSite: http.SameSiteStrictMode,
	}))
	apiGroup.Use(middleware.AuthorizationMiddleware)
	apiGroup.GET("/user/:id", apiRoutes.UserController.GetUserByID)
	apiGroup.PUT("/user/:id", apiRoutes.UserController.UpdateUser)
	apiGroup.DELETE("/user/:id", apiRoutes.UserController.DeleteUser)
}