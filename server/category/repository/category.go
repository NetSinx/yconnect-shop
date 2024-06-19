package repository

import (
	"github.com/NetSinx/yconnect-shop/server/category/model/entity"
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

func (c categoryRepository) ListCategory(categories []entity.Category) ([]entity.Category, error) {
	if err := c.DB.Select("id", "name", "slug").Find(&categories).Error; err != nil {
		return nil, err
	}
	
	return categories, nil
}

func (c categoryRepository) CreateCategory(categories entity.Category) (entity.Category, error) {
	if err := c.DB.Create(&categories).Error; err != nil {
		return categories, err
	}

	return categories, nil
}

func (c categoryRepository) UpdateCategory(categories entity.Category, id string) (entity.Category, error) {
	if err := c.DB.Where("id = ?", id).Updates(&categories).Error; err != nil {
		return categories, err
	}

	if err := c.DB.First(&categories, "id = ?", id).Error; err != nil {
		return categories, err
	}
	
	return categories, nil
}

func(c categoryRepository) DeleteCategory(category entity.Category, id string) error {
	if err := c.DB.First(&category, "id = ?", id).Error; err != nil {
		return err
	}

	c.DB.Delete(&category)

	return nil
}

func(c categoryRepository) GetCategory(categories entity.Category, id string) (entity.Category, error) {
	if err := c.DB.First(&categories, "id = ?", id).Error; err != nil {
		return categories, err
	}

	return categories, nil
}