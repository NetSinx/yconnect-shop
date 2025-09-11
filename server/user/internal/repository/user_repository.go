package repository

import (
	"github.com/NetSinx/yconnect-shop/server/user/internal/entity"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type UserRepository struct {
	Log *logrus.Logger
}

func NewUserRepository(log *logrus.Logger) *UserRepository {
	return &UserRepository{
		Log: log,
	}
}

func (u *UserRepository) RegisterUser(db *gorm.DB, userEntity *entity.User) error {
	if err := db.Create(userEntity).Error; err != nil {
		return err
	}

	return nil
}

func (u *UserRepository) UpdateUser(db *gorm.DB, userEntity *entity.User) error {
	if err := db.Save(&userEntity.Alamat).Error; err != nil {
		return err
	}

	if err := db.Save(userEntity).Error; err != nil {
		return err
	}
	
	return nil
}

func (u *UserRepository) GetUserByUsername(db *gorm.DB, entity *entity.User, username string) (*entity.User, error) {
	if err := db.Preload("Alamat").First(entity, "username = ?", username).Error; err != nil {
		return nil, err
	}

	return entity, nil
}

func (u *UserRepository) DeleteUser(db *gorm.DB, entity *entity.User, username string) error {
	if err := db.Delete(entity, "username = ?", username).Error; err != nil {
		return err
	}

	return nil
}