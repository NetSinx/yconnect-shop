package usecase

import (
	"github.com/NetSinx/yconnect-shop/server/authentication/internal/entity"
	"github.com/NetSinx/yconnect-shop/server/authentication/internal/model"
	"github.com/labstack/echo/v4"
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
		a.Log.WithError(err).Error("error creating user data mirror")
		return echo.ErrInternalServerError
	}

	if err := tx.Commit().Error; err != nil {
		a.Log.WithError(err).Error("error creating user data mirror")
		return echo.ErrInternalServerError
	}

	return nil
}

func (a *AuthUseCase) Update(user *model.UserEvent) error {
	tx := a.DB.Begin()
	defer tx.Rollback()

	entity := new(entity.Authentication)
	if err := a.AuthRepository.GetByID(tx, entity, user.ID); err != nil {
		a.Log.WithError(err).Error("error getting user data mirror")
		return echo.ErrNotFound
	}

	entity.ID = user.ID
	entity.Email = user.Email
	entity.Role = user.Role
	entity.Password = user.Password

	if err := a.AuthRepository.Update(tx, entity); err != nil {
		a.Log.WithError(err).Error("error updating user data mirror")
		return echo.ErrInternalServerError
	}

	if err := tx.Commit().Error; err != nil {
		a.Log.WithError(err).Error("error updating user data mirror")
		return echo.ErrInternalServerError
	}

	return nil
}

func (a *AuthUseCase) Delete(user *model.UserEvent) error {
	tx := a.DB.Begin()
	defer tx.Rollback()

	entity := new(entity.Authentication)
	if err := a.AuthRepository.GetByID(tx, entity, user.ID); err != nil {
		a.Log.WithError(err).Error("error getting user data mirror")
		return echo.ErrNotFound
	}

	entity.ID = user.ID
	entity.Email = user.Email
	entity.Role = user.Role
	entity.Password = user.Password

	if err := a.AuthRepository.Delete(tx, entity); err != nil {
		a.Log.WithError(err).Error("error deleting user data mirror")
		return echo.ErrInternalServerError
	}

	if err := tx.Commit().Error; err != nil {
		a.Log.WithError(err).Error("error deleting user data mirror")
		return echo.ErrInternalServerError
	}

	return nil
}