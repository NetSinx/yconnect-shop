package usecase

import (
	"context"
	"github.com/NetSinx/yconnect-shop/server/category/internal/entity"
	"github.com/NetSinx/yconnect-shop/server/category/internal/gateway/messaging"
	"github.com/NetSinx/yconnect-shop/server/category/internal/helpers"
	"github.com/NetSinx/yconnect-shop/server/category/internal/model"
	"github.com/NetSinx/yconnect-shop/server/category/internal/model/converter"
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
	Helpers            *helpers.Helpers
	CategoryRepository *repository.CategoryRepository
	CategoryPublisher  *messaging.Publisher
}

func NewCategoryUseCase(db *gorm.DB, log *logrus.Logger, validator *validator.Validate, helpers *helpers.Helpers,
	categoryRepository *repository.CategoryRepository, categoryPublisher *messaging.Publisher) *CategoryUseCase {
	return &CategoryUseCase{
		DB:                 db,
		Log:                log,
		Validator:          validator,
		Helpers:            helpers,
		CategoryRepository: categoryRepository,
		CategoryPublisher: categoryPublisher,
	}
}

func (c *CategoryUseCase) ListCategory(ctx context.Context, categoryRequest *model.ListCategoryRequest) ([]model.CategoryResponse, int64, error) {
	tx := c.DB.WithContext(ctx).Begin()
	defer tx.Rollback()

	if err := c.Validator.Struct(categoryRequest); err != nil {
		c.Log.WithError(err).Error("error validating request body")
		return nil, 0, echo.ErrBadRequest
	}

	listCategories, total, err := c.CategoryRepository.ListCategory(tx, categoryRequest)
	if err != nil {
		c.Log.WithError(err).Error("error listing categories")
		return nil, 0, echo.ErrInternalServerError
	}

	if err := tx.Commit().Error; err != nil {
		c.Log.WithError(err).Error("error listing categories")
		return nil, 0, echo.ErrInternalServerError
	}

	responses := make([]model.CategoryResponse, len(listCategories))
	for i, category := range listCategories {
		responses[i] = *converter.CategoryToResponse(&category)
	}

	return responses, total, nil
}

func (c *CategoryUseCase) CreateCategory(ctx context.Context, categoryRequest *model.CreateCategoryRequest) (*model.CategoryResponse, error) {
	tx := c.DB.WithContext(ctx).Begin()
	defer tx.Rollback()

	if err := validator.New().Struct(categoryRequest); err != nil {
		c.Log.WithError(err).Error("error validating request body")
		return nil, echo.ErrBadRequest
	}

	category := &entity.Category{
		Nama: c.Helpers.ToTitle(categoryRequest.Nama),
		Slug: c.Helpers.ToSlug(categoryRequest.Nama),
	}

	categoryID, err := c.CategoryRepository.CreateCategory(tx, category)
	if err != nil {
		c.Log.WithError(err).Error("error creating category")
		return nil, echo.ErrInternalServerError
	}

	if err := tx.Commit().Error; err != nil {
		c.Log.WithError(err).Error("error creating category")
		return nil, echo.ErrInternalServerError
	}

	category.ID = categoryID
	event := converter.CategoryToEvent(category)
	c.CategoryPublisher.Send("category.created", event)

	response := converter.CategoryToResponse(category)
	
	return response, nil
}

func (c *CategoryUseCase) UpdateCategory(ctx context.Context, categoryRequest *model.UpdateCategoryRequest) (*model.CategoryResponse, error) {
	tx := c.DB.WithContext(ctx).Begin()
	defer tx.Rollback()

	if err := validator.New().Struct(categoryRequest); err != nil {
		c.Log.WithError(err).Error("error validating request body")
		return nil, echo.ErrBadRequest
	}

	category := new(entity.Category)
	resultCategory, err := c.CategoryRepository.GetCategoryBySlug(tx, category, categoryRequest.Slug)
	if err != nil {
		c.Log.WithError(err).Error("error getting category")
		return nil, echo.ErrNotFound
	}
	
	category.ID = resultCategory.ID
	category.Nama = c.Helpers.ToTitle(categoryRequest.Nama)
	category.Slug = c.Helpers.ToSlug(categoryRequest.Nama)
	
	if err := c.CategoryRepository.UpdateCategory(tx, category); err != nil {
		c.Log.WithError(err).Error("error updating category")
		return nil, echo.ErrInternalServerError
	}

	if err := tx.Commit().Error; err != nil {
		c.Log.WithError(err).Error("error updating category")
		return nil, echo.ErrInternalServerError
	}

	event := converter.CategoryToEvent(category)
	c.CategoryPublisher.Send("category.updated", event)

	response := converter.CategoryToResponse(category)

	return response, nil
}

func (c *CategoryUseCase) DeleteCategory(ctx context.Context, categoryRequest *model.DeleteCategoryRequest) (*model.CategoryResponse, error) {
	tx := c.DB.WithContext(ctx).Begin()
	defer tx.Rollback()

	if err := c.Validator.Struct(categoryRequest); err != nil {
		c.Log.WithError(err).Error("error validating request body")
		return nil, echo.ErrBadRequest
	}

	var category *entity.Category

	resultCategory, err := c.CategoryRepository.GetCategoryBySlug(tx, category, categoryRequest.Slug)
	if err != nil {
		c.Log.WithError(err).Error("error getting category")
		return nil, echo.ErrNotFound
	}

	if err := c.CategoryRepository.DeleteCategory(tx, category, categoryRequest.Slug); err != nil {
		c.Log.WithError(err).Error("error deleting category")
		return nil, echo.ErrInternalServerError
	}

	if err := tx.Commit().Error; err != nil {
		c.Log.WithError(err).Error("error deleting category")
		return nil, echo.ErrInternalServerError
	}

	event := converter.CategoryToEvent(resultCategory)
	c.CategoryPublisher.Send("category.deleted", event)

	response := converter.CategoryToResponse(resultCategory)

	return response, nil
}

func (c *CategoryUseCase) GetCategoryBySlug(ctx context.Context, categoryRequest *model.GetCategoryBySlugRequest) (*entity.Category, error) {
	tx := c.DB.WithContext(ctx).Begin()
	defer tx.Rollback()

	var category *entity.Category

	getCategory, err := c.CategoryRepository.GetCategoryBySlug(tx, category, categoryRequest.Slug)
	if err != nil {
		c.Log.WithError(err).Error("error getting category")
		return nil, echo.ErrNotFound
	}

	return getCategory, nil
}
