package controller

import "github.com/labstack/echo/v4"

type CartControl interface {
	ListCart(c echo.Context) error
	AddToCart(c echo.Context) error
	DeleteProductInCart(c echo.Context) error
	GetCart(c echo.Context) error
	GetCartByUser(c echo.Context) error
}