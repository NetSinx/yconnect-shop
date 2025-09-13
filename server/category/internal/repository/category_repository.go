package repository

import (
	"github.com/NetSinx/yconnect-shop/server/category/internal/entity"
	"github.com/NetSinx/yconnect-shop/server/category/internal/model"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type CategoryRepository struct {
	Log *logrus.Logger
}

func NewCategoryRepository(log *logrus.Logger) *CategoryRepository {
	return &CategoryRepository{
		Log: log,
	}
}

func (r *CategoryRepository) ListCategory(db *gorm.DB, request *model.ListCategoryRequest) ([]entity.Category, int64, error) {
	var categories []entity.Category

	if err := db.Offset((request.Page - 1) * request.Size).Limit(request.Size).Find(&categories).Error; err != nil {
		return nil, 0, err
	}

	var total int64
	if err := db.Model(&entity.Category{}).Offset((request.Page - 1) * request.Size).Limit(request.Size).Count(&total).Error; err != nil {
		return nil, 0, err
	}

	return categories, total, nil
}

func (r *CategoryRepository) CreateCategory(db *gorm.DB, category *entity.Category) (uint, error) {
	if err := db.Create(category).Error; err != nil {
		return 0, err
	}

	return category.ID, nil
}

func (r *CategoryRepository) UpdateCategory(db *gorm.DB, category *entity.Category) error {
	return db.Save(category).Error
}

func (r *CategoryRepository) DeleteCategory(db *gorm.DB, category *entity.Category, slug string) error {
	return db.Delete(category, "slug = ?", slug).Error
}

func(c *CategoryRepository) GetCategoryID(db *gorm.DB, category *entity.Category) (uint, error) {
	if err := db.Select("id").First(category, "slug = ?", category.Slug).Error; err != nil {
		return 0, err
	}

	return category.ID, nil
}

func(c *CategoryRepository) GetCategoryBySlug(db *gorm.DB, category *entity.Category, slug string) (*entity.Category, error) {
	if err := db.First(category, "slug = ?", slug).Error; err != nil {
		return nil, err
	}

	return category, nil
}