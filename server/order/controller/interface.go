package controller

import "github.com/labstack/echo/v4"

type OrderController interface {
	ListOrder(c echo.Context) error
	AddOrder(c echo.Context) error
	GetOrder(c echo.Context) error
	DeleteOrder(c echo.Context) error
}