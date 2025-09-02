package usecase

import (
	"github.com/NetSinx/yconnect-shop/server/authentication/internal/model"
	"github.com/go-playground/validator/v10"
	"github.com/sirupsen/logrus"
)

type AuthUseCase struct {
	Log *logrus.Logger
	Validator *validator.Validate
}

func NewAuthUseCase(log *logrus.Logger, validator *validator.Validate) *AuthUseCase {
	return &AuthUseCase{
		Log: log,
		Validator: validator,
	}
}

func (a *AuthUseCase) LoginUser(loginRequest *model.LoginRequest) string {
	if err := a.Validator.Struct(loginRequest); err != nil {
		a.Log.WithError(err).Error("error validating request body")
		return ""
	}
}