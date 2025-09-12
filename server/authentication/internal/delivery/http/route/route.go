package route

import (
	"net/http"
	httpController "github.com/NetSinx/yconnect-shop/server/authentication/internal/delivery/http"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type APIRoutes struct {
	AppGroup       *echo.Group
	AuthController *httpController.AuthController
}

func NewAPIRoutes(apiRoutes *APIRoutes) {
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
	apiGroup.GET("/auth/csrf-token", apiRoutes.AuthController.GetCSRFToken)
	apiGroup.POST("/auth/register", apiRoutes.AuthController.RegisterUser)
	apiGroup.POST("/auth/login", apiRoutes.AuthController.LoginUser)
	apiGroup.POST("/auth/verify", apiRoutes.AuthController.Verify)
	apiGroup.POST("/auth/refresh", apiRoutes.AuthController.RefreshToken)
	apiGroup.POST("/auth/logout", apiRoutes.AuthController.LogoutUser)
}
