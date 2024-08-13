package entity

import (
	"database/sql"
	"time"
	cartModel "github.com/NetSinx/yconnect-shop/server/cart/model"
)

type User struct {
	Id              uint             `json:"id" gorm:"primaryKey"`
	Name            string           `json:"name"`
	Username        string           `json:"username" gorm:"unique"`
	Avatar          string           `json:"avatar"`
	Email           string           `json:"email" gorm:"unique" validate:"required,email"`
	Role            string           `json:"role" validate:"required"`
	Alamat          Alamat           `json:"alamat" validate:"required"`
	NoTelp          string           `json:"no_telp" gorm:"unique"`
	Password        string           `json:"password" validate:"required,min=5,containsany=!@#&*,containsany=0123456789,containsany=ABCDEFGHIJKLMNOPQRSTUVWXYZ"`
	EmailVerified   bool             `json:"email_verified" validate:"required"`
	EmailVerifiedAt sql.NullTime
	Cart            []cartModel.Cart `json:"cart"`
	CreatedAt       time.Time
	UpdatedAt       time.Time
}

type Alamat struct {
	Id          uint   `json:"id" gorm:"primaryKey"`
	AlamatRumah string `json:"alamat_rumah" validate:"required,min=5"`
	RT          int    `json:"rt" validate:"required"`
	RW          int    `json:"rw" validate:"required"`
	Kelurahan   string `json:"kelurahan" validate:"required"`
	Kecamatan   string `json:"kecamatan" validate:"required"`
	Kota        string `json:"kota" validate:"required"`
	KodePos     string `json:"kode_pos" validate:"required"`
	UserID      uint   `json:"user_id"`
}

type UserLogin struct {
	UsernameorEmail string `json:"UsernameorEmail"`
	Password        string `json:"password" validate:"required,min=5"`
}
