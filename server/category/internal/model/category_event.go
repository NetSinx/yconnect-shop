package model

type CategoryEvent struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	Nama      string    `json:"nama"`
	Slug      string    `json:"slug" gorm:"uniqueIndex"`
}