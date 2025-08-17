package service

import (
	"encoding/json"
	"fmt"
	"net/http"
	"github.com/NetSinx/yconnect-shop/server/authentication/domain"
	"github.com/NetSinx/yconnect-shop/server/authentication/utils"
	"github.com/NetSinx/yconnect-shop/server/user/model/entity"
	"github.com/go-playground/validator/v10"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

type authService struct {
	authServ  AuthService	
}

func AuthServ() *authService {
	var authServ authService
	return &authService{
		authServ: authServ.authServ,
	}
}

func (as *authService) LoginUser(userLogin domain.UserLogin) (string, string, error) {
	if err := validator.New().Struct(userLogin); err != nil {
		respUser, err := http.Get(fmt.Sprintf("http://localhost:8082/user?username=%s", userLogin.UsernameorEmail))
		if err != nil {
			return "", "", err
		} else if respUser.StatusCode == 401 {
			return "", "", fmt.Errorf("email atau password salah")
		} else if respUser.StatusCode == 400 {
			var respBadRequest domain.MessageResp
			if err := json.NewDecoder(respUser.Body).Decode(&respBadRequest); err != nil {
				return "", "", err
			}
			return "", "", fmt.Errorf("%v", respBadRequest.Message)
		}

		var user entity.User
		if err := json.NewDecoder(respUser.Body).Decode(&user); err != nil {
			return "", "", fmt.Errorf("failed to decode response")
		}

		err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(userLogin.Password))
		if err != nil {
			return "", "", fmt.Errorf("email atau password salah")
		}

		accessToken := utils.GenerateAccessToken(user.Username, user.Role)
		if user.Role == "admin" {
			signKey := []byte(utils.AdminJwtKey)
			_, err := jwt.ParseWithClaims(accessToken, &utils.CustomClaims{}, func(t *jwt.Token) (any, error) {
				return signKey, nil
			})
			if err != nil {
				return "", "", fmt.Errorf("failed to parse token with claims")
			}
			refreshToken := utils.GenerateRefreshToken(user.Username, user.Role)

			return accessToken, refreshToken, nil
		}

		signKey := []byte(utils.CustomerJwtKey)
		_, err = jwt.ParseWithClaims(accessToken, &utils.CustomClaims{}, func(t *jwt.Token) (any, error) {
			return signKey, nil
		})
		if err != nil {
			return "", "", fmt.Errorf("failed to parse token with claims")
		}

		refreshToken := utils.GenerateRefreshToken(user.Username, user.Role)

		return accessToken, refreshToken, nil
	}

	respUser, err := http.Get(fmt.Sprintf("http://localhost:8082/user?email=%s", userLogin.UsernameorEmail))
	if err != nil {
		return "", "", err
	} else if respUser.StatusCode == 401 {
		return "", "", fmt.Errorf("email atau password salah")
	} else if respUser.StatusCode == 400 {
		var respBadRequest domain.MessageResp
		if err := json.NewDecoder(respUser.Body).Decode(&respBadRequest); err != nil {
			return "", "", err
		}
		return "", "", fmt.Errorf("%s", respBadRequest.Message)
	}

	var user entity.User
	if err := json.NewDecoder(respUser.Body).Decode(&user); err != nil {
		return "", "", fmt.Errorf("failed to decode response")
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(userLogin.Password))
	if err != nil {
		return "", "", fmt.Errorf("email atau password salah")
	}

	if user.Role == "admin" {
		accessToken := utils.GenerateAccessToken(user.Username, user.Role)
		signKey := []byte(utils.AdminJwtKey)
		_, err := jwt.ParseWithClaims(accessToken, &utils.CustomClaims{}, func(t *jwt.Token) (any, error) {
			return signKey, nil
		})
		if err != nil {
			return "", "", fmt.Errorf("failed to parse token with claims")
		}
	
		refreshToken := utils.GenerateRefreshToken(user.Username, user.Role)
	
		return accessToken, refreshToken, nil
	} else {
		accessToken := utils.GenerateAccessToken(user.Username, user.Role)
		signKey := []byte(utils.CustomerJwtKey)
		_, err := jwt.ParseWithClaims(accessToken, &utils.CustomClaims{}, func(t *jwt.Token) (any, error) {
			return signKey, nil
		})
		if err != nil {
			return "", "", fmt.Errorf("failed to parse token with claims")
		}
	
		refreshToken := utils.GenerateRefreshToken(user.Username, user.Role)
	
		return accessToken, refreshToken, nil
	}
}