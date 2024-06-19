package service

import (
	"github.com/NetSinx/yconnect-shop/server/category/model/entity"
)
	
type CategoryServ interface {
	ListCategory(categories []entity.Category) ([]entity.Category, error)
	CreateCategory(categories entity.Category) (entity.Category, error)
	UpdateCategory(categories entity.Category, id string) (entity.Category, error)
	DeleteCategory(category entity.Category, id string) error
	GetCategory(categories entity.Category, id string) (entity.Category, error)
}