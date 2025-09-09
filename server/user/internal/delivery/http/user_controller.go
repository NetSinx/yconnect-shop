package http

import (
	"net/http"
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
	var userRequest *model.UserRequest
	if err := c.Bind(&userRequest); err != nil {
		u.Log.WithError(err).Error("error binding request to JSON")
		return err
	}

	username := c.Param("username")
	response, err := u.UserUseCase.UpdateUser(c.Request().Context(), userRequest, username)
	if err != nil {
		u.Log.WithError(err).Error("error updating user")
		return err
	}

	return c.JSON(http.StatusOK, response)
}

func (u *UserController) GetUserByUsername(c echo.Context) error {
	userRequest := new(model.GetUserByUsernameRequest)
	userRequest.Username = c.Param("username")
	if err := c.Bind(userRequest); err != nil {
		u.Log.WithError(err).Error("error binding request to JSON")
		return err
	}

	response, err := u.UserUseCase.GetUserByUsername(c.Request().Context(), userRequest)
	if err != nil {
		u.Log.WithError(err).Error("error getting user")
		return err
	}

	return c.JSON(http.StatusOK, response)
}

func (u *UserController) DeleteUser(c echo.Context) error {
	userRequest := new(model.DeleteUserRequest)
	userRequest.Username = c.Param("username")
	if err := c.Bind(userRequest); err != nil {
		u.Log.WithError(err).Error("error binding request to JSON")
		return err
	}

	if err := u.UserUseCase.DeleteUser(c.Request().Context(), userRequest); err != nil {
		u.Log.WithError(err).Error("error deleting user")
		return err
	}

	return c.JSON(http.StatusOK, &model.MessageResp{
		Message: "user deleted successfully",
	})
}
