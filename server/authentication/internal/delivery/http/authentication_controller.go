package http

import (
	"fmt"
	"net/http"
	"strings"
	"time"
	"github.com/NetSinx/yconnect-shop/server/authentication/internal/helpers"
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

func (a *AuthController) RegisterUser(ctx echo.Context) error {
	registerRequest := new(model.RegisterRequest)
	if err := ctx.Bind(registerRequest); err != nil {
		a.Log.WithError(err).Error("error binding request to json")
		return err
	}

	response, err := a.AuthUseCase.RegisterUser(ctx.Request().Context(), registerRequest)
	if err != nil {
		a.Log.WithError(err).Error("error registering user")
		return err
	}

	return ctx.JSON(http.StatusOK, response)
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

	helpers.SetCookie(ctx, "auth_token", response.RefreshToken, time.Now().Add(30*24*time.Hour))

	return ctx.JSON(http.StatusOK, &model.AuthenticationResponse{
		AuthToken: response.AccessToken,
	})
}

func (a *AuthController) Verify(ctx echo.Context) error {
	authTokenRequest := strings.Split(ctx.Request().Header.Get("Authorization"), " ")[1]
	if authTokenRequest == "" {
		a.Log.Error("error getting auth token in cookie")
		return echo.ErrBadRequest
	}

	authRequest := &model.AuthTokenRequest{
		AuthToken: authTokenRequest,
	}

	id, role, err := a.AuthUseCase.Verify(ctx.Request().Context(), authRequest)
	if err != nil {
		a.Log.WithError(err).Error("error verify user")
		return err
	}

	ctx.Response().Header().Add("X-User-ID", fmt.Sprint(id))
	ctx.Response().Header().Add("X-User-Role", role)

	return ctx.NoContent(http.StatusNoContent)
}

func (a *AuthController) RefreshToken(ctx echo.Context) error {
	refreshToken, err := ctx.Cookie("auth_token")
	if err != nil {
		a.Log.WithError(err).Error("error getting refresh token in cookie")
		return echo.ErrBadRequest
	}

	refreshTokenRequest := &model.AuthTokenRequest{
		AuthToken: refreshToken.Value,
	}

	response, err := a.AuthUseCase.RefreshToken(ctx.Request().Context(), refreshTokenRequest)
	if err != nil {
		a.Log.WithError(err).Error("error generating jwt token")
		return err
	}

	return ctx.JSON(http.StatusOK, &model.AuthenticationResponse{
		AuthToken: response.AccessToken,
	})
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

	helpers.SetCookie(ctx, "auth_token", "", time.Now())

	return ctx.NoContent(http.StatusNoContent)
}

func (a *AuthController) GetCSRFToken(ctx echo.Context) error {
	csrfToken := fmt.Sprintf("%v", ctx.Get("csrf_token"))

	return ctx.JSON(http.StatusOK, map[string]string{
		"csrf_token": csrfToken,
	})
}
