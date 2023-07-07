package service

import "github.com/NetSinx/yconnect-shop/user/app/model"

type UserServ interface {
	RegisterUser(users model.User) error
	LoginUser(email string) (model.User, error)
	ListUsers(users []model.User) ([]model.User, error)
	FindUser(users model.User, id uint) (model.User, error)
	UpdateUser(users model.User, id uint) error
	DeleteUser(users model.User, id uint) error
}