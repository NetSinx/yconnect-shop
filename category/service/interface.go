package service

import (
	"github.com/NetSinx/yconnect-shop/category/app/model"	
)
	
type CategoryServ interface {
	ListCategory(categories []model.Category) ([]model.Category, error)
	CreateCategory(categories model.Category) error
	UpdateCategory(categories model.Category, slug string) error
	DeleteCategory(category model.Category, slug string) error
	GetCategory(categories model.Category, slug string) (model.Category, error)
}