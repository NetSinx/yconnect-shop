package model

import (
	"time"
	"github.com/NetSinx/yconnect-shop/server/product/app/model"
)

type Category struct {
	Id        uint            `json:"id" gorm:"primaryKey"`
	Name      string          `json:"name" gorm:"unique" validate:"required,min=3"`
	Slug      string          `json:"slug" gorm:"unique" validate:"required,min=3,lowercase"`
	Product   []model.Product `json:"product" gorm:"-"`
	CreatedAt time.Time
	UpdatedAt time.Time
}