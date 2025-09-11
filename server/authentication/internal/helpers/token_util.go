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
	SecretKey   string
	RedisClient *redis.Client
}

func NewTokenUtil(secretKey string, redisClient *redis.Client) *TokenUtil {
	return &TokenUtil{
		SecretKey:   secretKey,
		RedisClient: redisClient,
	}
}

func (t *TokenUtil) CreateToken(ctx context.Context, role string, id uint) (string, error) {	
	claims := model.CustomClaims{
		Role: role,
		RegisteredClaims: jwt.RegisteredClaims{
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour)),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	jwt, err := token.SignedString(t.SecretKey)
	if err != nil {
		return "", err
	}

	valueAuth := map[string]any{"role": role}
	byteValue, _ := json.Marshal(valueAuth)
	t.RedisClient.Set(ctx, "authToken:"+jwt, byteValue, time.Hour)

	return jwt, nil
}

func (t *TokenUtil) ParseToken(authToken string) error {
	token, err := jwt.ParseWithClaims(authToken, &model.CustomClaims{}, func(token *jwt.Token) (any, error) {
		return []byte(t.SecretKey), nil
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