package service

import "github.com/NetSinx/yconnect-shop/server/authentication/model/domain"

type AuthService interface {
	LoginUser(userLogin domain.UserLogin) (string, string, string, error)
}