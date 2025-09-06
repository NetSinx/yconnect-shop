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
	"golang.org/x/crypto/bcrypt"
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
	userEntity.NoHP = userRequest.NoHP

	if err := u.UserRepository.UpdateUser(tx, userEntity, alamatEntity); err != nil {
		u.Log.WithError(err).Error("error updating user")
		return nil, echo.ErrInternalServerError
	}

	if err := tx.Commit().Error; err != nil {
		u.Log.WithError(err).Error("error updating user")
		return nil, echo.ErrInternalServerError
	}

	return converter.UserToResponse(entity), nil
}

func (u *UserUseCase) GetUserByUsername(ctx context.Context, userRequest *model.GetUserByUsernameRequest) (*model.UserResponse, error) {
	tx := u.DB.WithContext(ctx).Begin()
	defer tx.Rollback()

	entity := new(entity.User)
	user, err := u.UserRepository.GetUserByUsername(tx, entity, userRequest.Username)
	if err != nil {
		u.Log.WithError(err).Error("error getting user")
		return nil, echo.ErrNotFound
	}

	return converter.UserToResponse(user), nil
}

func (u *UserUseCase) DeleteUser(user entity.User, username, email string) error {
	getUser, err := u.userRepository.GetUser(user, username, email)
	if err != nil {
		return fmt.Errorf("user tidak ditemukan")
	}

	if getUser.Avatar != "" {
		os.Remove("." + getUser.Avatar)
	}

	err = u.userRepository.DeleteUser(user, username)
	if err != nil && err == gorm.ErrRecordNotFound {
		return fmt.Errorf("user tidak ditemukan")
	} else if err != nil {
		return err
	}

	return nil
}
