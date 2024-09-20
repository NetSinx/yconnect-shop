package repository

import (
	"github.com/NetSinx/yconnect-shop/server/category/model/entity"
)

type CategoryRepo interface {
	ListCategory(categories []entity.Kategori) ([]entity.Kategori, error)
	CreateCategory(categories entity.Kategori) (entity.Kategori, error)
	UpdateCategory(categories entity.Kategori, id string) (entity.Kategori, error)
	DeleteCategory(category entity.Kategori, id string) error
	GetCategory(categories entity.Kategori, id string) (entity.Kategori, error)
}