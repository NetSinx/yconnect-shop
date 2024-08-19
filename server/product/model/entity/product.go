package entity

import (
	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	Id          uint         `json:"id" gorm:"primaryKey"`
	Nama        string       `json:"nama" gorm:"unique" validate:"required,max=255"`
	Slug        string       `json:"slug" gorm:"unique" validate:"required"`
	Gambar      []Gambar     `json:"gambar" validate:"required"`
	Deskripsi   string       `json:"deskripsi" validate:"required"`
	KategoriId  uint         `json:"kategori_id" validate:"required"`
	Harga       int          `json:"harga" validate:"required"`
	Stok        int          `json:"stok" validate:"required"`
	Rating      float32      `json:"rating" validate:"required"`
	Kategori    Kategori     `json:"kategori"`
}

type Kategori struct {
	Id        uint        `json:"id"`
	Name      string      `json:"name"`
	Slug      string      `json:"slug"`
}

type Gambar struct {
	Id         uint   `json:"id"`
	Nama       string `json:"nama"`
	ProductID  uint   `json:"product_id"`
}