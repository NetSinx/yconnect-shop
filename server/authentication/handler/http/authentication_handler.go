package http

import (
	"net/http"
	"time"
	"github.com/NetSinx/yconnect-shop/server/authentication/handler/dto"
	"github.com/NetSinx/yconnect-shop/server/authentication/service"
	"github.com/NetSinx/yconnect-shop/server/authentication/helpers"
	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
)

type AuthHandl interface {
	LoginUser(ctx echo.Context) error
	LogoutUser(ctx echo.Context) error
	RefreshToken(ctx echo.Context) error
}

type authHandler struct {
	authService service.AuthService
}

func AuthHandler(authServ service.AuthService) authHandler {
	return authHandler{
		authService: authServ,
	}
}

func (ac authHandler) LoginUser(ctx echo.Context) error {
	var userLogin dto.UserLogin

	if err := ctx.Bind(&userLogin); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, dto.MessageResp{
			Message: err.Error(),
		})
	}

	accessToken, refreshToken, err := ac.authService.LoginUser(userLogin)
	if err != nil && err.Error() == "email atau password salah" {
		return echo.NewHTTPError(http.StatusUnauthorized, dto.MessageResp{
			Message: err.Error(),
		})
	} else if err != nil && (err.Error() == "failed to decode response" || err.Error() == "failed to parse token with claims") {
		return echo.NewHTTPError(http.StatusInternalServerError, dto.MessageResp{
			Message: err.Error(),
		})
	} else if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, dto.MessageResp{
			Message: err.Error(),
		})
	}

	helpers.SetCookies(ctx, "user_session", accessToken, time.Now().Add(15 * time.Minute))
	helpers.SetCookies(ctx, "refresh_token", refreshToken, time.Now().Add(2 * time.Hour))

	return ctx.JSON(http.StatusOK, dto.MessageResp{
		Message: "User berhasil login",
	})
}

func (ac authHandler) UserLogout(ctx echo.Context) error {
	session, err := ctx.Cookie("user_session")
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, dto.MessageResp{
			Message: "user session in cookie is not available",
		})
	}
	session.Path = "/"
	session.MaxAge = -1

	refreshToken, err := ctx.Cookie("refresh_token")
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, dto.MessageResp{
			Message: "refresh token in cookie is not available",
		})
	}
	refreshToken.Path = "/"
	refreshToken.MaxAge = -1

	ctx.SetCookie(session)
	ctx.SetCookie(refreshToken)

	return ctx.JSON(http.StatusOK, dto.MessageResp{
		Message: "User berhasil logout",
	})
}

func (ac authHandler) RefreshToken(ctx echo.Context) error {
	refreshToken, err := ctx.Cookie("refresh_token")
	if err != nil {
		return echo.NewHTTPError(http.StatusUnauthorized, dto.MessageResp{
			Message: "refresh token not available",
		})
	}

	token, err := jwt.ParseWithClaims(refreshToken.Value, &helpers.CustomClaims{}, func(t *jwt.Token) (any, error) {
		return []byte(helpers.AdminJwtKey), nil
	})
	if err != nil {
		return echo.NewHTTPError(http.StatusUnauthorized, dto.MessageResp{
			Message: err.Error(),
		})
	}
	if !token.Valid {
		token, err := jwt.ParseWithClaims(refreshToken.Value, &helpers.CustomClaims{}, func(t *jwt.Token) (any, error) {
			return []byte(helpers.CustomerJwtKey), nil
		})
		if err != nil {
			return echo.NewHTTPError(http.StatusUnauthorized, dto.MessageResp{
				Message: err.Error(),
			})
		}
		if !token.Valid {
			return echo.NewHTTPError(http.StatusUnauthorized, dto.MessageResp{
				Message: "your token is invalid",
			})
		}

		claims := token.Claims.(*helpers.CustomClaims)
		if claims.Username == "" && claims.Role == "" {
			return echo.NewHTTPError(http.StatusUnauthorized, dto.MessageResp{
				Message: "your claims is invalid",
			})
		}

		newAccessToken := helpers.GenerateAccessToken(claims.Username, claims.Role)
		helpers.SetCookies(ctx, "user_session", newAccessToken, time.Now().Add(15 * time.Minute))

		return ctx.JSON(http.StatusOK, dto.MessageResp{
			Message: "Access token regenerated successfully",
		})
	}

	claims := token.Claims.(*helpers.CustomClaims)
	if claims.Username == "" && claims.Role == "" {
		return echo.NewHTTPError(http.StatusUnauthorized, dto.MessageResp{
			Message: "your claims is invalid",
		})
	}

	newAccessToken := helpers.GenerateAccessToken(claims.Username, claims.Role)
	helpers.SetCookies(ctx, "user_session", newAccessToken, time.Now().Add(15 * time.Minute))

	return ctx.JSON(http.StatusOK, dto.MessageResp{
		Message: "Access token regenerated successfully",
	})
}