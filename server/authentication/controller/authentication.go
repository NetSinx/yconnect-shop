package controller

import (
	"net/http"
	"time"
	"github.com/NetSinx/yconnect-shop/server/authentication/domain"
	"github.com/NetSinx/yconnect-shop/server/authentication/service"
	"github.com/NetSinx/yconnect-shop/server/authentication/utils"
	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
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

func (ac *authController) UserLogout(ctx echo.Context) error {
	session, err := ctx.Cookie("user_session")
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, domain.MessageResp{
			Message: "user session in cookie is not available",
		})
	}
	session.Path = "/"
	session.MaxAge = -1

	refreshToken, err := ctx.Cookie("refresh_token")
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, domain.MessageResp{
			Message: "refresh token in cookie is not available",
		})
	}
	refreshToken.Path = "/"
	refreshToken.MaxAge = -1

	user_id, err := ctx.Cookie("user_id")
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, domain.MessageResp{
			Message: "user id in cookie is not available",
		})
	}
	user_id.Path = "/"
	user_id.MaxAge = -1

	tz, err := ctx.Cookie("tz")
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, domain.MessageResp{
			Message: "timezone in cookie is not available",
		})
	}
	tz.Path = "/"
	tz.MaxAge = -1

	ctx.SetCookie(session)
	ctx.SetCookie(refreshToken)
	ctx.SetCookie(user_id)
	ctx.SetCookie(tz)

	return ctx.JSON(http.StatusOK, domain.MessageResp{
		Message: "User berhasil logout",
	})
}

func (ac *authController) RefreshToken(ctx echo.Context) error {
	refreshToken, err := ctx.Cookie("refresh_token")
	if err != nil {
		return echo.NewHTTPError(http.StatusUnauthorized, domain.MessageResp{
			Message: "refresh token not available",
		})
	}

	token, err := jwt.ParseWithClaims(refreshToken.Value, &utils.CustomClaims{}, func(t *jwt.Token) (interface{}, error) {
		return []byte("adminyasinnetsinx_15"), nil
	})
	if err != nil {
		return echo.NewHTTPError(http.StatusUnauthorized, domain.MessageResp{
			Message: err.Error(),
		})
	}
	if !token.Valid {
		return echo.NewHTTPError(http.StatusUnauthorized, domain.MessageResp{
			Message: "your token is invalid",
		})
	}

	claims := token.Claims.(*utils.CustomClaims)
	if claims.Username == "" && claims.Role == "" && claims.Email == "" {
		return echo.NewHTTPError(http.StatusUnauthorized, domain.MessageResp{
			Message: "your claims is invalid",
		})
	}

	newAccessToken := utils.GenerateAccessToken(claims.Username, claims.Email, claims.Role)
	utils.SetCookies(ctx, "user_session", newAccessToken, time.Now().Add(15 * time.Minute))
	utils.SetCookies(ctx, "user_id", claims.Username, time.Now().Add(15 * time.Minute))
	utils.SetCookies(ctx, "tz", time.Now().String(), time.Now().Add(15 * time.Minute))

	return ctx.JSON(http.StatusOK, domain.MessageResp{
		Message: "Access token regenerated successfully",
	})
}