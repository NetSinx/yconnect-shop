package repository

import (
	"github.com/NetSinx/yconnect-shop/server/product/model"
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
	if err := p.DB.Preload("Image").Find(&products).Error; err != nil {
		return nil, err
	}

	return products, nil
}

func (p productRepository) CreateProduct(products model.Product, image []model.Image) (model.Product, error) {
	if err := p.DB.Create(&model.Product{Name: products.Name, Slug: products.Slug, Description: products.Description, Image: image, SellerId: products.SellerId, CategoryId: products.CategoryId, Price: products.Price, Stock: products.Stock}).Error; err != nil {
		return products, err
	}

	return products, nil
}

func (p productRepository) UpdateProduct(products model.Product, image []model.Image, slug string, id string) (model.Product, error) {
	err := p.DB.Model(&model.Image{}).Where("product_id = ?", id).Save(&image).Error
	if err != nil {
		return products, err
	}

	err = p.DB.Where("slug = ?", slug).Updates(&model.Product{Name: products.Name, Slug: products.Slug, Description: products.Description, SellerId: products.SellerId, CategoryId: products.CategoryId, Price: products.Price, Stock: products.Stock}).Error
	if err != nil {
		return products, err
	}

	if err := p.DB.First(&products, "slug = ?", slug).Error; err != nil {
		return products, err
	}

	return products, nil
}

func (p productRepository) DeleteProduct(products model.Product, image []model.Image, slug string, id string) error {
	if err := p.DB.Preload("Image").First(&products, "slug = ?", slug).Error; err != nil {
		return err
	}

	p.DB.Delete(&image, "product_id = ?", id)

	if err := p.DB.Delete(&products).Error; err != nil {
		return err
	}

	return nil
}

func (p productRepository) GetProduct(products model.Product, slug string) (model.Product, error) {
	if err := p.DB.Where("slug = ?", slug).Preload("Image").First(&products).Error; err != nil {
		return products, err
	}

	return products, nil
}

func (p productRepository) GetProductByCategory(products []model.Product, id string) ([]model.Product, error) {
	if err := p.DB.Preload("Image").Find(&products, "category_id = ?", id).Error; err != nil {
		return nil, err
	}

	return products, nil
}

func (p productRepository) GetProductBySeller(products []model.Product, id string) ([]model.Product, error) {
	if err := p.DB.Preload("Image").Find(&products, "seller_id = ?", id).Error; err != nil {
		return nil, err
	}

	return products, nil
}