package entity

import (
	"time"
	"github.com/NetSinx/yconnect-shop/server/product/model/entity"
)

type Order struct {
	Id        uint            `json:"id" gorm:"primaryKey"`
	ProductID int             `json:"product_id"`
	Product   entity.Product  `json:"product" gorm:"-"`
	UserID    int             `json:"user_id"`
	Kuantitas int             `json:"kuantitas"`
	Status    string          `json:"status"`
	Estimasi  time.Time       `json:"estimasi"`
	CreatedAt time.Time       `json:"created_at"`
	UpdatedAt time.Time       `json:"updated_at"`
}