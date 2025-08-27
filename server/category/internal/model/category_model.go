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
	Size int `json:"size" validate:"min=1,max=100"`
}

type GetCategoryIDRequest struct {
	Slug string `json:"slug" validate:"required,max=50"`
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

type ListCategoryResponse struct {
	Data      *[]CategoryResponse `json:"data"`
	Page      int                 `json:"page"`
	Size      int                 `json:"size"`
	TotalItem int                 `json:"total_item"`
	TotalPage int                 `json:"total_page"`
}
