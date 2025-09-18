package repository

import (
	"github.com/NetSinx/yconnect-shop/server/product/internal/entity"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type ProductRepository struct {
	Log *logrus.Logger
}

func NewProductRepository(log *logrus.Logger) *ProductRepository {
	return &ProductRepository{
		Log: log,
	}
}

func (p *ProductRepository) GetAll(db *gorm.DB, entity *entity.Product) error {
	if err := db.Preload("Gambar").Find(entity).Error; err != nil {
		return err
	}

	return nil
}

func (p *ProductRepository) Create(db *gorm.DB, entity *entity.Product) error {
	if err := db.Create(entity).Error; err != nil {
		return err
	}

	return nil
}

func (p *ProductRepository) Update(productReq model.Product, gambar []model.Gambar, slug string) error {
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

func (p *ProductRepository) DeleteProduct(product model.Product, slug string) error {
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

func (p *ProductRepository) GetProductByID(product model.Product, id string) (model.Product, error) {
	if err := p.DB.Preload("Gambar").First(&product, "id = ?", id).Error; err != nil {
		return product, err
	}

	return product, nil
}

func (p *ProductRepository) GetProductBySlug(product model.Product, slug string) (model.Product, error) {
	if err := p.DB.Preload("Gambar").First(&product, "slug = ?", slug).Error; err != nil {
		return product, err
	}

	return product, nil
}

func (p *ProductRepository) GetCategoryProduct(product model.Product, slug string) (model.CategoryMirror, error) {
	if err := p.DB.First(&product, "slug = ?", slug).Error; err != nil {
		return model.CategoryMirror{}, err
	}

	var categoryMirror model.CategoryMirror

	if err := p.DB.First(&categoryMirror, "slug = ?", product.KategoriSlug).Error; err != nil {
		return categoryMirror, err
	}

	return categoryMirror, nil
}

func (p *ProductRepository) GetProductName(product model.Product, slug string) (model.Product, error) {
	if err := p.DB.Select("nama", "slug").First(&product, "slug = ?", slug).Error; err != nil {
		return product, err
	}

	return product, nil
}

func (p *ProductRepository) GetProductByCategory(products []model.Product, slug string) ([]model.Product, error) {
	if err := p.DB.Preload("Gambar").Find(&products, "kategori_slug = ?", slug).Error; err != nil {
		return products, err
	}

	return products, nil
}

func (p *ProductRepository) GetMirrorCategory(slug string) error {
	var categoryMirror model.CategoryMirror

	if err := p.DB.First(&categoryMirror, "slug = ?", slug).Error; err != nil {
		return err
	}

	return nil
}