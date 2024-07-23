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

func (u userRepository) LoginUser(userLogin entity.UserLogin) (entity.User, error) {
	var users entity.User
	
	if err := u.DB.Select("password").First(&users, "email = ? OR username = ?", userLogin.UsernameorEmail, userLogin.UsernameorEmail).Error; err != nil {
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
	if err := u.DB.Where("username = ?", username).Updates(&entity.User{Name: users.Name, Username: users.Username, Avatar: users.Avatar, Email: users.Email, Alamat: users.Alamat, NoTelp: users.NoTelp, Password: users.Password, Cart: users.Cart}).Error; err != nil {
		return err
	}
	
	if err := u.DB.First(&users, "username = ?", username).Error; err != nil {
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