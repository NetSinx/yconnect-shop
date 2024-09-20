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
	if err := p.DB.Preload("Image").Find(&products).Error; err != nil {
		return nil, err
	}

	return products, nil
}

func (p productRepository) CreateProduct(products entity.Product) (entity.Product, error) {
	if err := p.DB.Create(&products).Error; err != nil {
		return products, err
	}

	return products, nil
}

func (p productRepository) UpdateProduct(products entity.Product, slug string) (entity.Product, error) {
	if err := p.DB.First(&products, "slug = ?", slug).Error; err != nil {
		return products, err
	}
	
	if err := p.DB.Updates(&products).Error; err != nil {
		return products, err
	}

	return products, nil
}

func (p productRepository) DeleteProduct(products entity.Product, slug string) error {
	if err := p.DB.First(&products, "slug = ?", slug).Error; err != nil {
		return err
	}

	if err := p.DB.Delete(&products).Error; err != nil {
		return err
	}

	return nil
}

func (p productRepository) GetProduct(products entity.Product, slug string) (entity.Product, error) {
	if err := p.DB.Where("slug = ?", slug).Preload("Image").First(&products).Error; err != nil {
		return products, err
	}

	return products, nil
}

func (p productRepository) GetProductByCategory(products []entity.Product, id string) ([]entity.Product, error) {
	if err := p.DB.Preload("Image").Find(&products, "category_id = ?", id).Error; err != nil {
		return nil, err
	}

	return products, nil
}