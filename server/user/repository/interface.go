package repository

import (
	"github.com/NetSinx/yconnect-shop/server/user/model/entity"
	"github.com/NetSinx/yconnect-shop/server/user/model/domain"
)

type UserRepo interface {
	RegisterUser(users entity.User) error
	ListUsers(users []entity.User) ([]entity.User, error)
	GetUser(user entity.User, username, email string) (entity.User, error)
	UpdateUser(user entity.User, username string) error
	VerifyOTP(verifyEmail domain.VerifyEmail) error
	VerifyEmail(verifyEmail domain.VerifyEmail) error
	DeleteUser(user entity.User, username string) error
}