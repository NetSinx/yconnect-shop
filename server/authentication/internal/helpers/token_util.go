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
		ID: id,
		Role: role,
		RegisteredClaims: jwt.RegisteredClaims{
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	jwt, err := token.SignedString(t.SecretKey)
	if err != nil {
		return "", err
	}

	t.RedisClient.Set(ctx, "authToken:"+jwt, id, 30*time.Minute)

	return jwt, nil
}

func (t *TokenUtil) ParseToken(authToken string) error {
	token, err := jwt.ParseWithClaims(authToken, &model.CustomClaims{}, func(token *jwt.Token) (any, error) {
		return []byte(t.SecretKey), nil
	})
	if err != nil {
		return err
	}

	claims := token.Claims.(*model.CustomClaims)
	if claims.ExpiresAt.UnixMilli() < time.Now().UnixMilli() {
		return echo.ErrUnauthorized
	}
}