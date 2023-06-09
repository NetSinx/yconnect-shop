package repository

import (
	"github.com/NetSinx/yconnect-shop/category/app/model"
)

type CategoryRepo interface {
	ListCategory(categories []model.Category) ([]model.Category, error)
	CreateCategory(categories model.Category) error
	UpdateCategory(categories model.Category, slug string) error
	DeleteCategory(category model.Category, slug string) error
	GetCategoryBySlug(categories model.Category, slug string) (model.Category, error)
	GetCategoryById(categories model.Category, id string) (model.Category, error)
}