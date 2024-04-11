package controller

import (
	"github.com/labstack/echo/v4"
)

type SellerCtrllr interface {
	ListSeller(c echo.Context) error
	RegisterSeller(c echo.Context) error
	UpdateSeller(c echo.Context) error
	DeleteSeller(c echo.Context) error
	GetSeller(c echo.Context) error
}