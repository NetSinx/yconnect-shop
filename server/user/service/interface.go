package service

import "github.com/NetSinx/yconnect-shop/server/user/model"

type UserServ interface {
	RegisterUser(users model.User) error
	LoginUser(userLogin model.UserLogin) (string, error)
	ListUsers(users []model.User) ([]model.User, error)
	GetUser(users model.User, username string) (model.User, error)
	UpdateUser(users model.User, username string) error
	DeleteUser(users model.User, username string) error
}