package controller

import (
	"github.com/labstack/echo/v4"
)

type CategoryContr interface {
	ListCategory(c echo.Context) error
	CreateCategory(c echo.Context) error
	UpdateCategory(c echo.Context) error
	DeleteCategory(c echo.Context) error
	GetCategory(c echo.Context) error
}