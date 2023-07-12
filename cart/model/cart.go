package model

import "time"

type Cart struct {
	Id         uint     `json:"id" gorm:"primaryKey"`
	Name       string   `json:"name" gorm:"unique"`
	Slug       string   `json:"slug" gorm:"unique"`
	Price      int      `json:"price"`
	Item       int      `json:"item"`
	UserId     uint     `json:"user_id" gorm:"unique"`
	CategoryId uint     `json:"category_id"`
	Category   Category `json:"category" gorm:"-"`
	User       User     `json:"user" gorm:"-"`
	CreatedAt  time.Time
	UpdatedAt  time.Time
}

type Category struct {
	Id   uint   `json:"id"`
	Name string `json:"name"`
	Slug string `json:"slug"`
}

type User struct {
	Id       uint   `json:"id"`
	Name     string `json:"name"`
	Username string `json:"username"`
}