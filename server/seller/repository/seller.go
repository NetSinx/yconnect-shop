package repository

import (
	"github.com/NetSinx/yconnect-shop/server/seller/model/entity"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type sellerRepository struct {
	DB *gorm.DB
}

func SellerRepository(db *gorm.DB) sellerRepository {
	return sellerRepository{
		DB: db,
	}
}

func (sr sellerRepository) ListSeller() ([]entity.Seller, error) {
	var seller []entity.Seller

	if err := sr.DB.Find(&seller).Error; err != nil {
		return []entity.Seller{}, err
	}

	return seller, nil
}

func (sr sellerRepository) RegisterSeller(seller entity.Seller) (entity.Seller, error) {
	if err := sr.DB.Create(&seller).Error; err != nil {
		return entity.Seller{}, err
	}

	return seller, nil
}

func (sr sellerRepository) UpdateSeller(username string, seller entity.Seller) (entity.Seller, error) {
	if err := sr.DB.Clauses(clause.Returning{}).Where("username = ?", username).Updates(&seller).Error; err != nil {
		return entity.Seller{}, err
	}

	return seller, nil
}

func (sr sellerRepository) DeleteSeller(username string) error {
	var seller entity.Seller

	if err := sr.DB.First(&seller, "username = ?", username).Error; err != nil {
		return err
	}

	if err := sr.DB.Delete(&seller, "username = ?", username).Error; err != nil {
		return err
	}

	return nil
}

func (sr sellerRepository) GetSeller(username string) (entity.Seller, error) {
	var seller entity.Seller

	if err := sr.DB.First(&seller, "username = ?", username).Error; err != nil {
		return entity.Seller{}, err
	}

	return seller, nil
}