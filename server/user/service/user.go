package service

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"github.com/NetSinx/yconnect-shop/server/user/app/model"
	"github.com/NetSinx/yconnect-shop/server/user/repository"
	"github.com/NetSinx/yconnect-shop/server/user/utils"
	"github.com/go-playground/validator/v10"
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

func (u userService) RegisterUser(users model.User) error {
	if err := validator.New().Struct(users); err != nil {
		return errors.New("request tidak sesuai")
	}

	passwdHash, _ := bcrypt.GenerateFromPassword([]byte(users.Password), 15)

	users.Password = string(passwdHash)

	reqUser := []byte(fmt.Sprintf(`{"username": "%s"}`, users.Username))

	_, err := http.Post("http://kong-gateway:8001/consumers", "application/json", bytes.NewBuffer(reqUser))
	if err != nil {
		return errors.New("consumer gagal dibuat")
	}

	if users.Username == "netsinx_15" {
		reqJwt := []byte(`{"key": "jwtnetsinxadmin", "secret": "netsinxadmin", "algorithm": "HS512"}`)

		_, err := http.Post(fmt.Sprintf("http://kong-gateway:8001/consumers/%s/jwt", users.Username), "application/json", bytes.NewBuffer(reqJwt))
		if err != nil {
			return errors.New("token gagal dibuat")
		}
	} else {
		reqJwt := []byte(`{"key": "jwtyasinganteng", "secret": "yasinganteng15", "algorithm": "HS512"}`)

		_, err := http.Post(fmt.Sprintf("http://kong-gateway:8001/consumers/%s/jwt", users.Username), "application/json", bytes.NewBuffer(reqJwt))
		if err != nil {
			return errors.New("token gagal dibuat")
		}
	}

	err = u.userRepository.RegisterUser(users)
	if err != nil {
		return err
	}

	return nil
}

func (u userService) LoginUser(userLogin model.UserLogin) (string, error) {
	if userLogin.Email != "" {
		containsAt := false

		for _, word := range userLogin.Email {
			if word == '@' {
				containsAt = true
				break
			}
		}
	
		if !containsAt {
			return "", errors.New("email tidak mengandung karakter '@'")
		}
	}

	users, err := u.userRepository.LoginUser(userLogin)
	if err != nil {
		return "", errors.New("email atau password salah")
	}

	err = bcrypt.CompareHashAndPassword([]byte(users.Password), []byte(userLogin.Password))
	if err != nil {
		return "", err
	}

	return users.Token, nil
}

func (u userService) ListUsers(users []model.User) ([]model.User, error) {
	listUsers, err := u.userRepository.ListUsers(users)
	if err != nil {
		return nil, err
	}

	for i := range listUsers {
		var preloadProduct utils.PreloadProducts
		var preloadCart utils.PreloadCarts
		
		respProduct, err := http.Get(fmt.Sprintf("http://product-service:8081/product/seller/%d", listUsers[i].Id))
		if err != nil {
			return listUsers, nil
		} else if respProduct.StatusCode == 200 {
			json.NewDecoder(respProduct.Body).Decode(&preloadProduct)
			
			listUsers[i].Seller.Product = preloadProduct.Data
		}

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

func (u userService) UpdateUser(users model.User, id string) error {
	if err := validator.New().Struct(users); err != nil {
		return err
	}

	passwdHash, _ := bcrypt.GenerateFromPassword([]byte(users.Password), 15)
	users.Password = string(passwdHash)

	err := u.userRepository.UpdateUser(users, id)
	if err != nil && err == gorm.ErrRecordNotFound {
		return errors.New("user tidak ditemukan")
	} else if err != nil && err != gorm.ErrRecordNotFound {
		return errors.New("user sudah pernah dibuat")
	}

	return nil
}

func (u userService) GetUser(users model.User, id string) (model.User, error) {
	findUser, err := u.userRepository.GetUser(users, id)
	if err != nil {
		return users, err
	}

	respCart, err := http.Get(fmt.Sprintf("http://cart-service:8083/cart/user/%d", findUser.Id))
		if err != nil {
			return findUser, nil
		} else if respCart.StatusCode != 200 {
			var preloadCart utils.PreloadCarts

			json.NewDecoder(respCart.Body).Decode(&preloadCart)
			
			findUser.Cart = preloadCart.Data
		}

	return findUser, nil
}

func (u userService) GetSeller(users model.User, id string) (model.User, error) {
	getSeller, err := u.userRepository.GetSeller(users, id)
	if err != nil {
		return users, err
	}

	resp, err := http.Get(fmt.Sprintf("http://product-service:8081/product/seller/%d", getSeller.Id))
	if err != nil {
		return users, nil
	} else if resp.StatusCode == 200 {
		var preloadProduct utils.PreloadProducts

		json.NewDecoder(resp.Body).Decode(&preloadProduct)
	
		getSeller.Seller.Product = preloadProduct.Data
	}

	return getSeller, nil
}

func (u userService) DeleteUser(users model.User, id string) error {
	var httpClient http.Client
	
	getUser, err := u.userRepository.GetUser(users, id)
	if err != nil {
		return err
	}

	req, _ := http.NewRequest("DELETE", fmt.Sprintf("http://kong-gateway:8001/consumers/%s", getUser.Username), nil)
	
	httpClient.Do(req)
	
	if err := u.userRepository.DeleteUser(users, id); err != nil {
		return errors.New("gagal hapus user")
	}

	return nil
}