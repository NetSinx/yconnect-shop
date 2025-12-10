package usecase

import (
	"bytes"
	"context"
	"encoding/json"
	"github.com/NetSinx/yconnect-shop/server/authentication/internal/entity"
	"github.com/NetSinx/yconnect-shop/server/authentication/internal/gateway/messaging"
	"github.com/NetSinx/yconnect-shop/server/authentication/internal/helpers"
	"github.com/NetSinx/yconnect-shop/server/authentication/internal/model"
	"github.com/NetSinx/yconnect-shop/server/authentication/internal/model/converter"
	"github.com/NetSinx/yconnect-shop/server/authentication/internal/repository"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/redis/go-redis/v9"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"net/http"
	"time"
)

type AuthUseCase struct {
	Config         *viper.Viper
	DB             *gorm.DB
	Log            *logrus.Logger
	Validator      *validator.Validate
	Publisher      *messaging.Publisher
	RedisClient    *redis.Client
	AuthRepository *repository.AuthRepository
	TokenUtil      *helpers.TokenUtil
}

func NewAuthUseCase(config *viper.Viper, db *gorm.DB, log *logrus.Logger, validator *validator.Validate, publisher *messaging.Publisher, redisClient *redis.Client, authRepository *repository.AuthRepository, tokenUtil *helpers.TokenUtil) *AuthUseCase {
	return &AuthUseCase{
		Config:         config,
		DB:             db,
		Log:            log,
		Validator:      validator,
		Publisher:      publisher,
		RedisClient:    redisClient,
		AuthRepository: authRepository,
		TokenUtil:      tokenUtil,
	}
}

func (a *AuthUseCase) RegisterUser(ctx context.Context, registerRequest *model.RegisterRequest) (*model.DataResponse[*model.RegisterResponse], error) {
	tx := a.DB.WithContext(ctx).Begin()
	defer tx.Rollback()

	if err := a.Validator.Struct(registerRequest); err != nil {
		a.Log.WithError(err).Error("error validating request body")
		return nil, echo.ErrBadRequest
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(registerRequest.Password), bcrypt.DefaultCost)
	if err != nil {
		a.Log.WithError(err).Error("error hashing password")
		return nil, echo.ErrInternalServerError
	}

	entity := &entity.UserAuthentication{
		NamaLengkap: registerRequest.NamaLengkap,
		Username:    registerRequest.Username,
		Email:       registerRequest.Email,
		Role:        "customer",
		NoHP:        registerRequest.NoHP,
		Password:    string(hashedPassword),
	}

	if err := a.AuthRepository.Create(tx, entity); err != nil {
		a.Log.WithError(err).Error("error registering user")
		return nil, echo.ErrInternalServerError
	}

	if err := tx.Commit().Error; err != nil {
		a.Log.WithError(err).Error("error registering user")
		return nil, echo.ErrInternalServerError
	}

	registerUserEvent := &model.RegisterUserEvent{
		NamaLengkap: registerRequest.NamaLengkap,
		Username:    registerRequest.Username,
		Email:       registerRequest.Email,
		NoHP:        registerRequest.NoHP,
		Role:        "customer",
	}

	if a.Config.GetBool("rabbitmq.enabled") {
		a.Publisher.Send(ctx, registerUserEvent)
	} else {
		dataByte, err := json.Marshal(registerUserEvent)
		if err != nil {
			a.Log.WithError(err).Error("error marshaling data")
			return nil, echo.ErrInternalServerError
		}

		if _, err = http.Post("http://user-service:8082/user/register", "application/json", bytes.NewReader(dataByte)); err != nil {
			a.Log.WithError(err).Error("error getting response from user service")
			return nil, echo.ErrInternalServerError
		}
	}

	response := &model.DataResponse[*model.RegisterResponse]{
		Data: converter.UserRegisterToResponse(entity),
	}

	return response, nil
}

func (a *AuthUseCase) LoginUser(ctx context.Context, loginRequest *model.LoginRequest) (*model.DataResponse[*model.TokenResponse], error) {
	tx := a.DB.WithContext(ctx).Begin()
	defer tx.Rollback()

	if err := validator.New().Struct(loginRequest); err != nil {
		a.Log.WithError(err).Error("error validating request body")
		return nil, echo.ErrBadRequest
	}

	authEntity := new(entity.UserAuthentication)
	if err := a.AuthRepository.GetByEmail(tx, authEntity, loginRequest.Email); err != nil {
		a.Log.WithError(err).Error("error getting user")
		return nil, echo.ErrUnauthorized
	}

	if err := tx.Commit().Error; err != nil {
		a.Log.WithError(err).Error("error commit the changes in a transaction")
		return nil, echo.ErrInternalServerError
	}

	if err := bcrypt.CompareHashAndPassword([]byte(authEntity.Password), []byte(loginRequest.Password)); err != nil {
		a.Log.WithError(err).Error("error to compare hash and password")
		return nil, echo.ErrUnauthorized
	}

	jwtRefresh, err := a.TokenUtil.CreateRefreshToken(ctx, authEntity.ID, authEntity.Role)
	if err != nil {
		a.Log.WithError(err).Error("error generating jwt token")
		return nil, echo.ErrInternalServerError
	}

	jwtAccess, err := a.TokenUtil.CreateAccessToken(ctx, authEntity.ID, authEntity.Role)
	if err != nil {
		a.Log.WithError(err).Error("error generating jwt token")
		return nil, echo.ErrInternalServerError
	}

	if a.Config.GetBool("redis.enabled") {
		valueAuth := map[string]any{"id": authEntity.ID, "role": authEntity.Role}
		byteValue, err := json.Marshal(valueAuth)
		if err != nil {
			a.Log.WithError(err).Error("error marshaling json data")
			return nil, echo.ErrInternalServerError
		}
		a.RedisClient.Set(ctx, "refresh_token:"+jwtRefresh, byteValue, 30*24*time.Hour)
	}

	response := &model.DataResponse[*model.TokenResponse]{
		Data: &model.TokenResponse{
			AccessToken:  jwtAccess,
			RefreshToken: jwtRefresh,
		},
	}

	return response, nil
}

func (a *AuthUseCase) Verify(ctx context.Context, authTokenRequest *model.AuthTokenRequest) (uint, string, error) {
	if err := a.Validator.Struct(authTokenRequest); err != nil {
		a.Log.WithError(err).Error("error validating request")
		return 0, "", echo.ErrBadRequest
	}

	id, role, err := a.TokenUtil.ParseAccessToken(authTokenRequest.AuthToken)
	if err != nil {
		a.Log.WithError(err).Error("error parsing token")
		return 0, "", echo.ErrUnauthorized
	}

	return id, role, nil
}

func (a *AuthUseCase) RefreshToken(ctx context.Context, authTokenRequest *model.AuthTokenRequest) (*model.TokenResponse, error) {
	if err := a.Validator.Struct(authTokenRequest); err != nil {
		a.Log.WithError(err).Error("error validating request")
		return nil, echo.ErrBadRequest
	}

	if err := a.TokenUtil.ParseRefreshToken(authTokenRequest.AuthToken); err != nil {
		a.Log.WithError(err).Error("error parsing refresh token")
		return nil, err
	}

	result, err := a.RedisClient.Get(ctx, "refresh_token:"+authTokenRequest.AuthToken).Result()
	if err != nil {
		a.Log.WithError(err).Error("error getting token")
		return nil, echo.ErrUnauthorized
	}

	refreshTokenResponse := new(model.RefreshTokenResponse)
	if err := json.Unmarshal([]byte(result), refreshTokenResponse); err != nil {
		a.Log.WithError(err).Error("error unmarshaling data")
		return nil, echo.ErrInternalServerError
	}

	accessToken, err := a.TokenUtil.CreateAccessToken(ctx, refreshTokenResponse.ID, refreshTokenResponse.Role)
	if err != nil {
		a.Log.WithError(err).Error("error generating jwt token")
		return nil, err
	}

	response := &model.TokenResponse{
		AccessToken: accessToken,
	}

	return response, nil
}

func (a *AuthUseCase) LogoutUser(ctx context.Context, authTokenRequest *model.AuthTokenRequest) error {
	if err := a.Validator.Struct(authTokenRequest); err != nil {
		a.Log.WithError(err).Error("error validating request")
		return echo.ErrBadRequest
	}

	if err := a.RedisClient.Del(ctx, "refresh_token:"+authTokenRequest.AuthToken).Err(); err != nil {
		a.Log.WithError(err).Error("error deleting token")
		return echo.ErrInternalServerError
	}

	return nil
}

func (a *AuthUseCase) DeleteUserAuthentication(ctx context.Context, deleteUserEvent *model.DeleteUserEvent) error {
	tx := a.DB.WithContext(ctx).Begin()
	defer tx.Rollback()

	entity := new(entity.UserAuthentication)
	if err := a.AuthRepository.GetByEmail(tx, entity, deleteUserEvent.Email); err != nil {
		a.Log.WithError(err).Error("error getting user authentication")
		return echo.ErrNotFound
	}

	if err := a.AuthRepository.Delete(tx, entity); err != nil {
		a.Log.WithError(err).Error("error deleting user authentication")
		return echo.ErrInternalServerError
	}

	if err := tx.Commit().Error; err != nil {
		a.Log.WithError(err).Error("error deleting user authentication")
		return echo.ErrInternalServerError
	}

	return nil
}
