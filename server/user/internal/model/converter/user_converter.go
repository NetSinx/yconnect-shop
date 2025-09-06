package converter

import (
	"github.com/NetSinx/yconnect-shop/server/user/internal/entity"
	"github.com/NetSinx/yconnect-shop/server/user/internal/model"
)

func UserToResponse(userEntity *entity.User) *model.UserResponse {
	return &model.UserResponse{
		ID: userEntity.ID,
		NamaLengkap: userEntity.NamaLengkap,
		Username: userEntity.Username,
		Email: userEntity.Email,
		NoHP: userEntity.NoHP,
	}
}