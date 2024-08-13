package entity

import (
	"time"
	prodEntity "github.com/NetSinx/yconnect-shop/server/product/model/entity"
	"gorm.io/gorm"
)

type Order struct {
	gorm.Model
	ProductID int                `json:"product_id" validate:"required"`
	Product   prodEntity.Product `json:"product"`
	UserID    int                `json:"user_id" validate:"required"`
	Kuantitas int                `json:"kuantitas" validate:"required"`
	Status    string             `json:"status" validate:"required"`
	Estimasi  time.Time          `json:"estimasi" validate:"required"`
}