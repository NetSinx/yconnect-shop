package http

import (
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

func (u *UserController) UpdateUser(c echo.Context) error {
	userRequest := new(model.UserRequest)
	if err := c.Bind(userRequest); err != nil {
		u.Log.WithError(err).Error("error binding request to JSON")
		return err
	}

	id, _ := strconv.ParseUint(c.Param("id"), 10, strconv.IntSize)
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
