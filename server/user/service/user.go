package service

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"github.com/NetSinx/yconnect-shop/server/user/model/domain"
	"github.com/NetSinx/yconnect-shop/server/user/model/entity"
	"github.com/NetSinx/yconnect-shop/server/user/repository"
	"github.com/NetSinx/yconnect-shop/server/user/utils"
	"github.com/go-playground/validator/v10"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type userService struct {
	userRepository repository.UserRepo
}

func UserService(userRepo repository.UserRepo) userService {
	return userService{
		userRepository: userRepo,
	}
}

func (u userService) RegisterUser(users entity.User) error {
	users.EmailVerified = false

	if users.Username == "netsinx_15" || users.Email == "yasin03ckm@gmail.com" {
		users.Role = "admin"
	} else {
		users.Role = "customer"
	}

	if err := validator.New().Struct(users); err != nil {
		return err
	}

	passwdHash, _ := bcrypt.GenerateFromPassword([]byte(users.Password), 8)
	users.Password = string(passwdHash)

	if err := u.userRepository.RegisterUser(users); err != nil {
		return err
	}

	return nil
}

func (u userService) LoginUser(userLogin domain.UserLogin) (string, string, string, error) {
	if err := validator.New().Struct(userLogin); err != nil {
		return "", "", "", err
	}

	users, err := u.userRepository.LoginUser(userLogin)
	if err != nil {
		return "", "", "", fmt.Errorf("username / email atau password salah")
	}

	err = bcrypt.CompareHashAndPassword([]byte(users.Password), []byte(userLogin.Password))
	if err != nil {
		return "", "", "", err
	}
	
	accessToken := utils.GenerateAccessToken(users.Username, users.Role)
	signKey := []byte("yasinnetsinx15")
	token, err := jwt.ParseWithClaims(accessToken, &utils.CustomClaims{}, func(t *jwt.Token) (interface{}, error) {
		return signKey, nil
	})
	if err != nil {
		return "", "", "", err
	}

	refreshToken := utils.GenerateRefreshToken(users.Username, users.Role)
	
	claims := token.Claims.(*utils.CustomClaims)
	
	return accessToken, refreshToken, claims.Username, nil
}

func (u userService) ListUsers(users []entity.User) ([]entity.User, error) {
	listUsers, err := u.userRepository.ListUsers(users)
	if err != nil {
		return nil, err
	}

	return listUsers, nil
}

func (u userService) UpdateUser(user entity.User, username string) error {
	if err := validator.New().Struct(user); err != nil {
		return err
	}

	passwdHash, _ := bcrypt.GenerateFromPassword([]byte(user.Password), 15)
	user.Password = string(passwdHash)

	err := u.userRepository.UpdateUser(user, username)
	if err != nil && err == gorm.ErrRecordNotFound {
		return fmt.Errorf("user tidak ditemukan")
	} else if err != nil && err == gorm.ErrDuplicatedKey {
		return fmt.Errorf("user sudah terdaftar")
	} else if err != nil  {
		return err
	}

	return nil
}

func (u userService) SendOTP(verifyEmail domain.VerifyEmail) (string, error) {
	otpCode := utils.GenerateOTP()
	verifyEmail.OTP = otpCode

	if err := validator.New().Struct(&verifyEmail); err != nil {
		return "", err
	}

	if err := u.userRepository.SendOTP(verifyEmail); err != nil {
		return "", err
	}

	reqBody, err := json.Marshal(&verifyEmail)
	if err != nil {
		return "", err
	}

	resp, err := http.Post("http://mail-service:8085/sendOTP", "application/json", bytes.NewReader(reqBody))
	if err != nil {
		return "", fmt.Errorf("OTP tidak bisa dikirim")
	}

	if err := utils.CacheOTP(otpCode); err != nil {
		return "", err
	}

	var response domain.RespData

	json.NewDecoder(resp.Body).Decode(&response)

	if resp.StatusCode != 200 {
		return "", err
	}

	return "Kode OTP berhasil dikirim.", nil
}

func (u userService) VerifyEmail(verifyEmail domain.VerifyEmail) error {
	if err := validator.New().Struct(&verifyEmail); err != nil {
		return err
	}

	if err := utils.GetOTPFromCache(verifyEmail.OTP); err != nil {
		return err
	}

	if err := u.userRepository.VerifyEmail(verifyEmail); err != nil {
		return err
	}

	return nil
}

func (u userService) GetUser(user entity.User, username string) (entity.User, error) {
	findUser, err := u.userRepository.GetUser(user, username)
	if err != nil {
		return user, fmt.Errorf("user tidak ditemukan")
	}

	return findUser, nil
}

func (u userService) DeleteUser(user entity.User, username string) error {
	getUser, err := u.userRepository.GetUser(user, username)
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
