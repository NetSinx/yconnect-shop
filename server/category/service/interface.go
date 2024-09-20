package service

import (
	"github.com/NetSinx/yconnect-shop/server/category/model/entity"
)
	
type CategoryServ interface {
	ListCategory(categories []entity.Kategori) ([]entity.Kategori, error)
	CreateCategory(category entity.Kategori) (entity.Kategori, error)
	UpdateCategory(category entity.Kategori, id string) (entity.Kategori, error)
	DeleteCategory(category entity.Kategori, id string) error
	GetCategory(category entity.Kategori, id string) (entity.Kategori, error)
}