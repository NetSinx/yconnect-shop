package entity

import (
	"time"
	"github.com/NetSinx/yconnect-shop/server/product/model/entity"
)

type Cart struct {
	Id         uint           `json:"id" gorm:"primaryKey"`
	ProductID  int            `json:"product_id" validate:"required"`
	Product    entity.Product  `json:"product" gorm:"-"`
	Item       int            `json:"item" validate:"required"`
	UserID     int            `json:"user_id" `
	CreatedAt  time.Time
	UpdatedAt  time.Time
}
