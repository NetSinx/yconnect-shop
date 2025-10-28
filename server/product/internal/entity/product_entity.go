package entity

import "time"

type Gambar struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	Path      string    `json:"path" gorm:"type:varchar(255);not null"`
	ProductID uint      `json:"product_id" gorm:"not null"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type Product struct {
	ID           uint      `json:"id" gorm:"primaryKey"`
	Nama         string    `json:"nama" gorm:"type:varchar(255);not null"`
	Slug         string    `json:"slug" gorm:"type:varchar(255);uniqueIndex;not null"`
	Gambar       []Gambar  `json:"gambar" gorm:"foreignKey:ProductID"`
	Deskripsi    string    `json:"deskripsi" gorm:"not null"`
	KategoriSlug string    `json:"kategori_slug" gorm:"not null"`
	Harga        int64     `json:"harga" gorm:"not null"`
	Stok         int       `json:"stok" gorm:"not null"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}

type CategoryMirror struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	Nama      string    `json:"nama" gorm:"not null"`
	Slug      string    `json:"slug" gorm:"uniqueIndex;not null"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type Tabler interface {
	TableName() string
}

func (Gambar) TableName() string {
	return "gambar"
}
