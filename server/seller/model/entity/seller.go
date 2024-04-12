package entity

import (
	"time"
	"github.com/NetSinx/yconnect-shop/server/product/app/model"
)

type Seller struct {
	Id        uint             `json:"id" gorm:"primaryKey"`
	Name      string           `json:"name" gorm:"unique" validate:"required,min=5"`
	Username  string           `json:"username" gorm:"unique"`
	Avatar    string           `json:"avatar"`
	Email     string           `json:"email" gorm:"unique"`
	Alamat    string           `json:"alamat"`
	NoTelp    string           `json:"no_telp" gorm:"unique"`
	Product   []model.Product  `json:"product" gorm:"-"`
	UserID    uint             `json:"user_id"`
	CreatedAt time.Time
	UpdatedAt time.Time
}