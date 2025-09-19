package model

type CategoryEvent struct {
	Id   uint   `json:"id" gorm:"primaryKey"`
	Name string `json:"name" gorm:"unique" validate:"required,min=3"`
	Slug string `json:"slug" validate:"required,min=3"`
}
