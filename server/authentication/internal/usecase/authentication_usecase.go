package usecase

import (
	"context"
	"github.com/NetSinx/yconnect-shop/server/authentication/internal/model"
	"github.com/go-playground/validator/v10"
	"github.com/redis/go-redis/v9"
	"github.com/sirupsen/logrus"
	"golang.org/x/crypto/bcrypt"
)

type AuthUseCase struct {
	Log         *logrus.Logger
	Validator   *validator.Validate
	RedisClient *redis.Client
}

func NewAuthUseCase(log *logrus.Logger, validator *validator.Validate, redisClient *redis.Client) *AuthUseCase {
	return &AuthUseCase{
		Log:         log,
		Validator:   validator,
		RedisClient: redisClient,
	}
}

func (a *AuthUseCase) LoginUser(ctx *context.Context, loginRequest *model.LoginRequest) string {
	if err := validator.New().Struct(loginRequest); err != nil {
		a.Log.WithError(err).Error("error validating request body")
		return ""
	}

	hashedPassword, err := a.RedisClient.Get(*ctx, "email:"+loginRequest.Email).Bytes()
	if err != nil {
		a.Log.WithError(err).Error("error getting password")
		
	}

	if err := bcrypt.CompareHashAndPassword(hashedPassword, []byte(loginRequest.Password)); err != nil {
		a.Log.WithError(err).Error("invalid username or password")
		return ""
	}
}
