package entity

import "time"

type Order struct {
	Id        uint      `json:"id" gorm:"primaryKey"`
	ProductID int       `json:"product_id"`
	UserID    int       `json:"user_id"`
	Kuantitas int       `json:"kuantitas" validate:"required"`
	Status    string    `json:"status" validate:"required"`
	Estimasi  time.Time `json:"estimasi" validate:"required"`
	CreatedAt time.Time
	UpdatedAt time.Time
}