package repository

import (
	"github.com/NetSinx/yconnect-shop/category/app/model"
	"gorm.io/gorm"
)

type categoryRepository struct {
	DB *gorm.DB
}

func CategoryRepository(db *gorm.DB) categoryRepository {
	return categoryRepository{
		DB: db,
	}
}

func (c categoryRepository) ListCategory(categories []model.Category) ([]model.Category, error) {

	if err := c.DB.Select("id", "name", "slug").Find(&categories).Error; err != nil {
		return nil, err
	}
	
	return categories, nil
}

func (c categoryRepository) CreateCategory(categories model.Category) error {
	if err := c.DB.Create(&categories).Error; err != nil {
		return err
	}

	return nil
}

func (c categoryRepository) UpdateCategory(categories model.Category, slug string) error {
	err := c.DB.Where("slug = ?", slug).Updates(&categories).Error; if err != nil {
		return err
	}

	if err := c.DB.First(&categories, "slug = ?", slug).Error; err != nil {
		return err
	}
	
	return nil
}

func(c categoryRepository) DeleteCategory(category model.Category, slug string) error {
	if err := c.DB.First(&category, "slug = ?", slug).Error; err != nil {
		return err
	}

	c.DB.Delete(&category)

	return nil
}

func(c categoryRepository) GetCategory(categories model.Category, slug string) (model.Category, error) {
	if err := c.DB.First(&categories, "slug = ?", slug).Error; err != nil {
		return categories, err
	}

	return categories, nil
}