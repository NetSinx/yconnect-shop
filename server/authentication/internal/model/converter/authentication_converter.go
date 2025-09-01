package converter

import (
	"github.com/NetSinx/yconnect-shop/server/authentication/internal/entity"
	"github.com/NetSinx/yconnect-shop/server/authentication/internal/model"
)

func AuthenticationToResponse(authEntity *entity.Authentication) *model.LoginResponse {
	return &model.LoginResponse{
		Email: authEntity.Email,
		Password: authEntity.Password,
	}
}