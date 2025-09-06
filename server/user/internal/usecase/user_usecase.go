package usecase

import (
	"context"

	"github.com/NetSinx/yconnect-shop/server/user/internal/entity"
	"github.com/NetSinx/yconnect-shop/server/user/internal/model"
	"github.com/NetSinx/yconnect-shop/server/user/internal/model/converter"
	"github.com/NetSinx/yconnect-shop/server/user/internal/repository"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type UserUseCase struct {
	DB             *gorm.DB
	Log            *logrus.Logger
	Validator      *validator.Validate
	UserRepository *repository.UserRepository
}

func NewUserUseCase(db *gorm.DB, log *logrus.Logger, validator *validator.Validate, userRepository *repository.UserRepository) *UserUseCase {
	return &UserUseCase{
		DB: db,
		Log: log,
		Validator: validator,
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
	user, err := u.UserRepository.GetUserByUsername(tx, userEntity, userRequest.Username)
	if err != nil {
		u.Log.WithError(err).Error("error getting user")
		return nil, echo.ErrNotFound
	}

	alamatEntity := &entity.Alamat{
		ID: user.Alamat.ID,
		NamaJalan: userRequest.Alamat.NamaJalan,
		RT: userRequest.Alamat.RT,
		RW: userRequest.Alamat.RW,
		Kelurahan: userRequest.Alamat.Kelurahan,
		Kecamatan: userRequest.Alamat.Kecamatan,
		Kota: userRequest.Alamat.Kota,
		KodePos: userRequest.Alamat.KodePos,
		UserID: user.ID,
	}

	userEntity.ID = user.ID
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

	return converter.UserToResponse(userEntity), nil
}

func (u *UserUseCase) GetUserByUsername(ctx context.Context, userRequest *model.GetUserByUsernameRequest) (*model.UserResponse, error) {
	tx := u.DB.WithContext(ctx).Begin()
	defer tx.Rollback()

	
	if err := u.Validator.Struct(userRequest); err != nil {
		u.Log.WithError(err).Error("error validating request body")
		return nil, echo.ErrBadRequest
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
	if _, err := u.UserRepository.GetUserByUsername(tx, entity, userRequest.Username); err != nil {
		u.Log.WithError(err).Error("error getting user")
		return echo.ErrNotFound
	}

	if err := u.UserRepository.DeleteUser(tx, entity, userRequest.Username); err != nil {
		u.Log.WithError(err).Error("error deleting user")
		return echo.ErrInternalServerError
	}

	return nil
}
