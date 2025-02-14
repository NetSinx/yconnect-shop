package controller

import "github.com/labstack/echo/v4"

type ProductContr interface {
	ListProduct(c echo.Context) error
	CreateProduct(c echo.Context) error
	UpdateProduct(c echo.Context) error
	DeleteProduct(c echo.Context) error
	GetProductByID(c echo.Context) error
	GetProductBySlug(c echo.Context) error
	GetProductByCategory(c echo.Context) error
}