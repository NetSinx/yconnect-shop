package dto

type CategoryRequest struct {
	Name      string          `json:"name" gorm:"unique" validate:"required,min=3"`
	Slug      string          `json:"slug" gorm:"unique" validate:"required,min=3,lowercase"`
}