package service

import (
	"github.com/NetSinx/yconnect-shop/server/category/app/model"	
)
	
type CategoryServ interface {
	ListCategory(categories []model.Category) ([]model.Category, error)
	CreateCategory(categories model.Category) (model.Category, error)
	UpdateCategory(categories model.Category, id string) (model.Category, error)
	DeleteCategory(category model.Category, id string) error
	GetCategory(categories model.Category, id string) (model.Category, error)
}