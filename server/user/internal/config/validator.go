package config

import (
	"github.com/NetSinx/yconnect-shop/server/user/internal/helpers"
	"github.com/go-playground/validator/v10"
)

func NewValidator() *validator.Validate {
	return validator.New()
}

func init() {
	validate := NewValidator()
	validate.RegisterValidation("passwd", helpers.PasswordValidation)
}