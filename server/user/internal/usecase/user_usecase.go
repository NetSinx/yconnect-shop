package usecase

import (
	"context"
	"encoding/json"
	"time"

	"github.com/NetSinx/yconnect-shop/server/user/internal/entity"
	"github.com/NetSinx/yconnect-shop/server/user/internal/model"
	"github.com/NetSinx/yconnect-shop/server/user/internal/model/converter"
	"github.com/NetSinx/yconnect-shop/server/user/internal/repository"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/redis/go-redis/v9"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type UserUseCase struct {
	DB             *gorm.DB
	Log            *logrus.Logger
	Validator      *validator.Validate
	RedisClient    *redis.Client
	UserRepository *repository.UserRepository
}

func NewUserUseCase(db *gorm.DB, log *logrus.Logger, validator *validator.Validate, redisClient *redis.Client, userRepository *repository.UserRepository) *UserUseCase {
	return &UserUseCase{
		DB:             db,
		Log:            log,
		Validator:      validator,
		RedisClient:    redisClient,
		UserRepository: userRepository,
	}
}

func (u *UserUseCase) UpdateUser(ctx context.Context, userRequest *model.UpdateUserRequest, username string) (*model.UserResponse, error) {
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

	userByte, _ := json.Marshal(userEntity)
	if err := u.RedisClient.Set(ctx, "user:"+userRequest.Username, userByte, 20*time.Minute).Err(); err != nil {
		u.Log.WithError(err).Error("error caching user in redis")
		return nil, echo.ErrInternalServerError
	}
	
	if err := tx.Commit().Error; err != nil {
		u.Log.WithError(err).Error("error updating user")
		return nil, echo.ErrInternalServerError
	}

	return converter.UserToResponse(userEntity), nil
}

func (u *UserUseCase) GetUserByUsername(ctx context.Context, userRequest *model.GetUserByUsernameRequest) (*model.UserResponse, error) {
	tx := u.DB.WithContext(ctx).Begin()
	defer tx.Rollback()

	if err := u.Validator.Struct(userRequest); err != nil {
		u.Log.WithError(err).Error("error validating request body")
		return nil, echo.ErrBadRequest
	}

	result, err := u.RedisClient.Get(ctx, "user:"+userRequest.Username).Result()
	if err == nil {
		var userResponse *model.UserResponse
		json.Unmarshal([]byte(result), userResponse)
		return userResponse, nil
	}

	entity := new(entity.User)
	user, err := u.UserRepository.GetUserByUsername(tx, entity, userRequest.Username)
	if err != nil {
		u.Log.WithError(err).Error("error getting user")
		return nil, echo.ErrNotFound
	}

	userByte, _ := json.Marshal(user)
	if err := u.RedisClient.Set(ctx, "user:"+userRequest.Username, userByte, 20*time.Minute).Err(); err != nil {
		u.Log.WithError(err).Error("error caching user in redis")
		return nil, echo.ErrInternalServerError
	}

	if err := tx.Commit().Error; err != nil {
		u.Log.WithError(err).Error("error getting user")
		return nil, echo.ErrInternalServerError
	}

	return converter.UserToResponse(user), nil
}

func (u *UserUseCase) DeleteUser(ctx context.Context, userRequest *model.DeleteUserRequest) error {
	tx := u.DB.WithContext(ctx).Begin()
	defer tx.Rollback()

	if err := u.Validator.Struct(userRequest); err != nil {
		u.Log.WithError(err).Error("error validating request body")
		return echo.ErrBadRequest
	}

	entity := new(entity.User)
	if err := u.RedisClient.GetDel(ctx, "user:"+userRequest.Username).Err(); err != nil {
		if _, err := u.UserRepository.GetUserByUsername(tx, entity, userRequest.Username); err != nil {
			u.Log.WithError(err).Error("error getting user")
			return echo.ErrNotFound
		}
	}

	if err := u.UserRepository.DeleteUser(tx, entity, userRequest.Username); err != nil {
		u.Log.WithError(err).Error("error deleting user")
		return echo.ErrInternalServerError
	}

	return nil
}
