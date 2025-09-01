package http

import (
	"net/http"
	"github.com/NetSinx/yconnect-shop/server/authentication/internal/model"
	"github.com/NetSinx/yconnect-shop/server/authentication/internal/usecase"
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
)

type AuthController struct {
	Log         *logrus.Logger
	AuthUseCase *usecase.AuthUseCase
}

func NewAuthController(log *logrus.Logger, authUseCase *usecase.AuthUseCase) *AuthController {
	return &AuthController{
		Log:         log,
		AuthUseCase: authUseCase,
	}
}

func (a *AuthController) LoginUser(ctx echo.Context) error {
	var loginRequest *model.LoginRequest

	if err := ctx.Bind(loginRequest); err != nil {
		a.Log.WithError(err).Error("error binding request to json")
		return err
	}

	response, err := a.AuthUseCase.LoginUser(ctx.Request().Context(), loginRequest)
	if err != nil {
		a.Log.WithError(err).Error("error user login")
		return err
	}

	return ctx.JSON(http.StatusOK, response)
}
