package usecase

import (
	"context"
	"encoding/json"
	"time"
	"github.com/NetSinx/yconnect-shop/server/user/internal/entity"
	"github.com/NetSinx/yconnect-shop/server/user/internal/gateway/messaging"
	"github.com/NetSinx/yconnect-shop/server/user/internal/model"
	"github.com/NetSinx/yconnect-shop/server/user/internal/model/converter"
	"github.com/NetSinx/yconnect-shop/server/user/internal/repository"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/redis/go-redis/v9"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"gorm.io/gorm"
)

type UserUseCase struct {
	Config *viper.Viper
	DB             *gorm.DB
	Log            *logrus.Logger
	Validator      *validator.Validate
	RedisClient    *redis.Client
	UserRepository *repository.UserRepository
	Publisher *messaging.Publisher
}

func NewUserUseCase(config *viper.Viper, db *gorm.DB, log *logrus.Logger, validator *validator.Validate, redisClient *redis.Client, userRepository *repository.UserRepository, publisher *messaging.Publisher) *UserUseCase {
	return &UserUseCase{
		Config: config,
		DB:             db,
		Log:            log,
		Validator:      validator,
		RedisClient:    redisClient,
		UserRepository: userRepository,
		Publisher: publisher,
	}
}

func (u *UserUseCase) RegisterUser(ctx context.Context, userEvent *model.RegisterUserEvent) error {
	tx := u.DB.WithContext(ctx).Begin()
	defer tx.Rollback()

	entity := &entity.User{
		NamaLengkap: userEvent.NamaLengkap,
		Username: userEvent.Username,
		Email: userEvent.Email,
		Role: userEvent.Role,
		NoHP: userEvent.NoHP,
	}

	if err := u.UserRepository.RegisterUser(tx, entity); err != nil {
		u.Log.WithError(err).Error("error registering user")
		return echo.ErrInternalServerError
	}

	if err := tx.Commit().Error; err != nil {
		u.Log.WithError(err).Error("error registering user")
		return echo.ErrInternalServerError
	}

	return nil
}

func (u *UserUseCase) UpdateUser(ctx context.Context, userRequest *model.UserRequest, username string) (*model.DataResponse, error) {
	tx := u.DB.WithContext(ctx).Begin()
	defer tx.Rollback()

	if err := validator.New().Struct(userRequest); err != nil {
		u.Log.WithError(err).Error("error validating request body")
		return nil, echo.ErrBadRequest
	}

	userEntity := new(entity.User)
	
	var userResponse *model.UserResponse
	result, err := u.RedisClient.Get(ctx, "user:"+userRequest.Username).Result()
	if err != nil {
		user, err := u.UserRepository.GetUserByUsername(tx, userEntity, userRequest.Username)
		if err != nil {
			u.Log.WithError(err).Error("error getting user")
			return nil, echo.ErrNotFound
		}

		userResponse = converter.UserToResponse(user)
	}

	json.Unmarshal([]byte(result), userResponse)

	alamatEntity := &entity.Alamat{
		ID:        userResponse.Alamat.ID,
		NamaJalan: userRequest.Alamat.NamaJalan,
		RT:        userRequest.Alamat.RT,
		RW:        userRequest.Alamat.RW,
		Kelurahan: userRequest.Alamat.Kelurahan,
		Kecamatan: userRequest.Alamat.Kecamatan,
		Kota:      userRequest.Alamat.Kota,
		KodePos:   userRequest.Alamat.KodePos,
		UserID:    userResponse.ID,
	}

	userEntity.ID = userResponse.ID
	userEntity.NamaLengkap = userRequest.NamaLengkap
	userEntity.Username = userRequest.Username
	userEntity.Email = userRequest.Email
	userEntity.Alamat = alamatEntity
	userEntity.NoHP = userRequest.NoHP

	if err := u.UserRepository.UpdateUser(tx, userEntity); err != nil {
		u.Log.WithError(err).Error("error updating user")
		return nil, echo.ErrInternalServerError
	}

	if err := tx.Commit().Error; err != nil {
		u.Log.WithError(err).Error("error updating user")
		return nil, echo.ErrInternalServerError
	}
	
	userByte, _ := json.Marshal(userEntity)
	if err := u.RedisClient.Set(ctx, "user:"+userRequest.Username, userByte, 20*time.Minute).Err(); err != nil {
		u.Log.WithError(err).Error("error caching user in redis")
		return nil, echo.ErrInternalServerError
	}

	response := &model.DataResponse{
		Data: converter.UserToResponse(userEntity),
	}

	return response, nil
}

func (u *UserUseCase) GetUserByUsername(ctx context.Context, userRequest *model.GetUserByUsernameRequest) (*model.DataResponse, error) {
	tx := u.DB.WithContext(ctx).Begin()
	defer tx.Rollback()

	if err := u.Validator.Struct(userRequest); err != nil {
		u.Log.WithError(err).Error("error validating request body")
		return nil, echo.ErrBadRequest
	}

	result, err := u.RedisClient.Get(ctx, "user:"+userRequest.Username).Result()
	if err == nil {
		userResponse := new(model.DataResponse)
		json.Unmarshal([]byte(result), userResponse.Data)
		return userResponse, nil
	}

	entity := new(entity.User)
	user, err := u.UserRepository.GetUserByUsername(tx, entity, userRequest.Username)
	if err != nil {
		u.Log.WithError(err).Error("error getting user")
		return nil, echo.ErrNotFound
	}

	if err := tx.Commit().Error; err != nil {
		u.Log.WithError(err).Error("error getting user")
		return nil, echo.ErrInternalServerError
	}
	
	userByte, _ := json.Marshal(user)
	if err := u.RedisClient.Set(ctx, "user:"+userRequest.Username, userByte, 20*time.Minute).Err(); err != nil {
		u.Log.WithError(err).Error("error caching user in redis")
		return nil, echo.ErrInternalServerError
	}

	response := &model.DataResponse{
		Data: converter.UserToResponse(user),
	}

	return response, nil
}

func (u *UserUseCase) DeleteUser(ctx context.Context, userRequest *model.DeleteUserRequest) error {
	tx := u.DB.WithContext(ctx).Begin()
	defer tx.Rollback()

	if err := u.Validator.Struct(userRequest); err != nil {
		u.Log.WithError(err).Error("error validating request body")
		return echo.ErrBadRequest
	}

	entity := new(entity.User)
	result, err := u.RedisClient.GetDel(ctx, "user:"+userRequest.Username).Result()
	if err != nil {
		user, err := u.UserRepository.GetUserByUsername(tx, entity, userRequest.Username)
		if err != nil {
			u.Log.WithError(err).Error("error getting user")
			return echo.ErrNotFound
		}
		entity = user
	} else if result != "" {
		json.Unmarshal([]byte(result), entity)
	}

	if err := u.UserRepository.DeleteUser(tx, entity, userRequest.Username); err != nil {
		u.Log.WithError(err).Error("error deleting user")
		return echo.ErrInternalServerError
	}

	if err := tx.Commit().Error; err != nil {
		u.Log.WithError(err).Error("error deleting user")
		return echo.ErrInternalServerError
	}

	if u.Config.GetBool("rabitmq.enabled") {
		u.Publisher.Send(ctx, )
	}

	return nil
}
