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

func (p productRepository) CreateProduct(products entity.Product, image []entity.Image) (entity.Product, error) {
	if err := p.DB.Create(&entity.Product{Name: products.Name, Slug: products.Slug, Description: products.Description, Image: image, SellerId: products.SellerId, CategoryId: products.CategoryId, Price: products.Price, Stock: products.Stock}).Error; err != nil {
		return products, err
	}

	return products, nil
}

func (p productRepository) UpdateProduct(products entity.Product, image []entity.Image, slug string, id string) (entity.Product, error) {
	err := p.DB.Model(&entity.Image{}).Where("product_id = ?", id).Save(&image).Error
	if err != nil {
		return products, err
	}

	err = p.DB.Where("slug = ?", slug).Updates(&entity.Product{Name: products.Name, Slug: products.Slug, Description: products.Description, SellerId: products.SellerId, CategoryId: products.CategoryId, Price: products.Price, Stock: products.Stock}).Error
	if err != nil {
		return products, err
	}

	if err := p.DB.First(&products, "slug = ?", slug).Error; err != nil {
		return products, err
	}

	return products, nil
}

func (p productRepository) DeleteProduct(products entity.Product, image []entity.Image, slug string, id string) error {
	if err := p.DB.Preload("Image").First(&products, "slug = ?", slug).Error; err != nil {
		return err
	}

	p.DB.Delete(&image, "product_id = ?", id)

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

func (p productRepository) GetProductBySeller(products []entity.Product, id string) ([]entity.Product, error) {
	if err := p.DB.Preload("Image").Find(&products, "seller_id = ?", id).Error; err != nil {
		return nil, err
	}

	return products, nil
}