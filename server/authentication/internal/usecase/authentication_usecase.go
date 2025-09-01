package usecase

import (
	"context"
	"time"
	"github.com/NetSinx/yconnect-shop/server/authentication/internal/entity"
	"github.com/NetSinx/yconnect-shop/server/authentication/internal/model"
	"github.com/NetSinx/yconnect-shop/server/authentication/internal/repository"
	"github.com/go-playground/validator/v10"
	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
	"github.com/redis/go-redis/v9"
	"github.com/sirupsen/logrus"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type AuthUseCase struct {
	DB             *gorm.DB
	Log            *logrus.Logger
	Validator      *validator.Validate
	RedisClient    *redis.Client
	AuthRepository *repository.AuthRepository
}

func NewAuthUseCase(db *gorm.DB, log *logrus.Logger, validator *validator.Validate, redisClient *redis.Client, authRepository *repository.AuthRepository) *AuthUseCase {
	return &AuthUseCase{
		DB:             db,
		Log:            log,
		Validator:      validator,
		RedisClient:    redisClient,
		AuthRepository: authRepository,
	}
}

func (a *AuthUseCase) LoginUser(ctx context.Context, loginRequest *model.LoginRequest) (*model.AuthTokenResponse, error) {
	tx := a.DB.WithContext(ctx).Begin()
	defer tx.Rollback()

	if err := validator.New().Struct(loginRequest); err != nil {
		a.Log.WithError(err).Error("error validating request body")
		return nil, echo.ErrBadRequest
	}

	resultAuth := new(entity.Authentication)
	if err := a.RedisClient.HGetAll(ctx, "email:"+loginRequest.Email).Scan(resultAuth); err != nil {
		authEntity := new(entity.Authentication)
		result, err := a.AuthRepository.GetByEmail(tx, authEntity, loginRequest.Email)
		if err != nil {
			a.Log.WithError(err).Error("error getting user")
			return nil, echo.ErrForbidden
		}

		if err := bcrypt.CompareHashAndPassword([]byte(resultAuth.Password), []byte(loginRequest.Password)); err != nil {
			a.Log.WithError(err).Error("invalid password")
			return nil, echo.ErrForbidden
		}

		response := &model.AuthTokenResponse{
			AuthToken: jwt,
		}

		a.RedisClient.HSet(ctx, "email:"+loginRequest.Email, result, time.Minute)
		
		return response, nil
	}

	if err := bcrypt.CompareHashAndPassword([]byte(resultAuth.Password), []byte(loginRequest.Password)); err != nil {
		a.Log.WithError(err).Error("invalid password")
		return nil, echo.ErrForbidden
	}

	signingKey := []byte("test123")

	type CustomClaims struct {
		Username string `json:"username"`
		Email    string `json:"email"`
		jwt.RegisteredClaims
	}

	claims := CustomClaims{
		resultAuth.Username,
		resultAuth.Email,
		jwt.RegisteredClaims{
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	jwt, err := token.SignedString(signingKey)
	if err != nil {
		a.Log.WithError(err).Error("failed to generate token")
		return nil, echo.ErrInternalServerError
	}

	response := &model.AuthTokenResponse{
		AuthToken: jwt,
	}

	return response, nil
}
