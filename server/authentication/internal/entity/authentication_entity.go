package entity

type Authentication struct {
	ID       uint   `json:"id" gorm:"primaryKey"`
	Email    string `json:"email" gorm:"type:varchar(100);uniqueIndex"`
	Username string `json:"username" gorm:"type:varchar(100);unique"`
	Password string `json:"password" gorm:"type:varchar"`
}