package usecase

import (
	"context"
	"math"
	"github.com/NetSinx/yconnect-shop/server/product/internal/entity"
	"github.com/NetSinx/yconnect-shop/server/product/internal/model"
	"github.com/NetSinx/yconnect-shop/server/product/internal/model/converter"
	"github.com/NetSinx/yconnect-shop/server/product/internal/repository"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/redis/go-redis/v9"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type ProductUseCase struct {
	DB                *gorm.DB
	Log               *logrus.Logger
	RedisClient       *redis.Client
	productRepository *repository.ProductRepository
}

func NewProductUseCase(db *gorm.DB, log *logrus.Logger, redisClient *redis.Client, productRepository *repository.ProductRepository) *ProductUseCase {
	return &ProductUseCase{
		DB: db,
		Log: log,
		RedisClient: redisClient,
		productRepository: productRepository,
	}
}

func (p *ProductUseCase) GetAllProduct(ctx context.Context, productReq *model.GetAllProductRequest) (*model.DataResponse[[]model.ProductResponse], error) {
	tx := p.DB.WithContext(ctx).Begin()
	defer tx.Rollback()

	if productReq.Page <= 0 {
		productReq.Page = 1
	}

	if productReq.Size <= 0 {
		productReq.Size = 20
	}

	var entityProduct []entity.Product
	totalProduct, err := p.productRepository.GetAll(tx, entityProduct, productReq)
	if err != nil {
		return nil, echo.ErrInternalServerError
	}

	getAllResponse := make([]model.ProductResponse, len(entityProduct))
	for i, product := range entityProduct {
		getAllResponse[i] = *converter.ProductToResponse(&product)
	}

	response := &model.DataResponse[[]model.ProductResponse]{
		Data: getAllResponse,
		PageMetadata: &model.PageMetadataResponse{
			Page: productReq.Page,
			Size: productReq.Size,
			TotalItem: totalProduct,
			TotalPage: int64(math.Ceil(float64(totalProduct) / float64(productReq.Size))),
		},
	}

	return response, nil
}

func (p *ProductUseCase) CreateProduct(productReq dto.ProductRequest) error {
	if err := validator.New().Struct(productReq); err != nil {
		return err
	}

	if err := p.productRepository.GetCategoryMirror(productReq.KategoriSlug); err != nil {
		return err
	}

	slug, err := helpers.GenerateSlugByName(productReq.Nama)
	if err != nil {
		return err
	}

	product := model.Product{
		Nama:         productReq.Nama,
		Deskripsi:    productReq.Deskripsi,
		Slug:         slug,
		Gambar:       productReq.Gambar,
		KategoriSlug: productReq.KategoriSlug,
		Harga:        productReq.Harga,
		Stok:         productReq.Stok,
	}

	err = p.productRepository.CreateProduct(product)
	if err != nil {
		return err
	}

	return nil
}

func (p *ProductUseCase) UpdateProduct(productReq dto.ProductRequest, slug string) error {
	if err := validator.New().Struct(productReq); err != nil {
		return err
	}

	if err := p.productRepository.GetMirrorCategory(productReq.KategoriSlug); err != nil {
		return err
	}

	product := model.Product{
		Nama:         productReq.Nama,
		Deskripsi:    productReq.Deskripsi,
		Slug:         slug,
		KategoriSlug: productReq.KategoriSlug,
		Harga:        productReq.Harga,
		Stok:         productReq.Stok,
	}

	gambar := productReq.Gambar

	getProductName, err := p.productRepository.GetProductName(product, slug)
	if err != nil {
		return err
	}

	if getProductName.Nama != productReq.Nama {
		newSlug := helpers.ReplaceProductSlug(getProductName.Slug, productReq.Nama)
		product.Slug = newSlug
	}

	err = p.productRepository.UpdateProduct(product, gambar, slug)
	if err != nil {
		return err
	}

	return nil
}

func (p *ProductUseCase) DeleteProduct(product model.Product, slug string) error {
	err := p.productRepository.DeleteProduct(product, slug)
	if err != nil {
		return err
	}

	return nil
}

func (p *ProductUseCase) GetProductByID(product model.Product, id string) (model.Product, error) {
	getProduct, err := p.productRepository.GetProductByID(product, id)
	if err != nil {
		return getProduct, err
	}

	return getProduct, nil
}

func (p *ProductUseCase) GetProductBySlug(product model.Product, slug string) (model.Product, error) {
	getProduct, err := p.productRepository.GetProductBySlug(product, slug)
	if err != nil {
		return getProduct, err
	}

	return getProduct, nil
}

func (p *ProductUseCase) GetCategoryProduct(product model.Product, slug string) (model.CategoryMirror, error) {
	categoryProduct, err := p.productRepository.GetCategoryProduct(product, slug)
	if err != nil {
		return categoryProduct, err
	}

	return categoryProduct, nil
}

func (p *ProductUseCase) GetProductByCategory(products []model.Product, slug string) ([]model.Product, error) {
	getProductByCategory, err := p.productRepository.GetProductByCategory(products, slug)
	if err != nil {
		return getProductByCategory, err
	}

	return getProductByCategory, nil
}
