package usecase

import (
	"context"
	"crypto/rand"
	"encoding/base64"
	"time"
	"github.com/NetSinx/yconnect-shop/server/authentication/internal/entity"
	"github.com/NetSinx/yconnect-shop/server/authentication/internal/helpers"
	"github.com/NetSinx/yconnect-shop/server/authentication/internal/model"
	"github.com/NetSinx/yconnect-shop/server/authentication/internal/repository"
	"github.com/go-playground/validator/v10"
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
	TokenUtil      *helpers.TokenUtil
}

func NewAuthUseCase(db *gorm.DB, log *logrus.Logger, validator *validator.Validate, redisClient *redis.Client, authRepository *repository.AuthRepository, tokenUtil *helpers.TokenUtil) *AuthUseCase {
	return &AuthUseCase{
		DB:             db,
		Log:            log,
		Validator:      validator,
		RedisClient:    redisClient,
		AuthRepository: authRepository,
		TokenUtil:      tokenUtil,
	}
}

func (a *AuthUseCase) LoginUser(ctx context.Context, loginRequest *model.LoginRequest) (*model.AuthTokenResponse, error) {
	tx := a.DB.WithContext(ctx).Begin()
	defer tx.Rollback()

	if err := validator.New().Struct(loginRequest); err != nil {
		a.Log.WithError(err).Error("error validating request body")
		return nil, echo.ErrBadRequest
	}

	authEntity := new(entity.Authentication)
	result, err := a.AuthRepository.GetByEmail(tx, authEntity, loginRequest.Email)
	if err != nil {
		a.Log.WithError(err).Error("error getting user")
		return nil, echo.ErrUnauthorized
	}

	if err := tx.Commit().Error; err != nil {
		a.Log.WithError(err).Error("error commit the changes in a transaction")
		return nil, echo.ErrInternalServerError
	}

	if err := bcrypt.CompareHashAndPassword([]byte(result.Password), []byte(loginRequest.Password)); err != nil {
		a.Log.WithError(err).Error("error to compare hash and password")
		return nil, echo.ErrUnauthorized
	}

	jwtToken, err := a.TokenUtil.CreateToken(ctx, result.Role, result.ID)
	if err != nil {
		a.Log.WithError(err).Error("error generating jwt token")
		return nil, echo.ErrInternalServerError
	}

	response := &model.AuthTokenResponse{
		AuthToken: jwtToken,
	}

	return response, nil
}

func (a *AuthUseCase) Verify(ctx context.Context, authTokenRequest *model.AuthTokenRequest) (map[string]string, error) {
	if err := a.Validator.Struct(authTokenRequest); err != nil {
		a.Log.WithError(err).Error("error validating request")
		return nil, echo.ErrBadRequest
	}

	if err := a.TokenUtil.ParseToken(authTokenRequest.AuthToken); err != nil {
		a.Log.WithError(err).Error("error parsing token")
		return nil, err
	}

	result, err := a.RedisClient.HGetAll(ctx, "authToken:"+authTokenRequest.AuthToken).Result()
	if err != nil {
		a.Log.WithError(err).Error("error getting token")
		return nil, echo.ErrUnauthorized
	}
	
	return result, nil
}

func (a *AuthUseCase) LogoutUser(ctx context.Context, authTokenRequest *model.AuthTokenRequest) (*model.MessageResponse, error) {
	if err := a.Validator.Struct(authTokenRequest); err != nil {
		a.Log.WithError(err).Error("error validating request")
		return nil, echo.ErrBadRequest
	}

	if err := a.RedisClient.Del(ctx, "token").Err(); err != nil {
		a.Log.WithError(err).Error("error deleting token")
		return nil, echo.ErrInternalServerError
	}

	response := &model.MessageResponse{
		Message: "User logout successfully",
	}

	return response, nil
}

func (a *AuthUseCase) GetCSRFToken(ctx context.Context) (string, error) {
	b := make([]byte, 32)
	if _, err := rand.Read(b); err != nil {
		a.Log.WithError(err).Error("error generating random bytes")
		return "", echo.ErrInternalServerError
	}

	csrfToken := base64.RawURLEncoding.EncodeToString(b)

	if err := a.RedisClient.Set(ctx, "csrf:"+csrfToken, "valid", 5*time.Minute).Err(); err != nil {
		a.Log.WithError(err).Error("error set csrf token in redis")
		return "", echo.ErrInternalServerError
	}

	return csrfToken, nil
}