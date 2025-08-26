package model

import "time"

type CreateCategoryRequest struct {
	Nama string `json:"nama" validate:"required,max=50"`
}

type UpdateCategoryRequest struct {
	Nama string `json:"nama" validate:"required,max=50"`
	Slug string `json:"slug" validate:"required"`
}

type DeleteCategoryRequest struct {
	Slug string `json:"slug" validate:"required,max=50"`
}

type ListCategoryRequest struct {
	Page int `json:"page" validate:"min=1"`
	Size int `json:"size" validate:"min=1,max=20"`
}

type GetCategoryByIdRequest struct {
	ID int `json:"id" validate:"required"`
}

type GetCategoryBySlugRequest struct {
	Slug string `json:"slug" validate:"required,max=50"`
}

type CategoryResponse struct {
	ID        uint      `json:"id"`
	Nama      string    `json:"nama"`
	Slug      string    `json:"slug"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}