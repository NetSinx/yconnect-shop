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
	GetCategoryProduct(product model.Product, slug string) (model.CategoryMirror, error)
	GetProductName(product model.Product, slug string) (model.Product, error)
	GetProductByCategory(product []model.Product, slug string) ([]model.Product, error)
	GetMirrorCategory(slug string) error
}

type productRepository struct {
	DB *gorm.DB
}

func NewProductRepository(db *gorm.DB) *productRepository {
	return &productRepository{
		DB: db,
	}
}

func (p *productRepository) ListProduct(products []model.Product) ([]model.Product, error) {
	if err := p.DB.Preload("Gambar").Find(&products).Error; err != nil {
		return nil, err
	}

	return products, nil
}

func (p *productRepository) CreateProduct(product model.Product) error {
	if err := p.DB.Create(&product).Error; err != nil {
		return err
	}

	return nil
}

func (p *productRepository) UpdateProduct(productReq model.Product, gambar []model.Gambar, slug string) error {
	var product model.Product

	if err := p.DB.Where("slug = ?", slug).First(&product).Error; err != nil {
		return err
	}

	if err := p.DB.Where("product_id = ?", product.Id).Delete(&gambar).Error; err != nil {
		return err
	}

	if err := p.DB.Where("slug = ?", slug).Updates(&productReq).Error; err != nil {
		return err
	}

	if err := p.DB.Model(&product).Association("Gambar").Replace(gambar); err != nil {
		return err
	}

	return nil
}

func (p *productRepository) DeleteProduct(product model.Product, slug string) error {
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

func (p *productRepository) GetProductByID(product model.Product, id string) (model.Product, error) {
	if err := p.DB.Preload("Gambar").First(&product, "id = ?", id).Error; err != nil {
		return product, err
	}

	return product, nil
}

func (p *productRepository) GetProductBySlug(product model.Product, slug string) (model.Product, error) {
	if err := p.DB.Preload("Gambar").First(&product, "slug = ?", slug).Error; err != nil {
		return product, err
	}

	return product, nil
}

func (p *productRepository) GetCategoryProduct(product model.Product, slug string) (model.CategoryMirror, error) {
	if err := p.DB.First(&product, "slug = ?", slug).Error; err != nil {
		return model.CategoryMirror{}, err
	}

	var categoryMirror model.CategoryMirror

	if err := p.DB.First(&categoryMirror, "slug = ?", product.KategoriSlug).Error; err != nil {
		return categoryMirror, err
	}

	return categoryMirror, nil
}

func (p *productRepository) GetProductName(product model.Product, slug string) (model.Product, error) {
	if err := p.DB.Select("nama", "slug").First(&product, "slug = ?", slug).Error; err != nil {
		return product, err
	}

	return product, nil
}

func (p *productRepository) GetProductByCategory(products []model.Product, slug string) ([]model.Product, error) {
	if err := p.DB.Preload("Gambar").Find(&products, "kategori_slug = ?", slug).Error; err != nil {
		return products, err
	}

	return products, nil
}

func (p *productRepository) GetMirrorCategory(slug string) error {
	var categoryMirror model.CategoryMirror

	if err := p.DB.First(&categoryMirror, "slug = ?", slug).Error; err != nil {
		return err
	}

	return nil
}