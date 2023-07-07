package repository

import "github.com/NetSinx/yconnect-shop/product/app/model"

type ProductRepo interface {
	ListProduct(products []model.Product) ([]model.Product, error)
	CreateProduct(products model.Product) error
	UpdateProduct(products model.Product, slug string) error
	DeleteProduct(products model.Product, slug string) error
	GetProduct(products model.Product, slug string) (model.Product, error)
}