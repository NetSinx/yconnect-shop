package usecase

import (
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
		Username: registerRequest.Username,
		Email:    registerRequest.Email,
		Role:     "customer",
		NoHP: registerRequest.NoHP,
		Password: string(hashedPassword),
	}

	id, err := a.AuthRepository.Create(tx, entity)
	if err != nil {
		a.Log.WithError(err).Error("error registering user")
		return nil, echo.ErrInternalServerError
	}

	if err := tx.Commit().Error; err != nil {
		a.Log.WithError(err).Error("error registering user")
		return nil, echo.ErrInternalServerError
	}

	registerUserEvent := &model.RegisterUserEvent{
		NamaLengkap: registerRequest.NamaLengkap,
		Username: registerRequest.Username,
		Email: registerRequest.Email,
		NoHP: registerRequest.NoHP,
		Role: "customer",
	}

	if a.Config.GetBool("rabbitmq.enabled") {
		a.Publisher.Send(ctx, registerUserEvent)
	}

	entity.ID = id
	response := &model.DataResponse[*model.RegisterResponse]{
		Data: converter.UserRegisterToResponse(entity),
	}

	return response, nil
}

func (a *AuthUseCase) LoginUser(ctx context.Context, loginRequest *model.LoginRequest) (*model.LoginResponse, error) {
	tx := a.DB.WithContext(ctx).Begin()
	defer tx.Rollback()

	if err := validator.New().Struct(loginRequest); err != nil {
		a.Log.WithError(err).Error("error validating request body")
		return nil, echo.ErrBadRequest
	}

	authEntity := new(entity.UserAuthentication)
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

	jwtAccess, jwtRefresh, err := a.TokenUtil.CreateToken(ctx, result.Role)
	if err != nil {
		a.Log.WithError(err).Error("error generating jwt token")
		return nil, echo.ErrInternalServerError
	}

	response := &model.LoginResponse{
		Username: result.Username,
		Role: result.Role,
		AuthToken: jwtAccess,
		RefreshToken: jwtRefresh,
	}

	return response, nil
}

func (a *AuthUseCase) Verify(ctx context.Context, authTokenRequest *model.AuthTokenRequest) error {
	if err := a.Validator.Struct(authTokenRequest); err != nil {
		a.Log.WithError(err).Error("error validating request")
		return echo.ErrBadRequest
	}

	if err := a.TokenUtil.ParseAccessToken(authTokenRequest.AuthToken); err != nil {
		a.Log.WithError(err).Error("error parsing token")
		return err
	}

	return nil
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

	var authToken *model.RefreshTokenResponse
	if err := json.Unmarshal([]byte(result), &authToken); err != nil {
		a.Log.WithError(err).Error("error unmarshaling data")
		return nil, echo.ErrInternalServerError
	}

	accessToken, refreshToken, err := a.TokenUtil.CreateToken(ctx, authToken.Role)
	if err != nil {
		a.Log.WithError(err).Error("error generating jwt token")
		return nil, err
	}

	response := &model.TokenResponse{
		AccessToken: accessToken,
		RefreshToken: refreshToken,
	}

	return response, nil
}

func (a *AuthUseCase) LogoutUser(ctx context.Context, authTokenRequest *model.AuthTokenRequest) error {
	if err := a.Validator.Struct(authTokenRequest); err != nil {
		a.Log.WithError(err).Error("error validating request")
		return echo.ErrBadRequest
	}

	if err := a.RedisClient.Del(ctx, "authToken:"+authTokenRequest.AuthToken).Err(); err != nil {
		a.Log.WithError(err).Error("error deleting token")
		return echo.ErrInternalServerError
	}

	return nil
}

func (a *AuthUseCase) DeleteUserAuthentication(ctx context.Context, deleteUserEvent *model.DeleteUserEvent) error {
	tx := a.DB.WithContext(ctx).Begin()
	defer tx.Rollback()

	entity := new(entity.UserAuthentication)
	result, err := a.AuthRepository.GetByEmail(tx, entity, deleteUserEvent.Email)
	if err != nil {
		a.Log.WithError(err).Error("error getting user authentication")
		return echo.ErrNotFound
	}

	if err := a.AuthRepository.Delete(tx, result); err != nil {
		a.Log.WithError(err).Error("error deleting user authentication")
		return echo.ErrInternalServerError
	}

	if err := tx.Commit().Error; err != nil {
		a.Log.WithError(err).Error("error deleting user authentication")
		return echo.ErrInternalServerError
	}

	return nil
}
