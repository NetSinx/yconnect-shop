package usecase

import (
	"context"

	"github.com/NetSinx/yconnect-shop/server/authentication/internal/entity"
	"github.com/NetSinx/yconnect-shop/server/authentication/internal/model"
)

func (a *AuthUseCase) Create(ctx context.Context, user *model.UserEvent) error {
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