package entity

import (
	"time"
	"github.com/NetSinx/yconnect-shop/server/product/app/model"
)

type Seller struct {
	Id        uint             `json:"id" gorm:"primaryKey"`
	Name      string           `json:"name" form:"name"`
	Username  string           `json:"username" form:"username" gorm:"unique"`
	Avatar    string           `json:"avatar" form:"avatar"`
	Email     string           `json:"email" form:"email" gorm:"unique"`
	Alamat    string           `json:"alamat" form:"alamat"`
	NoTelp    string           `json:"no_telp" form:"no_telp" gorm:"unique"`
	Product   []model.Product  `json:"product" gorm:"-"`
	UserID    uint             `json:"user_id"`
	CreatedAt time.Time
	UpdatedAt time.Time
}