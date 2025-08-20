package repository

import (
	"github.com/NetSinx/yconnect-shop/server/product/model"
	"gorm.io/gorm"
)

type ProductRepo interface {
	ListProduct(products []model.Product) ([]model.Product, error)
	CreateProduct(product model.Product) error
	UpdateProduct(product model.Product, gambar []model.Gambar, slug string) error
	DeleteProduct(product model.Product, slug string) error
	GetProductByID(product model.Product, id string) (model.Product, error)
	GetProductBySlug(product model.Product, slug string) (model.Product, error)
	GetCategoryProduct(product model.Product, slug string) error
	GetProductName(product model.Product, slug string) (model.Product, error)
}

type productRepository struct {
	DB *gorm.DB
}

func ProductRepository(db *gorm.DB) productRepository {
	return productRepository{
		DB: db,
	}
}

func (p productRepository) ListProduct(products []model.Product) ([]model.Product, error) {
	if err := p.DB.Preload("Gambar").Find(&products).Error; err != nil {
		return nil, err
	}

	return products, nil
}

func (p productRepository) CreateProduct(product model.Product) error {
	if err := p.DB.Create(&product).Error; err != nil {
		return err
	}

	return nil
}

func (p productRepository) UpdateProduct(product model.Product, gambar []model.Gambar, slug string) error {
	if err := p.DB.Where("slug = ?", slug).First(&product).Error; err != nil {
		return err
	}

	if err := p.DB.Where("product_id = ?", product.Id).Delete(&gambar).Error; err != nil {
		return err
	}

	if err := p.DB.Updates(&product).Error; err != nil {
		return err
	}

	if err := p.DB.Model(&product).Association("Gambar").Replace(product.Gambar, gambar); err != nil {
		return err
	}

	return nil
}

func (p productRepository) DeleteProduct(product model.Product, slug string) error {
	if err := p.DB.Where("slug = ?", slug).First(&product).Error; err != nil {
		return err
	}
	
	if err := p.DB.Where("product_id = ?", product.Id).Delete(&model.Gambar{}).Error; err != nil {
		return err
	}

	if err := p.DB.Delete(&product).Error; err != nil {
		return err
	}

	return nil
}

func (p productRepository) GetProductByID(product model.Product, id string) (model.Product, error) {
	if err := p.DB.Where("id = ?", id).Preload("Gambar").First(&product).Error; err != nil {
		return product, err
	}

	return product, nil
}

func (p productRepository) GetProductBySlug(product model.Product, slug string) (model.Product, error) {
	if err := p.DB.Where("slug = ?", slug).Preload("Gambar").First(&product).Error; err != nil {
		return product, err
	}

	return product, nil
}

func (p productRepository) GetCategoryProduct(product model.Product, slug string) error {
	if err := p.DB.Where("slug = ?", slug).First(&product).Error; err != nil {
		return err
	}

	return nil
}

func (p productRepository) GetProductName(product model.Product, slug string) (model.Product, error) {
	if err := p.DB.Select("nama").Where("slug = ?", slug).First(&product).Error; err != nil {
		return product, err
	}

	return product, nil
}