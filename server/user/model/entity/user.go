package entity

import (
	"database/sql"
	"time"
)

type User struct {
	Id              uint                `json:"id" gorm:"primaryKey"`
	Name            string              `json:"name"`
	Username        string              `json:"username" gorm:"unique" validate:"required"`
	Avatar          string              `json:"avatar"`
	Email           string              `json:"email" gorm:"unique" validate:"required,email"`
	Role            string              `json:"role" validate:"required"`
	Alamat          Alamat              `json:"alamat"`
	NoTelp          string              `json:"no_telp" gorm:"unique" validate:"required"`
	Password        string              `json:"password" validate:"required,min=5,containsany=!@#&*,containsany=0123456789,containsany=ABCDEFGHIJKLMNOPQRSTUVWXYZ"`
	EmailVerified   bool                `json:"email_verified"`
	EmailVerifiedAt sql.NullTime        `json:"email_verified_at"`
	CreatedAt       time.Time           `json:"created_at"`
	UpdatedAt       time.Time           `json:"updated_at"`
}

type Alamat struct {
	Id          uint       `json:"id" gorm:"primaryKey"`
	AlamatRumah string     `json:"alamat_rumah" validate:"required,min=5"`
	RT          int        `json:"rt" validate:"required"`
	RW          int        `json:"rw" validate:"required"`
	Kelurahan   string     `json:"kelurahan" validate:"required"`
	Kecamatan   string     `json:"kecamatan" validate:"required"`
	Kota        string     `json:"kota" validate:"required"`
	KodePos     int        `json:"kode_pos" validate:"required"`
	UserID      uint       `json:"user_id"`
	CreatedAt   time.Time  `json:"created_at"`
	UpdatedAt   time.Time  `json:"updated_at"`
}
