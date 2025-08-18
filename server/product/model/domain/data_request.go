package domain

import "github.com/NetSinx/yconnect-shop/server/product/model/entity"

type ProductRequest struct {
	Nama       string       `json:"nama" validate:"required,max=255"`
	Gambar     []entity.Gambar `json:"gambar" gorm:"foreignKey:ProductID" validate:"required"`
	Deskripsi  string       `json:"deskripsi" validate:"required"`
	KategoriID uint         `json:"kategori_id" validate:"required"`
	Harga      int          `json:"harga" validate:"required"`
	Stok       int          `json:"stok" validate:"required"`
}