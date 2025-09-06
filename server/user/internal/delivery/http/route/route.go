package route

import (
	"github.com/NetSinx/yconnect-shop/server/user/internal/delivery/http"
	"github.com/labstack/echo/v4"
)

type APIRoutes struct {
	AppGroup       *echo.Group
	UserController *http.UserController
}

func NewApiRoutes(apiRoutes *APIRoutes) {
	apiGroup := apiRoutes.AppGroup
	apiGroup.GET("/users/:username", apiRoutes.UserController.GetUserByUsername)
	apiGroup.PUT("/users/:username", apiRoutes.UserController.UpdateUser)
	apiGroup.DELETE("/users/:username", apiRoutes.UserController.DeleteUser)
}