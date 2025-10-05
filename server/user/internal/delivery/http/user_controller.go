package http

import (
	"fmt"
	"net/http"
	"strconv"
	"github.com/NetSinx/yconnect-shop/server/user/internal/model"
	"github.com/NetSinx/yconnect-shop/server/user/internal/usecase"
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
)

type UserController struct {
	Log         *logrus.Logger
	UserUseCase *usecase.UserUseCase
}

func NewUserController(log *logrus.Logger, userUseCase *usecase.UserUseCase) *UserController {
	return &UserController{
		Log: log,
		UserUseCase: userUseCase,
	}
}

func (u *UserController) RegisterUser(c echo.Context) error {
	userRequest := new(model.RegisterUserRequest)
	if err := c.Bind(userRequest); err != nil {
		u.Log.WithError(err).Error("error binding request to JSON")
		return err
	}
	
	userRegister := &model.RegisterUserEvent{
		NamaLengkap: userRequest.NamaLengkap,
		Username: userRequest.Username,
		Email: userRequest.Email,
		NoHP: userRequest.NoHP,
		Role: userRequest.Role,
	}

	response, err := u.UserUseCase.RegisterUser(c.Request().Context(), userRegister)
	if err != nil {
		u.Log.WithError(err).Error("error registering user")
		return err
	}

	return c.JSON(http.StatusOK, response)
}

func (u *UserController) UpdateUser(c echo.Context) error {
	userRequest := new(model.UpdateUserRequest)
	if err := c.Bind(userRequest); err != nil {
		u.Log.WithError(err).Error("error binding request to JSON")
		return err
	}

	id, _ := strconv.ParseUint(c.Param("id"), 10, strconv.IntSize)
	headerID := c.Request().Header.Get("X-User-ID")
	if fmt.Sprint(id) != headerID {
		u.Log.Error("error validating request header and parameter")
		return echo.ErrForbidden
	}
	
	response, err := u.UserUseCase.UpdateUser(c.Request().Context(), userRequest, uint(id))
	if err != nil {
		u.Log.WithError(err).Error("error updating user")
		return err
	}

	return c.JSON(http.StatusOK, response)
}

func (u *UserController) GetUserByID(c echo.Context) error {
	userRequest := new(model.GetUserByIDRequest)
	
	id, _ := strconv.ParseUint(c.Param("id"), 10, strconv.IntSize)
	headerID := c.Request().Header.Get("X-User-ID")
	if fmt.Sprint(id) != headerID {
		u.Log.Error("error validating request header and parameter")
		return echo.ErrForbidden
	}

	userRequest.ID = uint(id)
	if err := c.Bind(userRequest); err != nil {
		u.Log.WithError(err).Error("error binding request to JSON")
		return err
	}

	response, err := u.UserUseCase.GetUserByID(c.Request().Context(), userRequest)
	if err != nil {
		u.Log.WithError(err).Error("error getting user")
		return err
	}

	return c.JSON(http.StatusOK, response)
}

func (u *UserController) DeleteUser(c echo.Context) error {
	userRequest := new(model.DeleteUserRequest)
	id, _ := strconv.ParseUint(c.Param("id"), 10, strconv.IntSize)
	headerID := c.Request().Header.Get("X-User-ID")
	if fmt.Sprint(id) != headerID {
		u.Log.Error("error validating request header and parameter")
		return echo.ErrForbidden
	}

	userRequest.ID = uint(id)
	if err := c.Bind(userRequest); err != nil {
		u.Log.WithError(err).Error("error binding request to JSON")
		return err
	}

	if err := u.UserUseCase.DeleteUser(c.Request().Context(), userRequest); err != nil {
		u.Log.WithError(err).Error("error deleting user")
		return err
	}

	return c.NoContent(http.StatusNoContent)
}
