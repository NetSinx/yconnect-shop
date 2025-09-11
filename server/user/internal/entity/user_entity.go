package entity

import "time"

type User struct {
	ID          uint      `json:"id" gorm:"primaryKey"`
	NamaLengkap string    `json:"nama_lengkap" gorm:"type:varchar(100);not null"`
	Username    string    `json:"username" gorm:"uniqueIndex;type:varchar(50);not null"`
	Email       string    `json:"email" gorm:"unique;type:varchar(100);not null"`
	Role        string    `json:"role" gorm:"type:varchar(20);not null"`
	Alamat      *Alamat   `json:"alamat"`
	NoHP        string    `json:"no_hp" gorm:"unique;type:varchar(16);not null"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type Alamat struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	NamaJalan string    `json:"nama_jalan" gorm:"type:varchar(100);not null"`
	RT        int       `json:"rt" gorm:"not null"`
	RW        int       `json:"rw" gorm:"not null"`
	Kelurahan string    `json:"kelurahan" gorm:"type:varchar(100);not null"`
	Kecamatan string    `json:"kecamatan" gorm:"type:varchar(100);not null"`
	Kota      string    `json:"kota" gorm:"type:varchar(100);not null"`
	KodePos   int       `json:"kode_pos" gorm:"not null"`
	UserID    uint      `json:"user_id" gorm:"not null"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type Tabler interface {
	TableName() string
}

func (Alamat) TableName() string {
	return "alamat"
}
