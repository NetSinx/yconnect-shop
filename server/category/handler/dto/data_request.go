package dto

type CategoryRequest struct {
	Name    string    `json:"name" gorm:"unique" validate:"required,min=3"`
}