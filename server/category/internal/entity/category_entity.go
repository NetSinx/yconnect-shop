package entity

import "time"

type Category struct {
	ID        uint            `json:"id" gorm:"primaryKey"`
	Nama      string          `json:"nama"`
	Slug      string          `json:"slug" gorm:"uniqueIndex"`
	CreatedAt time.Time       `json:"created_at"`
	UpdatedAt time.Time       `json:"updated_at"`
}