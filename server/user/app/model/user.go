package model

import (
	"time"
	cartModel "github.com/NetSinx/yconnect-shop/server/cart/model"
	"github.com/NetSinx/yconnect-shop/server/seller/model/entity"
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
	Seller    entity.Seller    `json:"seller" form:"seller"`
	Cart			[]cartModel.Cart `json:"cart" gorm:"-"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

type UserLogin struct {
	Username string  `json:"username"`
	Email    string  `json:"email"`
	Password string  `json:"password" validate:"required,min=5"`	
}