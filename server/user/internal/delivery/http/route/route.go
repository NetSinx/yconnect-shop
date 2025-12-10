package route

import (
	"net/http"
	httpController "github.com/NetSinx/yconnect-shop/server/user/internal/delivery/http"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type APIRoutes struct {
	App       *echo.Echo
	UserController *httpController.UserController
}

func NewApiRoutes(apiRoutes *APIRoutes) {
	apiRoutes.App.POST("/user/register", apiRoutes.UserController.RegisterUser)

	apiGroup := apiRoutes.App.Group("/api")
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