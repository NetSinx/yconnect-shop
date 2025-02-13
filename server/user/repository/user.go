package repository

import (
	"database/sql"
	"time"
	"github.com/NetSinx/yconnect-shop/server/user/model/domain"
	"github.com/NetSinx/yconnect-shop/server/user/model/entity"
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

func (u userRepository) LoginUser(userLogin domain.UserLogin) (entity.User, error) {
	var users entity.User
	
	if err := u.DB.Select("username", "password", "role").First(&users, "email = ? OR username = ?", userLogin.UsernameorEmail, userLogin.UsernameorEmail).Error; err != nil {
		return users, err
	}

	return users, nil
}

func (u userRepository) ListUsers(users []entity.User) ([]entity.User, error) {
	if err := u.DB.Preload("Alamat").Find(&users).Error; err != nil {
		return nil, err
	}

	return users, nil
}

func (u userRepository) UpdateUser(user entity.User, username string) error {
	var userModel entity.User

	if err := u.DB.First(&userModel, "username = ?", username).Error; err != nil {
		return err
	}

	if err := u.DB.Model(&userModel).Updates(&entity.User{Name: user.Name, Username: user.Username, Avatar: user.Avatar, Email: user.Email, NoTelp: user.NoTelp, Role: user.Role, Password: user.Password, EmailVerified: user.EmailVerified, EmailVerifiedAt: user.EmailVerifiedAt}).Error; err != nil {
		return err
	}

	if err := u.DB.Model(&userModel).Association("Alamat").Replace(&user.Alamat); err != nil {
		return err
	}
	
	return nil
}

func (u userRepository) SendOTP(verifyEmail domain.VerifyEmail) error {
	var user entity.User

	if err := u.DB.First(&user, "email = ?", verifyEmail.Email).Error; err != nil {
		return err
	}

	return nil
}

func (u userRepository) VerifyEmail(verifyEmail domain.VerifyEmail) error {
	var user entity.User
	var emailVerified domain.EmailVerified

	emailVerified.EmailVerified = true
	emailVerified.EmailVerifiedAt = sql.NullTime{Time: time.Now().UTC(), Valid: true}

	if err := u.DB.Model(&user).Where("email = ?", verifyEmail.Email).UpdateColumns(&emailVerified).Error; err != nil {
		return err
	}

	return nil
}

func (u userRepository) GetUser(user entity.User, username string) (entity.User, error) {
	if err := u.DB.Preload("Alamat").First(&user, "username = ?", username).Error; err != nil {
		return user, err
	}

	return user, nil
}

func (u userRepository) DeleteUser(user entity.User, username string) error {
	if err := u.DB.First(&user, "username = ?", username).Error; err != nil {
		return err
	}
	
	if err := u.DB.Model(&user).Association("Alamat").Clear(); err != nil {
		return err
	}

	if err := u.DB.Delete(&user).Error; err != nil {
		return err
	}

	return nil
}