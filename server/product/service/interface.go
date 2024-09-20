package service

import (
	"mime/multipart"

	"github.com/NetSinx/yconnect-shop/server/product/model/entity"
)

type ProductServ interface {
	ListProduct(products []entity.Product) ([]entity.Product, error)
	CreateProduct(products entity.Product, images []*multipart.FileHeader) (entity.Product, error)
	UpdateProduct(products entity.Product, images []*multipart.FileHeader, slug string) (entity.Product, error)
	DeleteProduct(products entity.Product, slug string) error
	GetProduct(products entity.Product, username string) (entity.Product, error)
	GetProductByCategory(products []entity.Product, id string) ([]entity.Product, error)
}