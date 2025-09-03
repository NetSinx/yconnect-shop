package usecase

import (
	"github.com/NetSinx/yconnect-shop/server/authentication/internal/entity"
	"github.com/NetSinx/yconnect-shop/server/authentication/internal/model"
)

func (a *AuthUseCase) Create(user *model.UserEvent) error {
	tx := a.DB.Begin()
	defer tx.Rollback()

	entity := &entity.Authentication{
		ID: user.ID,
		Email: user.Email,
		Role: user.Role,
		Password: user.Password,
	}

	if err := a.AuthRepository.Create(tx, entity); err != nil {
		return err
	}

	return nil
}