package entity

type Authentication struct {
	ID       uint   `json:"id" gorm:"primaryKey"`
	Email    string `json:"email" gorm:"type:varchar(100);uniqueIndex"`
	Role     string `json:"role" gorm:"type:varchar(10)"`
	Password string `json:"password" gorm:"type:varchar"`
}
