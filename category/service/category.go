package service

import (
	"errors"
	"github.com/NetSinx/yconnect-shop/category/app/model"
	"github.com/NetSinx/yconnect-shop/category/repository"
)

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
		return nil, errors.New("couldn't get all data")
	}

	return categories, nil
}

func (c categoryService) CreateCategory(categories model.Category) error {
	err := c.categoryRepo.CreateCategory(categories)

	if err != nil {
		return errors.New("cannot create data")
	}

	return nil
}

func (c categoryService) UpdateCategory(categories model.Category, slug string) error {
	err := c.categoryRepo.UpdateCategory(categories, slug)

	if err != nil {
		return err
	}

	return nil
}

func (c categoryService) DeleteCategory(category model.Category, slug string) error {
	if err := c.categoryRepo.DeleteCategory(category, slug); err != nil {
		return errors.New("cannot delete data")
	}

	return nil
}

func (c categoryService) GetCategory(categories model.Category, slug string) (model.Category, error) {
	getCategory, err := c.categoryRepo.GetCategory(categories, slug); if err != nil {
		return getCategory, errors.New("cannot get category")
	}

	return getCategory, nil
}