package service

import (
	"errors"
	"github.com/NetSinx/yconnect-shop/user/app/model"
	"github.com/NetSinx/yconnect-shop/user/repository"
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
	if err := u.userRepository.RegisterUser(users); err != nil {
		return errors.New("cannot register user")
	}

	return nil
}

func (u userService) LoginUser(email string) (model.User, error) {
	users, err := u.userRepository.LoginUser(email)
	if err != nil {
		return users, errors.New("cannot login user")
	}

	return users, nil
}

func (u userService) ListUsers(users []model.User) ([]model.User, error) {
	listUsers, err := u.userRepository.ListUsers(users)
	if err != nil {
		return nil, err
	}

	return listUsers, nil
}

func (u userService) UpdateUser(users model.User, id uint) error {
	if err := u.userRepository.UpdateUser(users, id); err != nil {
		return err
	}

	return nil
}

func (u userService) GetUser(users model.User, id uint) (model.User, error) {
	findUser, err := u.userRepository.GetUser(users, id)
	if err != nil {
		return users, errors.New("user cannot be found")
	}

	return findUser, nil
}

func (u userService) DeleteUser(users model.User, id uint) error {
	if err := u.userRepository.DeleteUser(users, id); err != nil {
		return errors.New("user cannot be found")
	}

	return nil
}