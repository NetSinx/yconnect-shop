package helpers

import (
	"context"
	"encoding/json"
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

func (t *TokenUtil) CreateToken(ctx context.Context, role string) (string, string, error) {
	refreshClaims := model.CustomClaims{
		Role: role,
		RegisteredClaims: jwt.RegisteredClaims{
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(30*24*time.Hour)),
		},
	}

	accessClaims := model.CustomClaims{
		Role: role,
		RegisteredClaims: jwt.RegisteredClaims{
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour)),
		},
	}

	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, refreshClaims)
	jwtRefresh, err := refreshToken.SignedString(t.RefreshKey)
	if err != nil {
		return "", "", err
	}

	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, accessClaims)
	jwtAccess, err := accessToken.SignedString(t.AccessKey)
	if err != nil {
		return "", "", err
	}

	valueAuth := map[string]any{"role": role}
	byteValue, _ := json.Marshal(valueAuth)
	t.RedisClient.Set(ctx, "refresh_token:"+jwtRefresh, byteValue, time.Hour)

	return jwtAccess, jwtRefresh, nil
}

func (t *TokenUtil) ParseAccessToken(authToken string) error {
	token, err := jwt.ParseWithClaims(authToken, &model.CustomClaims{}, func(token *jwt.Token) (any, error) {
		return []byte(t.AccessKey), nil
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
