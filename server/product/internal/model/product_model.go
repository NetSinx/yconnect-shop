package model

import "github.com/NetSinx/yconnect-shop/server/product/internal/entity"

type ProductRequest struct {
	Nama         string          `json:"nama" validate:"required,max=255"`
	Gambar       []entity.Gambar `json:"gambar" gorm:"foreignKey:ProductID" validate:"required"`
	Deskripsi    string          `json:"deskripsi" validate:"required"`
	KategoriSlug string          `json:"kategori_slug" validate:"required"`
	Harga        int             `json:"harga" validate:"required"`
	Stok         int             `json:"stok" validate:"required"`
}

type RespData struct {
	Data any `json:"data"`
}

type MessageResp struct {
	Message string `json:"message"`
}

type ResponseCSRF struct {
	CSRFToken string `json:"csrf_token"`
}

type CategoryEvent struct {
	Id   uint   `json:"id" gorm:"primaryKey"`
	Name string `json:"name" gorm:"unique" validate:"required,min=3"`
	Slug string `json:"slug" validate:"required,min=3"`
}
