package model

import (
	"github.com/NetSinx/yconnect-shop/server/product/internal/entity"
	"time"
)

type GetAllProductRequest struct {
	Page int `query:"page" validate:"min=1"`
	Size int `query:"size" validate:"min=1,max=100"`
}

type ParamRequest struct {
	Slug string `param:"slug"`
}

type ProductRequest struct {
	Nama         string          `form:"nama" validate:"required,max=255"`
	Gambar       []entity.Gambar `form:"gambar" validate:"required"`
	Deskripsi    string          `form:"deskripsi" validate:"required"`
	KategoriSlug string          `form:"kategori_slug" validate:"required"`
	Harga        int64           `form:"harga" validate:"required,min=0"`
	Stok         int             `form:"stok" validate:"required,min=0"`
}

type ProductResponse struct {
	ID           uint             `json:"id"`
	Nama         string           `json:"nama"`
	Slug         string           `json:"slug"`
	Gambar       []GambarResponse `json:"gambar"`
	Deskripsi    string           `json:"deskripsi"`
	KategoriSlug string           `json:"kategori_slug"`
	Harga        int64            `json:"harga"`
	Stok         int              `json:"stok"`
	CreatedAt    time.Time        `json:"created_at"`
	UpdatedAt    time.Time        `json:"updated_at"`
}

type GambarResponse struct {
	ID        uint      `json:"id"`
	Path      string    `json:"path"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type PageMetadataResponse struct {
	Page      int   `json:"page"`
	Size      int   `json:"size"`
	TotalItem int64 `json:"total_item"`
	TotalPage int64 `json:"total_page"`
}

type DataResponse[T any] struct {
	Data         T                     `json:"data"`
	PageMetadata *PageMetadataResponse `json:"paging,omitempty"`
}
