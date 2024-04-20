package service

import "github.com/NetSinx/yconnect-shop/server/user/model/entity"

type UserServ interface {
	RegisterUser(users entity.User) error
	LoginUser(userLogin entity.UserLogin) (string, error)
	ListUsers(users []entity.User) ([]entity.User, error)
	GetUser(users entity.User, username string) (entity.User, error)
	UpdateUser(users entity.User, username string) error
	DeleteUser(users entity.User, username string) error
	VerifyEmail(verifyEmail entity.VerifyEmail) (string, error)
}