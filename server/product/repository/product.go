package repository

import (
	"github.com/NetSinx/yconnect-shop/server/product/model/entity"
	"gorm.io/gorm"
)

type productRepository struct {
	DB *gorm.DB
}

func ProductRepository(db *gorm.DB) productRepository {
	return productRepository{
		DB: db,
	}
}

func (p productRepository) ListProduct(products []entity.Product) ([]entity.Product, error) {
	if err := p.DB.Preload("Gambar").Find(&products).Error; err != nil {
		return nil, err
	}

	return products, nil
}

func (p productRepository) CreateProduct(product entity.Product) error {
	if err := p.DB.Create(&product).Error; err != nil {
		return err
	}

	return nil
}

func (p productRepository) UpdateProduct(product entity.Product, slug string) error {
	if err := p.DB.First(&product, "slug = ?", slug).Error; err != nil {
		return err
	}
	
	if err := p.DB.Save(&product).Error; err != nil {
		return err
	}

	return nil
}

func (p productRepository) DeleteProduct(product entity.Product, slug string) error {
	if err := p.DB.Where("slug = ?", slug).Delete(&product).Error; err != nil {
		return err
	}

	return nil
}

func (p productRepository) GetProductByID(product entity.Product, id string) (entity.Product, error) {
	if err := p.DB.Where("id = ?", id).Preload("Gambar").First(&product).Error; err != nil {
		return product, err
	}

	return product, nil
}

func (p productRepository) GetProductBySlug(product entity.Product, slug string) (entity.Product, error) {
	if err := p.DB.Where("slug = ?", slug).Preload("Gambar").First(&product).Error; err != nil {
		return product, err
	}

	return product, nil
}

func (p productRepository) GetCategoryProduct(product entity.Product, slug string) error {
	if err := p.DB.Where("slug = ?", slug).First(&product).Error; err != nil {
		return err
	}

	return nil
}