package service

import (
	"strings"
	"github.com/NetSinx/yconnect-shop/server/category/model"
	"github.com/NetSinx/yconnect-shop/server/category/repository"
	"github.com/go-playground/validator/v10"
)

type CategoryServ interface {
	ListCategory(categories []model.Category) ([]model.Category, error)
	CreateCategory(category model.Category) (model.Category, error)
	UpdateCategory(category model.Category, id string) (model.Category, error)
	DeleteCategory(category model.Category, id string) error
	GetCategoryById(category model.Category, id string) (model.Category, error)
}

type categoryService struct {
	categoryRepo repository.CategoryRepo
}

func CategoryService(categoryrepo repository.CategoryRepo) categoryService {
	return categoryService{
		categoryRepo: categoryrepo,
	}
}

func (c categoryService) ListCategory(categories []model.Category) ([]model.Category, error) {
	categories, err := c.categoryRepo.ListCategory(categories)
	if err != nil {
		return nil, err
	}

	return categories, nil
}

func (c categoryService) CreateCategory(category model.Category) (model.Category, error) {
	category.Name = strings.ToTitle(category.Name)
	category.Slug = strings.ToLower(category.Slug)

	if err := validator.New().Struct(category); err != nil {
		return category, err
	}

	createCategory, err := c.categoryRepo.CreateCategory(category)
	if err != nil {
		return category, err
	}

	return createCategory, nil
}

func (c categoryService) UpdateCategory(category model.Category, id string) (model.Category, error) {
	category.Name = strings.Title(category.Name)
	category.Slug = strings.ToLower(category.Slug)
	
	if err := validator.New().Struct(category); err != nil {
		return category, err
	}

	updCategory, err := c.categoryRepo.UpdateCategory(category, id)
	if err != nil {
		return category, err
	}

	return updCategory, nil
}

func (c categoryService) DeleteCategory(category model.Category, id string) error {
	if err := c.categoryRepo.DeleteCategory(category, id); err != nil {
		return err
	}

	return nil
}

func (c categoryService) GetCategoryById(category model.Category, id string) (model.Category, error) {
	getCategory, err := c.categoryRepo.GetCategoryById(category, id)
	if err != nil {
		return getCategory, err
	}

	return getCategory, nil
}