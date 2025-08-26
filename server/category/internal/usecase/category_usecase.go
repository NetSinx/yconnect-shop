package usecase

import (
	"context"
	"github.com/NetSinx/yconnect-shop/server/category/internal/entity"
	"github.com/NetSinx/yconnect-shop/server/category/internal/helpers"
	"github.com/NetSinx/yconnect-shop/server/category/internal/model"
	"github.com/NetSinx/yconnect-shop/server/category/internal/repository"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type CategoryUseCase struct {
	DB                 *gorm.DB
	Log                *logrus.Logger
	Validator          *validator.Validate
	categoryRepository *repository.CategoryRepository
}

func NewCategoryUseCase(db *gorm.DB, log *logrus.Logger, validator *validator.Validate, categoryRepository *repository.CategoryRepository) *CategoryUseCase {
	return &CategoryUseCase{
		DB:                 db,
		Log:                log,
		Validator:          validator,
		categoryRepository: categoryRepository,
	}
}

func (c *CategoryUseCase) ListCategory(ctx context.Context, db *gorm.DB, categoryRequest *model.ListCategoryRequest) ([]entity.Category, int64, error) {
	tx := db.WithContext(ctx).Begin()
	defer tx.Rollback()

	if err := c.Validator.Struct(categoryRequest); err != nil {
		c.Log.WithError(err).Error("error validating request body")
		return nil, 0, echo.ErrBadRequest
	}

	listCategories, total, err := c.categoryRepository.ListCategory(tx, categoryRequest)
	if err != nil {
		c.Log.WithError(err).Error("error listing categories")
		return nil, 0, echo.ErrInternalServerError
	}

	if err := tx.Commit().Error; err != nil {
		c.Log.WithError(err).Error("error listing categories")
		return nil, 0, echo.ErrInternalServerError
	}

	return listCategories, total, nil
}

func (c *CategoryUseCase) CreateCategory(ctx context.Context, db *gorm.DB, categoryReq *model.CreateCategoryRequest) error {
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

func (c *CategoryUseCase) UpdateCategory(categoryReq dto.CategoryRequest, slug string) error {
	categoryReq.Name = helpers.ToTitle(categoryReq.Name)
	newSlug := helpers.ToSlug(categoryReq.Name)

	if err := validator.New().Struct(categoryReq); err != nil {
		return err
	}

	category := model.Category{
		Name: categoryReq.Name,
		Slug: newSlug,
	}

	id, err := c.categoryRepo.UpdateCategory(category, slug)
	if err != nil {
		return err
	}

	category.Id = id

	rabbitmq.Publisher(rabbitmq.RoutingCKUpdated, category)

	return nil
}

func (c *CategoryUseCase) DeleteCategory(category model.Category, slug string) error {
	if err := c.categoryRepo.DeleteCategory(category, slug); err != nil {
		return err
	}

	category = model.Category{
		Slug: slug,
	}
	rabbitmq.Publisher(rabbitmq.RoutingCKDeleted, category)

	return nil
}

func (c *CategoryUseCase) GetCategoryById(category model.Category, id string) (model.Category, error) {
	getCategory, err := c.categoryRepo.GetCategoryById(category, id)
	if err != nil {
		return getCategory, err
	}

	return getCategory, nil
}

func (c *CategoryUseCase) GetCategoryBySlug(category model.Category, slug string) (model.Category, error) {
	getCategory, err := c.categoryRepo.GetCategoryBySlug(category, slug)
	if err != nil {
		return getCategory, err
	}

	return getCategory, nil
}
