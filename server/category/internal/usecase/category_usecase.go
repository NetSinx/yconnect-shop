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
	"github.com/redis/go-redis"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"gorm.io/gorm"
)

type CategoryUseCase struct {
	Config             *viper.Viper
	DB                 *gorm.DB
	Log                *logrus.Logger
	RedisClient        *redis.Client
	Validator          *validator.Validate
	CategoryRepository *repository.CategoryRepository
	CategoryPublisher  *messaging.Publisher
}

func NewCategoryUseCase(config *viper.Viper, db *gorm.DB, log *logrus.Logger, redisClient *redis.Client, validator *validator.Validate,
	categoryRepository *repository.CategoryRepository, categoryPublisher *messaging.Publisher) *CategoryUseCase {
	return &CategoryUseCase{
		Config:             config,
		DB:                 db,
		Log:                log,
		RedisClient:        redisClient,
		Validator:          validator,
		CategoryRepository: categoryRepository,
		CategoryPublisher:  categoryPublisher,
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
		Nama: helpers.ToTitle(categoryRequest.Nama),
		Slug: helpers.ToSlug(categoryRequest.Nama),
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

	if c.Config.GetBool("rabbitmq.enabled") {
		c.CategoryPublisher.Send("category.created", event)
	}

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
	category.Nama = helpers.ToTitle(categoryRequest.Nama)
	category.Slug = helpers.ToSlug(categoryRequest.Nama)

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

func (c *CategoryUseCase) GetCategoryBySlug(ctx context.Context, categoryRequest *model.GetCategoryBySlugRequest) (*model.CategoryResponse, error) {
	tx := c.DB.WithContext(ctx).Begin()
	defer tx.Rollback()

	var category *entity.Category

	getCategory, err := c.CategoryRepository.GetCategoryBySlug(tx, category, categoryRequest.Slug)
	if err != nil {
		c.Log.WithError(err).Error("error getting category")
		return nil, echo.ErrNotFound
	}

	response := converter.CategoryToResponse(getCategory)

	return response, nil
}
