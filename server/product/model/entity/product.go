package entity

import "time"

type Product struct {
	Id          uint         `json:"id" gorm:"primaryKey"`
	Nama        string       `json:"nama" gorm:"unique" validate:"required,max=255"`
	Slug        string       `json:"slug" gorm:"unique" validate:"required"`
	Images      []Images     `json:"gambar" validate:"required"`
	Deskripsi   string       `json:"deskripsi" validate:"required"`
	KategoriId  uint         `json:"kategori_id" validate:"required"`
	Harga       int          `json:"harga" validate:"required"`
	Stok        int          `json:"stok" validate:"required"`
	Rating      float32      `json:"rating" validate:"required"`
	Kategori    Kategori     `json:"kategori" gorm:"-"`
	CreatedAt   time.Time    `json:"created_at"`
	UpdatedAt   time.Time    `json:"updated_at"`
}

type Kategori struct {
	Id        uint        `json:"id"`
	Name      string      `json:"name"`
	Slug      string      `json:"slug"`
}

type Images struct {
	Id         uint   `json:"id"`
	Path       string `json:"path"`
	ProductID  uint   `json:"product_id"`
}