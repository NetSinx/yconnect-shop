package usecase

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/NetSinx/yconnect-shop/server/authentication/internal/model"
	"github.com/NetSinx/yconnect-shop/server/user/model/entity"
	"github.com/go-playground/validator/v10"
	"github.com/golang-jwt/jwt/v5"
	"github.com/sirupsen/logrus"
	"golang.org/x/crypto/bcrypt"
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
	if err := validator.New().Struct(loginRequest); err != nil {
		a.Log.WithError(err).Error("error validating request body")
		return ""
	}

	
}