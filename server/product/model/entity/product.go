package entity

import "time"

type Gambar struct {
	Id           uint      `json:"id" gorm:"primaryKey"`
	Path         string    `json:"path" validate:"required"`
	ProductID    uint      `json:"-"`
}

type Product struct {
	Id          uint         `json:"id" gorm:"primaryKey"`
	Nama        string       `json:"nama" validate:"required,max=255"`
	Slug        string       `json:"slug" gorm:"unique" validate:"required"`
	Gambar      []Gambar     `json:"gambar" gorm:"foreignKey:ProductID" validate:"required"`
	Deskripsi   string       `json:"deskripsi" validate:"required"`
	KategoriID  uint         `json:"kategori_id" validate:"required"`
	Harga       int          `json:"harga" validate:"required"`
	Stok        int          `json:"stok" validate:"required"`
	CreatedAt   time.Time    `json:"created_at"`
	UpdatedAt   time.Time    `json:"updated_at"`
}

type Tabler interface {
  TableName() string
}

func (Gambar) TableName() string {
  return "gambar"
}