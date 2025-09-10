package converter

import (
	"github.com/NetSinx/yconnect-shop/server/authentication/internal/entity"
	"github.com/NetSinx/yconnect-shop/server/authentication/internal/model"
)

func UserRegisterToResponse(entity *entity.UserAuthentication) *model.RegisterResponse {
	return &model.RegisterResponse{
		ID:          entity.ID,
		NamaLengkap: entity.NamaLengkap,
		Username:    entity.Username,
		Email:       entity.Email,
		NoHP:        entity.NoHP,
	}
}

func UserAuthenticationToResponse(entity *entity.UserAuthentication) *model.AuthenticationResponse {
	return &model.AuthenticationResponse{
		ID:    entity.ID,
		Email: entity.Email,
		Role:  entity.Role,
	}
}
