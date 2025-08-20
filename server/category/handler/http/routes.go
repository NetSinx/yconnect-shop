package http

import (
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func ApiRoutes(e *echo.Echo, categoryHandler categoryHandler) {
	apiGroup := e.Group("/api")
	apiGroup.Use(middleware.KeyAuthWithConfig(middleware.KeyAuthConfig{
		KeyLookup: "header:api-token",
		Validator: func(auth string, c echo.Context) (bool, error) {
			return auth == "dfkgjdgj#753846873248358645*&#%^*$54%hgdf", nil
		},
	}))
	apiGroup.GET("/category", categoryHandler.ListCategory)
	apiGroup.GET("/category/id/:id", categoryHandler.GetCategoryById)
	adminGroup := apiGroup.Group("/admin")
	adminGroup.Use(echojwt.WithConfig(echojwt.Config{
		SigningKey: []byte("yasinnetsinx15"),
		SigningMethod: "HS512",
	}))
	adminGroup.POST("/category", categoryHandler.CreateCategory)
	adminGroup.PUT("/category/:id", categoryHandler.UpdateCategory)
	adminGroup.DELETE("/category/:id", categoryHandler.DeleteCategory)
}