package entity

import (
	"database/sql"
	"time"
	cartModel "github.com/NetSinx/yconnect-shop/server/cart/model"
	"github.com/NetSinx/yconnect-shop/server/seller/model/entity"
)

type User struct {
	Id        				uint             `json:"id" gorm:"primaryKey"`
	Name      				string           `json:"name"`
	Username  				string           `json:"username" gorm:"unique"`
	Avatar    				string           `json:"avatar"`
	Email     				string           `json:"email" gorm:"unique" validate:"email"`
	Alamat    				string           `json:"alamat"`
	NoTelp    				string           `json:"no_telp" gorm:"unique"`
	Password  				string           `json:"password" validate:"required,min=5,containsany=!@#&*,containsany=0123456789,containsany=ABCDEFGHIJKLMNOPQRSTUVWXYZ"`
	Token     				string           `json:"token"`
	Seller    				entity.Seller    `json:"seller"`
	Cart							[]cartModel.Cart `json:"cart" gorm:"-"`
	EmailVerified			bool						 `json:"email_verified"`
	EmailVerifiedAt		sql.NullTime
	CreatedAt 				time.Time
	UpdatedAt 				time.Time
}

type UserLogin struct {
	Username string  `json:"username"`
	Email    string  `json:"email"`
	Password string  `json:"password" validate:"required,min=5"`	
}