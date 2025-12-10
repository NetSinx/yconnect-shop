package config

import (
	"github.com/NetSinx/yconnect-shop/server/authentication/internal/helpers"
	"github.com/go-playground/validator/v10"
)

func NewValidator() *validator.Validate {
	validate := validator.New()
	validate.RegisterValidation("passwd", helpers.PasswordValidation)
	return validate
}
