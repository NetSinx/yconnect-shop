package http

import (
	"fmt"
	"net/http"
	"time"

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
	loginRequest := new(model.LoginRequest)
	if err := ctx.Bind(loginRequest); err != nil {
		a.Log.WithError(err).Error("error binding request to json")
		return err
	}

	response, err := a.AuthUseCase.LoginUser(ctx.Request().Context(), loginRequest)
	if err != nil {
		a.Log.WithError(err).Error("error user login")
		return err
	}

	ctx.SetCookie(&http.Cookie{
		Name: "auth_token",
		Path: "/",
		Value: response.AuthToken,
		Secure: true,
		HttpOnly: true,
		Expires: time.Now().Add(30 * time.Minute),
		SameSite: http.SameSiteLaxMode,
	})

	return ctx.NoContent(http.StatusNoContent)
}

func (a *AuthController) Verify(ctx echo.Context) error {
	authTokenRequest, err := ctx.Cookie("auth_token")
	if err != nil {
		a.Log.WithError(err).Error("error getting auth token in cookie")
		return echo.ErrBadRequest
	}

	authRequest := &model.AuthTokenRequest{
		AuthToken: authTokenRequest.Value,
	}

	if err := a.AuthUseCase.Verify(ctx.Request().Context(), authRequest); err != nil {
		a.Log.WithError(err).Error("error verify user")
		return err
	}

	return ctx.NoContent(http.StatusNoContent)
}

func (a *AuthController) LogoutUser(ctx echo.Context) error {
	authTokenRequest, err := ctx.Cookie("auth_token")
	if err != nil {
		a.Log.WithError(err).Error("error getting auth token in cookie")
		return echo.ErrBadRequest
	}

	authRequest := &model.AuthTokenRequest{
		AuthToken: authTokenRequest.Value,
	}

	if err := a.AuthUseCase.LogoutUser(ctx.Request().Context(), authRequest); err != nil {
		a.Log.WithError(err).Error("error logout user")
		return err
	}

	return ctx.NoContent(http.StatusNoContent)
}

func (a *AuthController) GetCSRFToken(ctx echo.Context) error {
	csrfToken := fmt.Sprintf("%v", ctx.Get("csrf_token"))

	ctx.SetCookie(&http.Cookie{
		Name:     "csrf_token",
		Path:     "/",
		Value:    csrfToken,
		HttpOnly: true,
		Secure:   true,
		SameSite: http.SameSiteStrictMode,
	})

	return ctx.NoContent(http.StatusNoContent)
}
