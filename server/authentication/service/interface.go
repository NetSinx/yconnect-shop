package service

import "github.com/NetSinx/yconnect-shop/server/authentication/domain"

type AuthService interface {
	LoginUser(userLogin domain.UserLogin) (string, string, error)
}