package service

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/NetSinx/yconnect-shop/server/user/model/domain"
	"github.com/NetSinx/yconnect-shop/server/user/model/entity"
	"github.com/NetSinx/yconnect-shop/server/user/repository"
	"github.com/NetSinx/yconnect-shop/server/user/utils"
	"github.com/go-playground/validator/v10"
	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
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
	// users.EmailVerified = false

	if users.Username == "netsinx_15" || users.Email == "yasin03ckm@gmail.com" {
		users.Role = "admin"
	} else {
		users.Role = "member"
	}

	if err := validator.New().Struct(&users); err != nil {
		return err
	}

	passwdHash, _ := bcrypt.GenerateFromPassword([]byte(users.Password), 15)
	users.Password = string(passwdHash)
	// reqUser := []byte(fmt.Sprintf(`{"username": "%s"}`, users.Username))

	// _, err := http.Post("http://kong-gateway:8001/consumers", "application/json", bytes.NewBuffer(reqUser))
	// if err != nil {
	// 	return fmt.Errorf("consumer gagal dibuat")
	// }

	// if users.Username == "netsinx_15" && users.Email == "yasin03ckm@gmail.com" {
	// 	reqJwt := []byte(`{"key": "jwtnetsinxadmin", "secret": "netsinxadmin", "algorithm": "HS512"}`)
	// 	_, err := http.Post(fmt.Sprintf("http://kong-gateway:8001/consumers/%s/jwt", users.Username), "application/json", bytes.NewBuffer(reqJwt))
	// 	if err != nil {
	// 		return fmt.Errorf("token gagal dibuat")
	// 	}
	// } else {
	// 	reqJwt := []byte(`{"key": "jwtyasinganteng", "secret": "yasinganteng15", "algorithm": "HS512"}`)
	// 	_, err := http.Post(fmt.Sprintf("http://kong-gateway:8001/consumers/%s/jwt", users.Username), "application/json", bytes.NewBuffer(reqJwt))
	// 	if err != nil {
	// 		return fmt.Errorf("token gagal dibuat")
	// 	}
	// }

	err := u.userRepository.RegisterUser(users)
	if err != nil {
		return err
	}

	return nil
}

func (u userService) LoginUser(userLogin entity.UserLogin) (string, string, error) {
	users, err := u.userRepository.LoginUser(userLogin)
	if err != nil {
		return "", "", fmt.Errorf("username / email atau password salah")
	}

	jwtToken := utils.JWTAuth(users.Username)

	err = bcrypt.CompareHashAndPassword([]byte(users.Password), []byte(userLogin.Password))
	if err != nil {
		return "", "", err
	}

	key1 := []byte("netsinxadmin")
	key2 := []byte("yasinganteng15")

	token, err := jwt.Parse(jwtToken, func(t *jwt.Token) (interface{}, error) {
		return key1, nil
	})
	if err != nil {
		token, err := jwt.Parse(jwtToken, func(t *jwt.Token) (interface{}, error) {
			return key2, nil
		})
		if err != nil {
			return "", "", err
		} else if token.Valid {
			for claim, value := range token.Claims.(jwt.MapClaims) {
				if claim == "username" {
					return jwtToken, value.(string), nil
				}
			}
		} else {
			return "", "", err
		}
	} else if token.Valid {
		for _, value := range token.Claims.(jwt.MapClaims) {
			if value == "netsinx_15" {
				return jwtToken, value.(string), nil
			}
		}
	}

	return "", "", err
}

func (u userService) ListUsers(users []entity.User) ([]entity.User, error) {
	listUsers, err := u.userRepository.ListUsers(users)
	if err != nil {
		return nil, err
	}

	for i := range listUsers {
		var preloadCart domain.PreloadCarts

		respCart, err := http.Get(fmt.Sprintf("http://cart-service:8083/cart/user/%d", listUsers[i].Id))
		if err != nil {
			return listUsers, nil
		} else if respCart.StatusCode == 200 {
			json.NewDecoder(respCart.Body).Decode(&preloadCart)

			listUsers[i].Cart = preloadCart.Data
		}
	}

	return listUsers, nil
}

func (u userService) UpdateUser(users entity.User, username string) error {
	if err := validator.New().Struct(users); err != nil {
		return err
	}

	passwdHash, _ := bcrypt.GenerateFromPassword([]byte(users.Password), 15)
	users.Password = string(passwdHash)

	err := u.userRepository.UpdateUser(users, username)
	if err != nil {
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

func (u userService) GetUser(users entity.User, username string) (entity.User, error) {
	findUser, err := u.userRepository.GetUser(users, username)
	if err != nil {
		return users, err
	}

	respCart, err := http.Get(fmt.Sprintf("http://cart-service:8083/cart/user/%d", findUser.Id))
	if err != nil {
		return findUser, nil
	} else if respCart.StatusCode == 200 {
		var preloadCart domain.PreloadCarts

		json.NewDecoder(respCart.Body).Decode(&preloadCart)

		findUser.Cart = preloadCart.Data
	}

	return findUser, nil
}

func (u userService) DeleteUser(users entity.User, username string) error {
	var httpClient http.Client

	getUser, err := u.userRepository.GetUser(users, username)
	if err != nil {
		return err
	}

	req, err := http.NewRequest("DELETE", fmt.Sprintf("http://kong-gateway:8001/consumers/%s", getUser.Username), nil)
	if err != nil {
		return nil
	}

	httpClient.Do(req)

	if err := u.userRepository.DeleteUser(users, username); err != nil {
		return err
	}

	return nil
}

func (u userService) Verify(token string) (*jwt.Token, error) {
	jwtKey1 := []byte("netsinxadmin")
	jwtKey2 := []byte("yasinganteng15")

	resultToken, err := jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {
		return jwtKey1, nil
	})
	if err != nil {
		resultToken, err := jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {
			return jwtKey2, nil
		})
		if err != nil {
			return nil, err
		} else if resultToken.Valid {
			return resultToken, nil
		} else {
			return nil, err
		}
	} else if resultToken.Valid {
		return resultToken, nil
	} else {
		return nil, err
	}
}
