package repository

import (
	"github.com/NetSinx/yconnect-shop/server/user/app/model"
	"github.com/NetSinx/yconnect-shop/server/user/utils"
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

func (u userRepository) LoginUser(userLogin model.UserLogin) (model.User, error) {
	var users model.User

	jwtToken := utils.JWTAuth(userLogin.Username, userLogin.Email)

	u.DB.Where("username = ? OR email = ?", userLogin.Username, userLogin.Email).Updates(&model.User{Token: jwtToken})
	
	if err := u.DB.Select("username", "password", "token").First(&users, "email = ? OR username = ?", userLogin.Email, userLogin.Username).Error; err != nil {
		return users, err
	}

	return users, nil
}

func (u userRepository) ListUsers(users []model.User) ([]model.User, error) {
	if err := u.DB.Preload("Seller").Find(&users).Error; err != nil {
		return nil, err
	}

	return users, nil
}

func (u userRepository) UpdateUser(users model.User, id string) error {
	if err := u.DB.Where("id = ?", id).Updates(&model.User{Name: users.Name, Username: users.Username, Avatar: users.Avatar, Email: users.Email, Alamat: users.Alamat, NoTelp: users.NoTelp, Password: users.Password, Token: users.Token, Seller: users.Seller, Cart: users.Cart}).Error; err != nil {
		return err
	}
	
	if err := u.DB.First(&users, "id = ?", id).Error; err != nil {
		return err
	}
	
	return nil
}

func (u userRepository) GetUser(users model.User, id string) (model.User, error) {
	if err := u.DB.First(&users, "id = ?", id).Error; err != nil {
		return users, err
	}

	return users, nil
}

func (u userRepository) GetSeller(users model.User, id string) (model.User, error) {
	if err := u.DB.Preload("Seller").First(&users, "id = ?", id).Error; err != nil {
		return users, err
	}

	return users, nil
}

func (u userRepository) DeleteUser(users model.User, id string) error {
	if err := u.DB.First(&users, "id = ?", id).Error; err != nil {
		return err
	}
	
	if err := u.DB.Delete(&users, "id = ?", id).Error; err != nil {
		return err
	}

	return nil
}