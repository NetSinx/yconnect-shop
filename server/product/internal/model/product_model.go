package model

import (
	"github.com/NetSinx/yconnect-shop/server/product/internal/entity"
	"time"
)

type GetAllProductRequest struct {
	Page int `json:"page" validate:"min=1"`
	Size int `json:"size" validate:"min=1,max=100"`
}

type ProductRequest struct {
	Nama         string          `json:"nama" validate:"required,max=255"`
	Gambar       []entity.Gambar `json:"gambar" validate:"required"`
	Deskripsi    string          `json:"deskripsi" validate:"required"`
	KategoriSlug string          `json:"kategori_slug" validate:"required"`
	Harga        int64           `json:"harga" validate:"required,min=0"`
	Stok         int             `json:"stok" validate:"required,min=0"`
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
