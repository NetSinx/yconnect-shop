package repository

import (
	"github.com/NetSinx/yconnect-shop/server/user/model/entity"
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

func (u userRepository) RegisterUser(users entity.User) error {
	if err := u.DB.Create(&users).Error; err != nil {
		return err
	}

	return nil
}

func (u userRepository) LoginUser(userLogin entity.UserLogin) (entity.User, error) {
	var users entity.User

	jwtToken := utils.JWTAuth(userLogin.Username, userLogin.Email)

	u.DB.Where("username = ? OR email = ?", userLogin.Username, userLogin.Email).Updates(&entity.User{Token: jwtToken})
	
	if err := u.DB.Select("username", "password", "token").First(&users, "email = ? OR username = ?", userLogin.Email, userLogin.Username).Error; err != nil {
		return users, err
	}

	return users, nil
}

func (u userRepository) ListUsers(users []entity.User) ([]entity.User, error) {
	if err := u.DB.Find(&users).Error; err != nil {
		return nil, err
	}

	return users, nil
}

func (u userRepository) UpdateUser(users entity.User, username string) error {
	if err := u.DB.Where("username = ?", username).Updates(&entity.User{Name: users.Name, Username: users.Username, Avatar: users.Avatar, Email: users.Email, Alamat: users.Alamat, NoTelp: users.NoTelp, Password: users.Password, Token: users.Token, Cart: users.Cart}).Error; err != nil {
		return err
	}
	
	if err := u.DB.First(&users, "username = ?", username).Error; err != nil {
		return err
	}
	
	return nil
}

func (u userRepository) VerifyEmail(verifyEmail entity.VerifyEmail) error {
	var user entity.User

	if err := u.DB.First(&user, "email = ?", verifyEmail.Email).Error; err != nil {
		return err
	}

	return nil
}

func (u userRepository) GetUser(users entity.User, username string) (entity.User, error) {
	if err := u.DB.First(&users, "username = ?", username).Error; err != nil {
		return users, err
	}

	return users, nil
}

func (u userRepository) DeleteUser(users entity.User, username string) error {
	if err := u.DB.First(&users, "username = ?", username).Error; err != nil {
		return err
	}
	
	if err := u.DB.Delete(&users, "username = ?", username).Error; err != nil {
		return err
	}

	return nil
}