package service

import (
	"github.com/NetSinx/yconnect-shop/server/user/model/domain"
	"github.com/NetSinx/yconnect-shop/server/user/model/entity"
)

type UserServ interface {
	RegisterUser(users entity.User) error
	ListUsers(users []entity.User) ([]entity.User, error)
	GetUser(user entity.User, username, email string) (entity.User, error)
	UpdateUser(user entity.User, username string) error
	DeleteUser(user entity.User, username, email string) error
	VerifyOTP(verifyEmail domain.VerifyEmail) (string, error)
	VerifyEmail(verifyEmail domain.VerifyEmail) error
}