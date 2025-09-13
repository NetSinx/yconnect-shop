package helpers

import (
	"context"
	"time"
	"github.com/NetSinx/yconnect-shop/server/authentication/internal/model"
	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
	"github.com/redis/go-redis/v9"
)

type TokenUtil struct {
	AccessKey   []byte
	RefreshKey  []byte
	RedisClient *redis.Client
}

func NewTokenUtil(accessKey []byte, refreshKey []byte, redisClient *redis.Client) *TokenUtil {
	return &TokenUtil{
		AccessKey:   accessKey,
		RefreshKey:  refreshKey,
		RedisClient: redisClient,
	}
}

func (t *TokenUtil) CreateAccessToken(ctx context.Context, id uint, role string) (string, error) {
	accessClaims := model.CustomClaims{
		ID: id,
		Role: role,
		RegisteredClaims: jwt.RegisteredClaims{
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour)),
		},
	}

	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, accessClaims)
	jwtAccess, err := accessToken.SignedString(t.AccessKey)
	if err != nil {
		return "", err
	}

	return jwtAccess, nil
}

func (t *TokenUtil) CreateRefreshToken(ctx context.Context, id uint, role string) (string, error) {
	refreshClaims := model.CustomClaims{
		ID: id,
		Role: role,
		RegisteredClaims: jwt.RegisteredClaims{
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(30*24*time.Hour)),
		},
	}

	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, refreshClaims)
	jwtRefresh, err := refreshToken.SignedString(t.RefreshKey)
	if err != nil {
		return "", err
	}

	return jwtRefresh, nil
}

func (t *TokenUtil) ParseAccessToken(authToken string) (uint, string, error) {
	token, err := jwt.ParseWithClaims(authToken, &model.CustomClaims{}, func(token *jwt.Token) (any, error) {
		return []byte(t.AccessKey), nil
	})
	if err != nil {
		return 0, "", echo.ErrInternalServerError
	}

	if !token.Valid {
		return 0, "", echo.ErrUnauthorized
	}

	claims := token.Claims.(*model.CustomClaims)
	if claims.ExpiresAt.UnixMilli() < time.Now().UnixMilli() {
		return 0, "", echo.ErrUnauthorized
	}

	return claims.ID, claims.Role, nil
}

func (t *TokenUtil) ParseRefreshToken(authToken string) error {
	token, err := jwt.ParseWithClaims(authToken, &model.CustomClaims{}, func(token *jwt.Token) (any, error) {
		return []byte(t.RefreshKey), nil
	})
	if err != nil {
		return echo.ErrInternalServerError
	}

	if !token.Valid {
		return echo.ErrUnauthorized
	}

	claims := token.Claims.(*model.CustomClaims)
	if claims.ExpiresAt.UnixMilli() < time.Now().UnixMilli() {
		return echo.ErrUnauthorized
	}

	return nil
}
