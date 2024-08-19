package entity

import (
	"time"
	prodEntity "github.com/NetSinx/yconnect-shop/server/product/model/entity"
)

type Order struct {
	Id        uint               `json:"id" gorm:"primaryKey"`
	ProductID int                `json:"product_id" validate:"required"`
	Product   prodEntity.Product `json:"product"`
	Username  string             `json:"username" validate:"required"`
	Kuantitas int                `json:"kuantitas" validate:"required"`
	Status    string             `json:"status" validate:"required"`
	Estimasi  time.Time          `json:"estimasi" validate:"required,datetime"`
	CreatedAt time.Time          `json:"created_at"`
	UpdatedAt time.Time          `json:"updated_at"`
}