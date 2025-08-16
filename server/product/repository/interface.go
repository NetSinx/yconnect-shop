package repository

import "github.com/NetSinx/yconnect-shop/server/product/model/entity"

type ProductRepo interface {
	ListProduct(products []entity.Product) ([]entity.Product, error)
	CreateProduct(product entity.Product) error
	UpdateProduct(product entity.Product, slug string) error
	DeleteProduct(product entity.Product, slug string) error
	GetProductByID(product entity.Product, id string) (entity.Product, error)
	GetProductBySlug(product entity.Product, slug string) (entity.Product, error)
	GetCategoryProduct(product entity.Product, slug string) error
}