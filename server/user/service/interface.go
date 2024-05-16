package service

import (
	"github.com/NetSinx/yconnect-shop/server/user/model/entity"
	"github.com/NetSinx/yconnect-shop/server/user/model/domain"
)

type UserServ interface {
	RegisterUser(users entity.User) error
	LoginUser(userLogin entity.UserLogin) (string, error)
	ListUsers(users []entity.User) ([]entity.User, error)
	GetUser(users entity.User, username string) (entity.User, error)
	UpdateUser(users entity.User, username string) error
	DeleteUser(users entity.User, username string) error
	SendOTP(verifyEmail domain.VerifyEmail) (string, error)
	VerifyEmail(verifyEmail domain.VerifyEmail) error
}