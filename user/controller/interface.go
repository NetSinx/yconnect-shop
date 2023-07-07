package controller

import "github.com/labstack/echo/v4"

type UserContr interface {
	RegisterUser(c echo.Context) error
	LoginUser(c echo.Context) error
	ListUsers(c echo.Context) error
	UpdateUser(c echo.Context) error
	FindUser(c echo.Context) error
	DeleteUser(c echo.Context) error
	GetProductFromUser(c echo.Context) error
}