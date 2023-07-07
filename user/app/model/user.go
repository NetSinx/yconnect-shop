package model

import (
	"time"
	"github.com/NetSinx/yconnect-shop/product/app/model"
)

type User struct {
	Id        uint            `json:"id" gorm:"primaryKey"`
	Name      string          `json:"name"`
	Username  string          `json:"username" gorm:"unique"`
	Email     string          `json:"email" gorm:"unique" validate:"email"`
	Alamat    string          `json:"alamat"`
	NoTelp    string          `json:"no_telp" gorm:"unique"`
	Password  string          `json:"password"`
	Product   []model.Product `json:"product" gorm:"-"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

type UserLogin struct {
	Email    string  `json:"email" validate:"required,email"`
	Password string  `json:"password" validate:"required"`	
}