package converter

import (
	"github.com/NetSinx/yconnect-shop/server/user/internal/entity"
	"github.com/NetSinx/yconnect-shop/server/user/internal/model"
)

func UserToResponse(userEntity *entity.User) *model.UserResponse {
	alamatResponse := &model.AlamatResponse{
		ID: userEntity.Alamat.ID,
		NamaJalan: userEntity.Alamat.NamaJalan,
		RT: userEntity.Alamat.RT,
		RW: userEntity.Alamat.RW,
		Kelurahan: userEntity.Alamat.Kelurahan,
		Kecamatan: userEntity.Alamat.Kecamatan,
		Kota: userEntity.Alamat.Kota,
		KodePos: userEntity.Alamat.KodePos,
	}

	return &model.UserResponse{
		ID: userEntity.ID,
		NamaLengkap: userEntity.NamaLengkap,
		Username: userEntity.Username,
		Email: userEntity.Email,
		Alamat: alamatResponse,
		NoHP: userEntity.NoHP,
	}
}

func UserToDeleteUserEvent(userEntity *entity.User) *model.DeleteUserEvent {
	return &model.DeleteUserEvent{
		Email: userEntity.Email,
	}
}