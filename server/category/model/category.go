package model

import "time"

type Category struct {
	Id        uint            `json:"id" gorm:"primaryKey"`
	Name      string          `json:"name" gorm:"unique" validate:"required,min=3"`
	Slug      string          `json:"slug" gorm:"unique" validate:"required,min=3,lowercase"`
	CreatedAt time.Time
	UpdatedAt time.Time
}