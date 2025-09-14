package usecase

import (
	"context"
	"encoding/json"
	"fmt"
	"math"
	"time"
	"github.com/NetSinx/yconnect-shop/server/category/internal/entity"
	"github.com/NetSinx/yconnect-shop/server/category/internal/gateway/messaging"
	"github.com/NetSinx/yconnect-shop/server/category/internal/helpers"
	"github.com/NetSinx/yconnect-shop/server/category/internal/model"
	"github.com/NetSinx/yconnect-shop/server/category/internal/model/converter"
	"github.com/NetSinx/yconnect-shop/server/category/internal/repository"
	"github.com/go-playground/validator/v10"
	"github.com/go-redis/redis"
	"github.com/labstack/echo/v4"
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

func (c *CategoryUseCase) ListCategory(ctx context.Context, categoryRequest *model.ListCategoryRequest) (*model.DataResponse[[]model.CategoryResponse], error) {
	if err := c.Validator.Struct(categoryRequest); err != nil {
		c.Log.WithError(err).Error("error validating request body")
		return nil, echo.ErrBadRequest
	}
	
	if c.Config.GetBool("redis.enabled") {
		key := fmt.Sprintf("categoriesCache:%d:%d", categoryRequest.Page, categoryRequest.Size)
		result, err := c.RedisClient.Get(key).Result()
		if err == nil {
			categories := new(model.DataResponse[[]model.CategoryResponse])
			if err := json.Unmarshal([]byte(result), categories); err != nil {
				return nil, echo.ErrInternalServerError
			}
	
			return categories, nil
		}
	}
	
	tx := c.DB.WithContext(ctx).Begin()
	defer tx.Rollback()

	listCategories, total, err := c.CategoryRepository.ListCategory(tx, categoryRequest)
	if err != nil {
		c.Log.WithError(err).Error("error listing categories")
		return nil, echo.ErrInternalServerError
	}

	if err := tx.Commit().Error; err != nil {
		c.Log.WithError(err).Error("error listing categories")
		return nil, echo.ErrInternalServerError
	}

	listCategoryResponse := make([]model.CategoryResponse, len(listCategories))
	for i, category := range listCategories {
		listCategoryResponse[i] = *converter.CategoryToResponse(&category)
	}

	response := &model.DataResponse[[]model.CategoryResponse]{
		Data: listCategoryResponse,
		PageMetadata: &model.PageMetadataResponse{
			Page:      categoryRequest.Page,
			Size:      categoryRequest.Size,
			TotalItem: total,
			TotalPage: int64(math.Ceil(float64(total) / float64(categoryRequest.Size))),
		},
	}

	if c.Config.GetBool("redis.enabled") {
		categories, _ := json.Marshal(&response)
		key := fmt.Sprintf("categoriesCache:%d:%d", categoryRequest.Page, categoryRequest.Size)
		if err := c.RedisClient.Set(key, categories, 5*time.Minute).Err(); err != nil {
			c.Log.WithError(err).Error("error caching categories in redis")
			return nil, echo.ErrInternalServerError
		}
	}

	return response, nil
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

	if err := c.CategoryRepository.CreateCategory(tx, category); err != nil {
		c.Log.WithError(err).Error("error creating category")
		return nil, echo.ErrInternalServerError
	}

	if err := tx.Commit().Error; err != nil {
		c.Log.WithError(err).Error("error creating category")
		return nil, echo.ErrInternalServerError
	}

	if c.Config.GetBool("rabbitmq.enabled") {
		event := converter.CategoryToEvent(category)
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
	if err := c.CategoryRepository.GetCategoryBySlug(tx, category, categoryRequest.Slug); err != nil {
		c.Log.WithError(err).Error("error getting category")
		return nil, echo.ErrNotFound
	}

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

	if c.Config.GetBool("rabbitmq.enabled") {
		event := converter.CategoryToEvent(category)
		c.CategoryPublisher.Send("category.updated", event)
	}

	response := converter.CategoryToResponse(category)

	return response, nil
}

func (c *CategoryUseCase) DeleteCategory(ctx context.Context, categoryRequest *model.DeleteCategoryRequest) error {
	tx := c.DB.WithContext(ctx).Begin()
	defer tx.Rollback()

	if err := c.Validator.Struct(categoryRequest); err != nil {
		c.Log.WithError(err).Error("error validating request body")
		return echo.ErrBadRequest
	}

	category := new(entity.Category)
	if err := c.CategoryRepository.GetCategoryBySlug(tx, category, categoryRequest.Slug); err != nil {
		c.Log.WithError(err).Error("error getting category")
		return echo.ErrNotFound
	}

	if err := c.CategoryRepository.DeleteCategory(tx, category, categoryRequest.Slug); err != nil {
		c.Log.WithError(err).Error("error deleting category")
		return echo.ErrInternalServerError
	}

	if err := tx.Commit().Error; err != nil {
		c.Log.WithError(err).Error("error deleting category")
		return echo.ErrInternalServerError
	}

	if c.Config.GetBool("rabbitmq.enabled") {
		event := converter.CategoryToEvent(category)
		c.CategoryPublisher.Send("category.deleted", event)
	}

	return nil
}

func (c *CategoryUseCase) GetCategoryBySlug(ctx context.Context, categoryRequest *model.GetCategoryBySlugRequest) (*model.CategoryResponse, error) {
	if err := c.Validator.Struct(categoryRequest); err != nil {
		c.Log.WithError(err).Error("error validating request body")
		return nil, echo.ErrBadRequest
	}
	
	if c.Config.GetBool("redis.enabled") {
		result, err := c.RedisClient.Get("categoryCache:"+categoryRequest.Slug).Result()
		if err == nil {
			categoryResponse := new(model.CategoryResponse)
			if err := json.Unmarshal([]byte(result), categoryResponse); err != nil {
				c.Log.WithError(err).Error("error unmarshaling data")
				return nil, echo.ErrInternalServerError
			}
	
			return categoryResponse, nil
		}
	}
	
	tx := c.DB.WithContext(ctx).Begin()
	defer tx.Rollback()
	
	category := new(entity.Category)
	if err := c.CategoryRepository.GetCategoryBySlug(tx, category, categoryRequest.Slug); err != nil {
		c.Log.WithError(err).Error("error getting category")
		return nil, echo.ErrNotFound
	}

	if err := tx.Commit().Error; err != nil {
		c.Log.WithError(err).Error("error getting category")
		return nil, echo.ErrInternalServerError
	}

	if c.Config.GetBool("redis.enabled") {
		categoryByte, _ := json.Marshal(category)
		if err := c.RedisClient.Set("categoryCache:"+categoryRequest.Slug, categoryByte, 5*time.Minute).Err(); err != nil {
			c.Log.WithError(err).Error("error caching category in redis")
			return nil, echo.ErrInternalServerError
		}
	}

	response := converter.CategoryToResponse(category)

	return response, nil
}
