package converter

import (
	"github.com/NetSinx/yconnect-shop/server/product/internal/entity"
	"github.com/NetSinx/yconnect-shop/server/product/internal/model"
)

func CategoryMirrorToResponse(categoryMirror *entity.CategoryMirror) *model.CategoryMirrorResponse {
	return &model.CategoryMirrorResponse{
		ID: categoryMirror.ID,
		Nama: categoryMirror.Nama,
		Slug: categoryMirror.Slug,
		CreatedAt: categoryMirror.CreatedAt,
		UpdatedAt: categoryMirror.UpdatedAt,
	}
}