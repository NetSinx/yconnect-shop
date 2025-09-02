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

func (a *AuthController) Verify(ctx echo.Context) error {
	authTokenRequest := &model.AuthTokenRequest{
		AuthToken: ctx.Request().Header.Get("Authorization"),
	}

	response, err := a.AuthUseCase.Verify(ctx.Request().Context(), authTokenRequest)
	if err != nil {
		a.Log.WithError(err).Error("error verify user")
		return err
	}

	return ctx.JSON(http.StatusOK, response)
}

func (a *AuthController) LogoutUser(ctx echo.Context) error {
	authTokenRequest := &model.AuthTokenRequest{
		AuthToken: ctx.Request().Header.Get("Authorization"),
	}

	response, err := a.AuthUseCase.LogoutUser(ctx.Request().Context(), authTokenRequest)
	if err != nil {
		a.Log.WithError(err).Error("error logout user")
		return err
	}

	return ctx.JSON(http.StatusOK, response)
}

func (a *AuthController) GetCSRFToken(ctx echo.Context) error {
	csrfToken := ctx.Get("csrf_token")

	return ctx.JSON(http.StatusOK, map[string]any{
		"csrf_token": csrfToken,
	})
}