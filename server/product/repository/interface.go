package repository

import "github.com/NetSinx/yconnect-shop/server/product/app/model"

type ProductRepo interface {
	ListProduct(products []model.Product) ([]model.Product, error)
	CreateProduct(products model.Product, image []model.Image) (model.Product, error)
	UpdateProduct(products model.Product, image []model.Image, id uint) (model.Product, error)
	DeleteProduct(products model.Product, image []model.Image, id string) error
	GetProduct(products model.Product, id string) (model.Product, error)
}