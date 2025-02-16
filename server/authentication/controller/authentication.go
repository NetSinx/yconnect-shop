package controller

import (
	"github.com/NetSinx/yconnect-shop/server/authentication/service"
	"github.com/NetSinx/yconnect-shop/server/authentication/model/domain"
	"github.com/NetSinx/yconnect-shop/server/authentication/utils"
	"github.com/labstack/echo/v4"
	"net/http"
	"time"
)

type authController struct {
	authService service.AuthService
}

func AuthContrllr(authServ service.AuthService) *authController {
	return &authController{
		authService: authServ,
	}
}

func (ac *authController) LoginUser(ctx echo.Context) error {
	var userLogin domain.UserLogin

	if err := ctx.Bind(&userLogin); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, domain.MessageResp{
			Message: err.Error(),
		})
	}

	accessToken, refreshToken, user_id, err := ac.authService.LoginUser(userLogin)
	if err != nil && err.Error() == "email atau password salah" {
		return echo.NewHTTPError(http.StatusUnauthorized, domain.MessageResp{
			Message: err.Error(),
		})
	} else if err != nil && (err.Error() == "failed to decode response" || err.Error() == "failed to parse token with claims") {
		return echo.NewHTTPError(http.StatusInternalServerError, domain.MessageResp{
			Message: err.Error(),
		})
	} else if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, domain.MessageResp{
			Message: err.Error(),
		})
	}

	utils.SetCookies(ctx, "user_session", accessToken, time.Now().Add(15 * time.Minute))
	utils.SetCookies(ctx, "refresh_token", refreshToken, time.Now().Add(1 * time.Hour))
	utils.SetCookies(ctx, "user_id", user_id, time.Now().Add(15 * time.Minute))
	utils.SetCookies(ctx, "tz", time.Now().String(), time.Now().Add(15 * time.Minute))

	return ctx.JSON(http.StatusOK, domain.MessageResp{
		Message: "User berhasil login",
	})
}