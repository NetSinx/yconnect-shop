package converter

import (
	"github.com/NetSinx/yconnect-shop/server/category/internal/entity"
	"github.com/NetSinx/yconnect-shop/server/category/internal/model"
)

func CategoryToEvent(category *entity.Category) *model.CategoryEvent {
	return &model.CategoryEvent{
		ID: category.ID,
		Nama: category.Nama,
		Slug: category.Slug,
	}
}

func CategoryToResponse(category *entity.Category) *model.CategoryResponse {
	return &model.CategoryResponse{
		ID: category.ID,
		Nama: category.Nama,
		Slug: category.Slug,
		CreatedAt: category.CreatedAt,
		UpdatedAt: category.UpdatedAt,
	}
}