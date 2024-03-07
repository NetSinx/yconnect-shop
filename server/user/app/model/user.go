package model

import (
	"time"
	"github.com/NetSinx/yconnect-shop/server/product/app/model"
	cartModel "github.com/NetSinx/yconnect-shop/server/cart/model"
)

type User struct {
	Id        uint             `json:"id" gorm:"primaryKey"`
	Name      string           `json:"name" form:"name"`
	Username  string           `json:"username" form:"username" gorm:"unique"`
	Avatar    string           `json:"avatar" form:"avatar"`
	Email     string           `json:"email" form:"email" gorm:"unique" validate:"email"`
	Alamat    string           `json:"alamat" form:"alamat"`
	NoTelp    string           `json:"no_telp" form:"no_telp" gorm:"unique"`
	Password  string           `json:"password" form:"password" validate:"required,min=5,containsany=!@#&*,containsany=0123456789,containsany=ABCDEFGHIJKLMNOPQRSTUVWXYZ"`
	Token     string           `json:"token" form:"token"`
	Seller    Seller           `json:"seller" form:"seller"`
	Cart			[]cartModel.Cart `json:"cart" gorm:"-"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

type UserLogin struct {
	Username string  `json:"username"`
	Email    string  `json:"email"`
	Password string  `json:"password" validate:"required,min=5"`	
}

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