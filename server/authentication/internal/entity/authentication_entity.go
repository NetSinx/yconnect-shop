package entity

import "time"

type UserAuthentication struct {
	ID          uint      `json:"id" gorm:"primaryKey"`
	NamaLengkap string    `json:"nama_lengkap" gorm:"type:varchar(100);not null"`
	Username    string    `json:"username" gorm:"uniqueIndex;type:varchar(50);not null"`
	Email       string    `json:"email" gorm:"uniqueIndex;type:varchar(100);not null"`
	Role        string    `json:"role" gorm:"type:varchar(20);not null"`
	NoHP        string    `json:"no_hp" gorm:"unique;type:varchar(12);not null"`
	Password string `json:"password" gorm:"type:varchar(255)"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}
