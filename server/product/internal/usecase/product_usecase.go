package usecase

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/NetSinx/yconnect-shop/server/product/internal/entity"
	"github.com/NetSinx/yconnect-shop/server/product/internal/gateway/messaging"
	"github.com/NetSinx/yconnect-shop/server/product/internal/helpers"
	"github.com/NetSinx/yconnect-shop/server/product/internal/model"
	"github.com/NetSinx/yconnect-shop/server/product/internal/model/converter"
	"github.com/NetSinx/yconnect-shop/server/product/internal/repository"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/redis/go-redis/v9"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"gorm.io/gorm"
	"math"
	"time"
)

type ProductUseCase struct {
	Config            *viper.Viper
	DB                *gorm.DB
	Log               *logrus.Logger
	Validator         *validator.Validate
	RedisClient       *redis.Client
	Publisher         *messaging.Publisher
	ProductRepository *repository.ProductRepository
}

func NewProductUseCase(config *viper.Viper, db *gorm.DB, log *logrus.Logger, redisClient *redis.Client, publisher *messaging.Publisher, productRepository *repository.ProductRepository) *ProductUseCase {
	return &ProductUseCase{
		Config:            config,
		DB:                db,
		Log:               log,
		RedisClient:       redisClient,
		Publisher:         publisher,
		ProductRepository: productRepository,
	}
}

func (p *ProductUseCase) GetAllProduct(ctx context.Context, productReq *model.GetAllProductRequest) (*model.DataResponse[[]model.ProductResponse], error) {
	tx := p.DB.WithContext(ctx).Begin()
	defer tx.Rollback()

	if err := p.Validator.Struct(productReq); err != nil {
		p.Log.WithError(err).Error("error validating request")
		return nil, echo.ErrBadRequest
	}

	if productReq.Page <= 0 {
		productReq.Page = 1
	}

	if productReq.Size <= 0 {
		productReq.Size = 20
	}

	result, err := p.RedisClient.Get(ctx, fmt.Sprintf("products:%d:%d", productReq.Page, productReq.Size)).Result()
	if err == nil {
		products := new(model.DataResponse[[]model.ProductResponse])
		if err := json.Unmarshal([]byte(result), products); err != nil {
			p.Log.WithError(err).Error("error unmarshaling data")
			return nil, echo.ErrInternalServerError
		}

		return products, nil
	}

	entityProduct := new([]entity.Product)
	totalProduct, err := p.ProductRepository.GetAll(tx, *entityProduct, productReq)
	if err != nil {
		p.Log.WithError(err).Error("error getting all products")
		return nil, echo.ErrInternalServerError
	}

	if err := tx.Commit().Error; err != nil {
		p.Log.WithError(err).Error("error getting all products")
		return nil, echo.ErrInternalServerError
	}

	getAllResponse := make([]model.ProductResponse, len(*entityProduct))
	for i, product := range *entityProduct {
		getAllResponse[i] = *converter.ProductToResponse(&product)
	}

	response := &model.DataResponse[[]model.ProductResponse]{
		Data: getAllResponse,
		PageMetadata: &model.PageMetadataResponse{
			Page:      productReq.Page,
			Size:      productReq.Size,
			TotalItem: totalProduct,
			TotalPage: int64(math.Ceil(float64(totalProduct) / float64(productReq.Size))),
		},
	}

	bytesData, err := json.Marshal(response)
	if err != nil {
		p.Log.WithError(err).Error("error marshaling data")
		return nil, echo.ErrInternalServerError
	}

	if err := p.RedisClient.Set(ctx, fmt.Sprintf("products:%d:%d", productReq.Page, productReq.Size), bytesData, 5*time.Minute).Err(); err != nil {
		p.Log.WithError(err).Error("error setting cache data in redis")
		return nil, echo.ErrInternalServerError
	}

	return response, nil
}

func (p *ProductUseCase) CreateProduct(ctx context.Context, productReq *model.ProductRequest) error {
	tx := p.DB.WithContext(ctx).Begin()
	defer tx.Rollback()

	if err := validator.New().Struct(productReq); err != nil {
		p.Log.WithError(err).Error("error validating request body")
		return echo.ErrBadRequest
	}

	categoryMirror := new(entity.CategoryMirror)
	if err := p.ProductRepository.GetCategoryMirror(tx, categoryMirror, productReq.KategoriSlug); err != nil {
		p.Log.WithError(err).Error("error getting category mirror")
		return echo.ErrNotFound
	}

	slug, err := helpers.GenerateSlugByName(productReq.Nama)
	if err != nil {
		p.Log.WithError(err).Error("error generating slug")
		return echo.ErrInternalServerError
	}

	product := &entity.Product{
		Nama:         productReq.Nama,
		Deskripsi:    productReq.Deskripsi,
		Slug:         slug,
		Gambar:       productReq.Gambar,
		KategoriSlug: productReq.KategoriSlug,
		Harga:        productReq.Harga,
		Stok:         productReq.Stok,
	}

	if err = p.ProductRepository.Create(tx, product); err != nil {
		p.Log.WithError(err).Error("error creating product")
		return echo.ErrInternalServerError
	}

	if err := tx.Commit().Error; err != nil {
		p.Log.WithError(err).Error("error creating product")
		return echo.ErrInternalServerError
	}

	eventCreated := converter.ProductToEvent(product)
	p.Publisher.Send("product.created", eventCreated)

	return nil
}

func (p *ProductUseCase) UpdateProduct(ctx context.Context, productReq *model.ProductRequest, slug string) error {
	tx := p.DB.WithContext(ctx).Begin()
	defer tx.Rollback()

	if err := p.Validator.Struct(productReq); err != nil {
		p.Log.WithError(err).Error("error validating request body")
		return echo.ErrBadRequest
	}

	categoryMirror := new(entity.CategoryMirror)
	if err := p.ProductRepository.GetCategoryMirror(tx, categoryMirror, productReq.KategoriSlug); err != nil {
		p.Log.WithError(err).Error("error getting category mirror")
		return echo.ErrNotFound
	}

	product := new(entity.Product)
	if err := p.ProductRepository.GetProductName(tx, product, slug); err != nil {
		p.Log.WithError(err).Error("error getting product name")
		return echo.ErrNotFound
	}

	if product.Nama != productReq.Nama {
		newSlug := helpers.ReplaceProductSlug(product.Slug, productReq.Nama)
		product.Nama = productReq.Nama
		product.Slug = newSlug
	}

	product.Gambar = productReq.Gambar
	product.Deskripsi = productReq.Deskripsi
	product.KategoriSlug = productReq.KategoriSlug
	product.Harga = productReq.Harga
	product.Stok = productReq.Stok

	if err := p.ProductRepository.Update(tx, product, slug); err != nil {
		p.Log.WithError(err).Error("error updating product")
		return echo.ErrInternalServerError
	}

	if err := tx.Commit().Error; err != nil {
		p.Log.WithError(err).Error("error updating product")
		return echo.ErrInternalServerError
	}

	eventUpdated := converter.ProductToEvent(product)
	p.Publisher.Send("product.updated", eventUpdated)

	return nil
}

func (p *ProductUseCase) DeleteProduct(ctx context.Context, slug string) error {
	tx := p.DB.WithContext(ctx).Begin()
	defer tx.Rollback()

	product := new(entity.Product)
	if err := p.ProductRepository.DeleteProduct(tx, product, slug); err != nil {
		p.Log.WithError(err).Error("error deleting product")
		return echo.ErrNotFound
	}

	if err := tx.Commit().Error; err != nil {
		p.Log.WithError(err).Error("error deleting product")
		return echo.ErrInternalServerError
	}

	eventDeleted := converter.ProductToEvent(product)
	p.Publisher.Send("product.deleted", eventDeleted)

	return nil
}

func (p *ProductUseCase) GetProductBySlug(ctx context.Context, slug string) (*model.DataResponse[*model.ProductResponse], error) {
	tx := p.DB.WithContext(ctx).Begin()
	defer tx.Rollback()

	result, err := p.RedisClient.Get(ctx, "product:"+slug).Result()
	if err == nil {
		response := new(model.DataResponse[*model.ProductResponse])
		if err := json.Unmarshal([]byte(result), response); err != nil {
			p.Log.WithError(err).Error("error unmarshaling data")
			return nil, echo.ErrInternalServerError
		}

		return response, nil
	}

	product := new(entity.Product)
	if err := p.ProductRepository.GetProductBySlug(tx, product, slug); err != nil {
		p.Log.WithError(err).Error("error getting product")
		return nil, echo.ErrNotFound
	}

	if err := tx.Commit().Error; err != nil {
		p.Log.WithError(err).Error("error getting product")
		return nil, echo.ErrInternalServerError
	}

	response := &model.DataResponse[*model.ProductResponse]{
		Data: converter.ProductToResponse(product),
	}

	dataBytes, err := json.Marshal(response)
	if err != nil {
		p.Log.WithError(err).Error("error marshaling data")
		return nil, echo.ErrInternalServerError
	}

	if err := p.RedisClient.Set(ctx, "product:"+slug, dataBytes, 5*time.Minute).Err(); err != nil {
		p.Log.WithError(err).Error("error setting cache data")
		return nil, echo.ErrInternalServerError
	}

	return response, nil
}

func (p *ProductUseCase) GetCategoryProduct(ctx context.Context, slug string) (*model.DataResponse[*model.CategoryMirrorResponse], error) {
	tx := p.DB.WithContext(ctx).Begin()
	defer tx.Rollback()

	result, err := p.RedisClient.Get(ctx, "category_product:"+slug).Result()
	if err == nil {
		response := new(model.DataResponse[*model.CategoryMirrorResponse])
		if err := json.Unmarshal([]byte(result), response); err != nil {
			p.Log.WithError(err).Error("error unmarshaling data")
			return nil, echo.ErrInternalServerError
		}

		return response, nil
	}

	categoryMirror := new(entity.CategoryMirror)
	if err := p.ProductRepository.GetCategoryProduct(tx, categoryMirror, slug); err != nil {
		p.Log.WithError(err).Error("error getting category product")
		return nil, echo.ErrNotFound
	}

	response := &model.DataResponse[*model.CategoryMirrorResponse]{
		Data: converter.CategoryMirrorToResponse(categoryMirror),
	}

	dataBytes, err := json.Marshal(response)
	if err != nil {
		p.Log.WithError(err).Error("error marshaling data")
		return nil, echo.ErrInternalServerError
	}

	if err := p.RedisClient.Set(ctx, "category_product:"+slug, dataBytes, 5*time.Minute).Err(); err != nil {
		p.Log.WithError(err).Error("error setting cache data")
		return nil, echo.ErrInternalServerError
	}

	return response, nil
}

func (p *ProductUseCase) GetProductByCategory(ctx context.Context, productReq *model.GetAllProductRequest, slug string) (*model.DataResponse[[]model.ProductResponse], error) {
	tx := p.DB.WithContext(ctx).Begin()
	defer tx.Rollback()

	key := fmt.Sprintf("products:%s:%d:%d", slug, productReq.Page, productReq.Size)
	result, err := p.RedisClient.Get(ctx, key).Result()
	if err == nil {
		response := new(model.DataResponse[[]model.ProductResponse])
		if err := json.Unmarshal([]byte(result), response); err != nil {
			p.Log.WithError(err).Error("error unmarshaling data")
			return nil, echo.ErrInternalServerError
		}

		return response, nil
	}

	product := new([]entity.Product)
	if err := p.ProductRepository.GetProductByCategory(tx, *product, productReq, slug); err != nil {
		p.Log.WithError(err).Error("error getting product")
		return nil, echo.ErrNotFound
	}

	products := make([]model.ProductResponse, len(*product))
	for i, p := range *product {
		products[i] = *converter.ProductToResponse(&p)
	}

	response := &model.DataResponse[[]model.ProductResponse]{
		Data: products,
	}

	dataBytes, err := json.Marshal(response)
	if err != nil {
		p.Log.WithError(err).Error("error marshaling data")
		return nil, echo.ErrInternalServerError
	}

	if err := p.RedisClient.Set(ctx, key, dataBytes, 5*time.Minute).Err(); err != nil {
		p.Log.WithError(err).Error("error setting cache data")
		return nil, echo.ErrInternalServerError
	}

	return response, nil
}

func (p *ProductUseCase) CreateCategoryMirror(ctx context.Context, categoryEvent *model.CategoryEvent) error {
	tx := p.DB.WithContext(ctx).Begin()
	defer tx.Rollback()

	
}
