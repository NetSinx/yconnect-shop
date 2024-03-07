package controller

import "github.com/labstack/echo/v4"

type ProductContr interface {
	ListProduct(c echo.Context) error
	CreateProduct(c echo.Context) error
	UpdateProduct(c echo.Context) error
	DeleteProduct(c echo.Context) error
	GetProduct(c echo.Context) error
	GetProductByCategory(c echo.Context) error
	GetProductBySeller(c echo.Context) error
}