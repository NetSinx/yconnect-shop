package entity

import (
	"time"
	"github.com/NetSinx/yconnect-shop/server/product/app/model"
)

type Seller struct {
	Id        uint             `json:"id" gorm:"primaryKey"`
	Name      string           `json:"name" form:"name" gorm:"unique" validate:"required,min=5"`
	Username  string           `json:"username" gorm:"unique" validate:"required"`
	Avatar    string           `json:"avatar"`
	Email     string           `json:"email" gorm:"unique" validate:"required"`
	Alamat    string           `json:"alamat" validate:"required"`
	NoTelp    string           `json:"no_telp" gorm:"unique" validate:"required"`
	Product   []model.Product  `json:"product" gorm:"-"`
	UserID    uint             `json:"user_id" validate:"required"`
	CreatedAt time.Time
	UpdatedAt time.Time
}