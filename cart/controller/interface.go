package controller

import "github.com/labstack/echo/v4"

type CartControl interface {
	ListCart(c echo.Context) error
	AddToCart(c echo.Context) error
	UpdateCart(c echo.Context) error
	DeleteProductInCart(c echo.Context) error
	GetCartById(c echo.Context) error
	GetCartBySlug(c echo.Context) error
}