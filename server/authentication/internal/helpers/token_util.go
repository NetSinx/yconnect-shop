package helpers

import (
	"context"
	"github.com/golang-jwt/jwt/v5"
	"github.com/redis/go-redis/v9"
	"time"
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
	type CustomClaims struct {
		ID   uint   `json:"id"`
		Role string `json:"role"`
		jwt.RegisteredClaims
	}

	claims := CustomClaims{
		id,
		role,
		jwt.RegisteredClaims{
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
