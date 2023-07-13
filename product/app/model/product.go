package model

import "time"

type User struct {
	Id        uint        `json:"id" gorm:"primaryKey"`
	Name      string      `json:"name"`
}

type Product struct {
	Id          uint       `json:"id" form:"id" gorm:"primaryKey"`
	Name        string     `json:"name" form:"name" gorm:"unique" validate:"required,max=255"`
	Slug        string     `json:"slug" form:"slug" gorm:"unique" validate:"required"`
	Description string     `json:"description" form:"description" gorm:"type:text" validate:"required"`
	CategoryId  uint       `json:"category_id" form:"category_id"`
	SellerId    uint       `json:"seller_id" form:"seller_id"`
	Price				int        `json:"price" form:"price" validate:"required"`
	Stock				int        `json:"stock" form:"stock" validate:"required"`
	Category    Category   `json:"category" gorm:"-"`
	User        User       `json:"user" gorm:"-"`
	CreatedAt		time.Time
	UpdatedAt		time.Time
}

type Category struct {
	Id        uint        `json:"id" gorm:"primaryKey"`
	Name      string      `json:"name" gorm:"unique" validate:"required,min=3"`
	Slug      string      `json:"slug" gorm:"unique" validate:"required,min=3,lowercase"`
}