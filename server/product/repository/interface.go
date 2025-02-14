package repository

import "github.com/NetSinx/yconnect-shop/server/product/model/entity"

type ProductRepo interface {
	ListProduct(products []entity.Product) ([]entity.Product, error)
	CreateProduct(products entity.Product) (entity.Product, error)
	UpdateProduct(products entity.Product, slug string) (entity.Product, error)
	DeleteProduct(products entity.Product, slug string) error
	GetProductByID(product entity.Product, id string) (entity.Product, error)
	GetProductBySlug(products entity.Product, slug string) (entity.Product, error)
	GetProductByCategory(products []entity.Product, id string) ([]entity.Product, error)
}