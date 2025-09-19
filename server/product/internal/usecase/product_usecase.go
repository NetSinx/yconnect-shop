package usecase

import (
	"github.com/NetSinx/yconnect-shop/server/product/internal/repository"
	"github.com/go-playground/validator/v10"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type ProductUseCase struct {
	DB                *gorm.DB
	Log               *logrus.Logger
	productRepository *repository.ProductRepository
}

func NewProductService(db *gorm.DB, log *logrus.Logger, productRepository *repository.ProductRepository) *ProductUseCase {
	return &ProductUseCase{
		DB: db,
		Log: log,
		productRepository: productRepository,
	}
}

func (p *ProductUseCase) GetAllProduct(product []model.Product) ([]model.Product, error) {
	listProduct, err := p.productRepository.ListProduct(products)
	if err != nil {
		return products, err
	}

	return listProduct, nil
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
