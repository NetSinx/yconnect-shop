package repository

import (
	"github.com/NetSinx/yconnect-shop/product/app/model"
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

func (p productRepository) ListProduct(products []model.Product) ([]model.Product, error) {
	if err := p.DB.Select("id", "name", "slug", "description", "category_id", "seller_id", "price", "stock").Find(&products).Error; err != nil {
		return nil, err
	}

	return products, nil
}

func (p productRepository) CreateProduct(products model.Product) error {
	if err := p.DB.Create(&products).Error; err != nil {
		return err
	}

	return nil
}

func (p productRepository) UpdateProduct(products model.Product, slug string) error {
	if err := p.DB.Where("slug = ?", slug).Updates(&products).Error; err != nil {
		return err
	}

	if err := p.DB.First(&products, "slug = ?", slug).Error; err != nil {
		return err
	}

	return nil
}

func (p productRepository) DeleteProduct(products model.Product, slug string) error {
	if err := p.DB.Where("slug = ?", slug).Delete(&products).Error; err != nil {
		return err
	}

	return nil
}

func (p productRepository) GetProduct(products model.Product, slug string) (model.Product, error) {
	if err := p.DB.Where("slug = ?", slug).First(&products).Error; err != nil {
		return products, err
	}

	return products, nil
}