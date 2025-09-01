package http

import (
	"net/http"
	"time"

	"github.com/NetSinx/yconnect-shop/server/authentication/internal/model"
	"github.com/NetSinx/yconnect-shop/server/authentication/internal/usecase"
	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
)

type AuthController struct {
	Log  *logrus.Logger
	AuthUseCase *usecase
}

func AuthHandler(authServ service.AuthService) authHandler {
	return authHandler{
		authService: authServ,
	}
}

func (ac authHandler) LoginUser(ctx echo.Context) error {
	var loginRequest *model.LoginRequest

	if err := ctx.Bind(loginRequest); err != nil {

	}
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