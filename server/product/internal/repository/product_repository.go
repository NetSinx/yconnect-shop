package repository

import (
	"github.com/NetSinx/yconnect-shop/server/product/internal/entity"
	"github.com/NetSinx/yconnect-shop/server/product/internal/model"
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

func (p *ProductRepository) GetAll(db *gorm.DB, entityProduct []entity.Product, productReq *model.GetAllProductRequest) (int64, error) {
	if err := db.Offset((productReq.Page - 1) * productReq.Size).Limit(productReq.Size).Preload("Gambar").Find(entityProduct).Error; err != nil {
		return 0, err
	}

	var total int64
	if err := db.Model(&entity.Product{}).Count(&total).Error; err != nil {
		return 0, err
	}

	return total, nil
}

func (p *ProductRepository) Create(db *gorm.DB, entity *entity.Product) error {
	if err := db.Create(entity).Error; err != nil {
		return err
	}

	return nil
}

func (p *ProductRepository) Update(db *gorm.DB, entityProduct *entity.Product, entityGambar []entity.Gambar, slug string) error {
	if err := db.Where("slug = ?", slug).First(entityProduct).Error; err != nil {
		return err
	}

	if err := db.Where("product_id = ?", entityProduct.ID).Delete(entityGambar).Error; err != nil {
		return err
	}

	if err := db.Where("slug = ?", slug).Updates(entityProduct).Error; err != nil {
		return err
	}

	if err := db.Model(entityProduct).Association("Gambar").Replace(entityGambar); err != nil {
		return err
	}

	return nil
}

func (p *ProductRepository) DeleteProduct(db *gorm.DB, entityProduct *entity.Product, entityGambar *entity.Gambar, slug string) error {
	if err := db.Where("slug = ?", slug).First(entityProduct).Error; err != nil {
		return err
	}
	
	if err := db.Where("product_id = ?", entityProduct.ID).Delete(entityGambar).Error; err != nil {
		return err
	}

	if err := db.Delete(entityProduct).Error; err != nil {
		return err
	}

	return nil
}

func (p *ProductRepository) GetProductBySlug(db *gorm.DB, entity *entity.Product, slug string) (*entity.Product, error) {
	if err := db.Preload("Gambar").First(entity, "slug = ?", slug).Error; err != nil {
		return nil, err
	}

	return entity, nil
}

func (p *ProductRepository) GetCategoryProduct(db *gorm.DB, entity *entity.Product, categoryMirror *entity.CategoryMirror, slug string) (*entity.CategoryMirror, error) {
	if err := db.First(entity, "slug = ?", slug).Error; err != nil {
		return nil, err
	}

	if err := db.First(categoryMirror, "slug = ?", entity.KategoriSlug).Error; err != nil {
		return nil, err
	}

	return categoryMirror, nil
}

func (p *ProductRepository) GetProductName(db *gorm.DB, entity *entity.Product, slug string) (*entity.Product, error) {
	if err := db.Select("nama", "slug").First(entity, "slug = ?", slug).Error; err != nil {
		return nil, err
	}

	return entity, nil
}

func (p *ProductRepository) GetProductByCategory(db *gorm.DB, entity []entity.Product, slug string) ([]entity.Product, error) {
	if err := db.Preload("Gambar").Find(entity, "kategori_slug = ?", slug).Error; err != nil {
		return nil, err
	}

	return entity, nil
}

func (p *ProductRepository) GetCategoryMirror(db *gorm.DB, categoryMirror *entity.CategoryMirror, slug string) error {
	if err := db.First(categoryMirror, "slug = ?", slug).Error; err != nil {
		return err
	}

	return nil
}