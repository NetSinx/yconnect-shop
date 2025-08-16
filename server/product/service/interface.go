package service

import (
	categoryEntity "github.com/NetSinx/yconnect-shop/server/category/model/entity"
	"github.com/NetSinx/yconnect-shop/server/product/model/entity"
)

type ProductServ interface {
	ListProduct(products []entity.Product) ([]entity.Product, error)
	CreateProduct(product entity.Product) error
	UpdateProduct(product entity.Product, slug string) error
	DeleteProduct(product entity.Product, slug string) error
	GetProductByID(product entity.Product, id string) (entity.Product, error)
	GetProductBySlug(product entity.Product, slug string) (entity.Product, error)
	GetCategoryProduct(product entity.Product, slug string) (categoryEntity.Category, error)
}