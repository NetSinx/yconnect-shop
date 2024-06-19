package repository

import "github.com/NetSinx/yconnect-shop/server/product/model/entity"

type ProductRepo interface {
	ListProduct(products []entity.Product) ([]entity.Product, error)
	CreateProduct(products entity.Product, image []entity.Image) (entity.Product, error)
	UpdateProduct(products entity.Product, image []entity.Image, slug string, id string) (entity.Product, error)
	DeleteProduct(products entity.Product, image []entity.Image, slug string, id string) error
	GetProduct(products entity.Product, slug string) (entity.Product, error)
	GetProductByCategory(products []entity.Product, id string) ([]entity.Product, error)
	GetProductBySeller(products []entity.Product, id string) ([]entity.Product, error)
}