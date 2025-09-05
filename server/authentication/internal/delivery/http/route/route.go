package route

import (
	httpController"github.com/NetSinx/yconnect-shop/server/authentication/internal/delivery/http"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"net/http"
)

type APIRoutes struct {
	AppGroup       *echo.Group
	AuthController *httpController.AuthController
}

func NewAPIRoutes(apiRoutes *APIRoutes) {
	apiGroup := apiRoutes.AppGroup
	apiGroup.Use(middleware.CSRFWithConfig(middleware.CSRFConfig{
		CookieName: "csrf_token",
		TokenLookup: "cookie:csrf_token",
		ContextKey: "csrf_token",
		CookiePath: "/",
		CookieHTTPOnly: true,
		CookieSecure: true,
		CookieMaxAge: 600,
		CookieSameSite: http.SameSiteStrictMode,
	}))
	apiGroup.GET("/auth/csrf-token", apiRoutes.AuthController.GetCSRFToken)
	apiGroup.POST("/auth/login", apiRoutes.AuthController.LoginUser)
	apiGroup.POST("/auth/verify", apiRoutes.AuthController.Verify)
	apiGroup.POST("/auth/logout", apiRoutes.AuthController.LogoutUser)
}
