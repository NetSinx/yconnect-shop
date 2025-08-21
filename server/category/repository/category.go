package repository

import (
	"github.com/NetSinx/yconnect-shop/server/category/model"
	"gorm.io/gorm"
)

type CategoryRepo interface {
	ListCategory(categories []model.Category) ([]model.Category, error)
	CreateCategory(category model.Category) error
	UpdateCategory(category model.Category, slug string) error
	DeleteCategory(category model.Category, slug string) error
	GetCategoryById(category model.Category, id string) (model.Category, error)
	GetCategoryBySlug(category model.Category, slug string) (model.Category, error)
}

type categoryRepository struct {
	DB *gorm.DB
}

func CategoryRepository(db *gorm.DB) categoryRepository {
	return categoryRepository{
		DB: db,
	}
}

func (c categoryRepository) ListCategory(categories []model.Category) ([]model.Category, error) {
	if err := c.DB.Find(&categories).Error; err != nil {
		return categories, err
	}
	
	return categories, nil
}

func (c categoryRepository) CreateCategory(category model.Category) error {
	if err := c.DB.Create(&category).Error; err != nil {
		return err
	}

	return nil
}

func (c categoryRepository) UpdateCategory(category model.Category, slug string) error {
	if err := c.DB.Where("slug = ?", slug).Updates(&category).Error; err != nil {
		return err
	}
	
	return nil
}

func(c categoryRepository) DeleteCategory(category model.Category, slug string) error {
	if err := c.DB.First(&category, "slug = ?", slug).Error; err != nil {
		return err
	}

	if err := c.DB.Delete(&category).Error; err != nil {
		return err
	}

	return nil
}

func(c categoryRepository) GetCategoryById(category model.Category, id string) (model.Category, error) {
	if err := c.DB.First(&category, "id = ?", id).Error; err != nil {
		return category, err
	}

	return category, nil
}

func(c categoryRepository) GetCategoryBySlug(category model.Category, slug string) (model.Category, error) {
	if err := c.DB.First(&category, "slug = ?", slug).Error; err != nil {
		return category, err
	}

	return category, nil
}
