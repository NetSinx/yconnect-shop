package repository

import (
	"github.com/NetSinx/yconnect-shop/server/authentication/internal/entity"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type AuthRepository struct {
	Log    *logrus.Logger
}

func NewAuthRepository(log *logrus.Logger) *AuthRepository {
	return &AuthRepository{
		Log: log,
	}
}

func (a *AuthRepository) Create(db *gorm.DB, entity *entity.Authentication) error {
	if err := db.Create(entity).Error; err != nil {
		return err
	}

	return nil
}

func (a *AuthRepository) Update(db *gorm.DB, entity *entity.Authentication) error {
	if err := db.Save(entity).Error; err != nil {
		return err
	}

	return nil
}

func (a *AuthRepository) Delete(db *gorm.DB, entity *entity.Authentication) error {
	if err := db.Delete(entity, "id = ?", entity.ID).Error; err != nil {
		return err
	}

	return nil
}

func (a *AuthRepository) GetByEmail(db *gorm.DB, entity *entity.Authentication, email string) (*entity.Authentication, error) {
	if err := db.Select("id", "role", "password").First(entity, "email = ?", email).Error; err != nil {
		return nil, err
	}

	return entity, nil
}