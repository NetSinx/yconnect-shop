package repository

import (
	"github.com/NetSinx/yconnect-shop/server/seller/model/entity"
	"gorm.io/gorm"
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