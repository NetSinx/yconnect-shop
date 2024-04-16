package service

import "github.com/NetSinx/yconnect-shop/server/product/model"

type ProductServ interface {
	ListProduct(products []model.Product) ([]model.Product, error)
	CreateProduct(products model.Product, image []model.Image) (model.Product, error)
	UpdateProduct(products model.Product, image []model.Image, slug string, id string) (model.Product, error)
	DeleteProduct(products model.Product, image []model.Image, slug string, id string) error
	GetProduct(products model.Product, username string) (model.Product, error)
	GetProductByCategory(products []model.Product, id string) ([]model.Product, error)
	GetProductBySeller(products []model.Product, id string) ([]model.Product, error)
}