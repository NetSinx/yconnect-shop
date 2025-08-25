package service

import (
	"github.com/NetSinx/yconnect-shop/server/category/handler/dto"
	"github.com/NetSinx/yconnect-shop/server/category/helpers"
	"github.com/NetSinx/yconnect-shop/server/category/model"
	"github.com/NetSinx/yconnect-shop/server/category/repository"
	"github.com/NetSinx/yconnect-shop/server/category/service/rabbitmq"
	"github.com/go-playground/validator/v10"
)

type CategoryServ interface {
	ListCategory(categories []model.Category) ([]model.Category, error)
	CreateCategory(categoryReq dto.CategoryRequest) error
	UpdateCategory(categoryReq dto.CategoryRequest, slug string) error
	DeleteCategory(category model.Category, slug string) error
	GetCategoryById(category model.Category, id string) (model.Category, error)
	GetCategoryBySlug(category model.Category, slug string) (model.Category, error)
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
	listCategories, err := c.categoryRepo.ListCategory(categories)
	if err != nil {
		return nil, err
	}

	return listCategories, nil
}

func (c categoryService) CreateCategory(categoryReq dto.CategoryRequest) error {
	categoryReq.Name = helpers.ToTitle(categoryReq.Name)
	slug := helpers.ToSlug(categoryReq.Name)

	if err := validator.New().Struct(categoryReq); err != nil {
		return err
	}

	category := model.Category{
		Name: categoryReq.Name,
		Slug: slug,
	}

	if err := c.categoryRepo.CreateCategory(category); err != nil {
		return err
	}

	rabbitmq.Publisher(rabbitmq.RoutingCKCreated, category)

	return nil
}

func (c categoryService) UpdateCategory(categoryReq dto.CategoryRequest, slug string) error {
	categoryReq.Name = helpers.ToTitle(categoryReq.Name)
	newSlug := helpers.ToSlug(categoryReq.Name)
	
	if err := validator.New().Struct(categoryReq); err != nil {
		return err
	}
	
	category := model.Category{
		Name: categoryReq.Name,
		Slug: newSlug,
	}

	if err := c.categoryRepo.UpdateCategory(category, slug); err != nil {
		return err
	}

	category = model.Category{
		Slug: slug,
	}
	
	rabbitmq.Publisher(rabbitmq.RoutingCKUpdated, category)

	return nil
}

func (c categoryService) DeleteCategory(category model.Category, slug string) error {
	if err := c.categoryRepo.DeleteCategory(category, slug); err != nil {
		return err
	}

	category = model.Category{
		Slug: slug,
	}
	rabbitmq.Publisher(rabbitmq.RoutingCKDeleted, category)

	return nil
}

func (c categoryService) GetCategoryById(category model.Category, id string) (model.Category, error) {
	getCategory, err := c.categoryRepo.GetCategoryById(category, id)
	if err != nil {
		return getCategory, err
	}

	return getCategory, nil
}

func (c categoryService) GetCategoryBySlug(category model.Category, slug string) (model.Category, error) {
	getCategory, err := c.categoryRepo.GetCategoryBySlug(category, slug)
	if err != nil {
		return getCategory, err
	}

	return getCategory, nil
}