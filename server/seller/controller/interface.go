package controller

import (
	"github.com/labstack/echo/v4"
)

type SellerCtrllr interface {
	ListSeller(c echo.Context) error
}