package repository

import (
	"github.com/NetSinx/yconnect-shop/user/app/model"
	"gorm.io/gorm"
)

type userRepository struct {
	DB *gorm.DB
}

func UserRepository(db *gorm.DB) userRepository {
	return userRepository{
		DB: db,
	}
}

func (u userRepository) RegisterUser(users model.User) error {
	if err := u.DB.Create(&users).Error; err != nil {
		return err
	}

	return nil
}

func (u userRepository) LoginUser(email string) (model.User, error) {
	var users model.User

	if err := u.DB.Select("username", "password").First(&users, "email = ?", email).Error; err != nil {
		return users, err
	}

	return users, nil
}

func (u userRepository) ListUsers(users []model.User) ([]model.User, error) {
	if err := u.DB.Find(&users).Error; err != nil {
		return nil, err
	}

	return users, nil
}

func (u userRepository) FindUser(users model.User, id uint) (model.User, error) {
	if err := u.DB.First(&users, "id = ?", id).Error; err != nil {
		return users, err
	}

	return users, nil
}

func (u userRepository) UpdateUser(users model.User, id uint) error {
	if err := u.DB.Where("id = ?", id).Updates(&users).Error; err != nil {
		return err
	}
	
	if err := u.DB.First(&users, "id = ?", id).Error; err != nil {
		return err
	}

	return nil
}

func (u userRepository) DeleteUser(users model.User, id uint) error {
	if err := u.DB.Where("id = ?", id).Delete(&users).Error; err != nil {
		return err
	}

	return nil
}