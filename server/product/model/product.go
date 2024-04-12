package model

import "time"

type Seller struct {
	Id        uint        `json:"id"`
	Name      string      `json:"name"`
}

type Product struct {
	Id          uint         `json:"id" gorm:"primaryKey"`
	Name        string       `json:"name" gorm:"unique" validate:"required,max=255"`
	Slug        string       `json:"slug" gorm:"unique" validate:"required"`
	Image       []Image      `json:"images" validate:"required"`
	Description string       `json:"description" gorm:"type:text" validate:"required"`
	CategoryId  uint         `json:"category_id" validate:"required"`
	SellerId    uint         `json:"seller_id" validate:"required"`
	Price       int          `json:"price" validate:"required"`
	Stock       int          `json:"stock" validate:"required"`
	Category    Category     `json:"category" gorm:"-"`
	Seller      Seller       `json:"seller" gorm:"-"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

type Category struct {
	Id        uint        `json:"id"`
	Name      string      `json:"name"`
	Slug      string      `json:"slug"`
}

type Image struct {
	Id         uint   `json:"id"`
	Name       string `json:"name"`
	ProductID  uint   `json:"product_id"`
}