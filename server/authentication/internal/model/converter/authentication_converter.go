package converter

import (
	"github.com/NetSinx/yconnect-shop/server/authentication/internal/entity"
	"github.com/NetSinx/yconnect-shop/server/authentication/internal/model"
)

func AuthenticationToResponse(entity *entity.Authentication) *model.AuthenticationResponse {
	return &model.AuthenticationResponse{
		ID: entity.ID,
		Email: entity.Email,
		Role: entity.Role,
		Password: entity.Password,
	}
}