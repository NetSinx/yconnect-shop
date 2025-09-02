package entity

type Authentication struct {
	ID       uint   `json:"id" gorm:"primaryKey"`
	Email    string `json:"email" gorm:"uniqueIndex;type:varchar(100)"`
	Role     string `json:"role" gorm:"type:varchar(10)"`
	Password string `json:"password" gorm:"type:varchar(255)"`
}
