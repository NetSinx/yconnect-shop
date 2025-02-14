package utils

import (
	"context"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"time"
	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
	"github.com/redis/go-redis/v9"
)

func CacheOTP(otp string) error {
	client := redis.NewClient(&redis.Options{
		Addr: "redis-cache:6379",
		Username: "test",
		Password: "test123",
		DB: 0,
	})
	defer client.Close()

	ctx := context.Background()

	if err := client.Set(ctx, "otp", otp, 2 * time.Minute).Err(); err != nil {
		return err
	}

	return nil
}

func GetOTPFromCache(otp string) error {
	client := redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
		Password: "test123",
		DB: 0,
	})
	defer client.Close()

	ctx := context.Background()

	result := client.Get(ctx, "otp").Val()
	if result != otp {
		return fmt.Errorf("kode OTP tidak valid")
	}

	return nil
}

func GenerateOTP() string {
	var token string
	strGenerator := []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")
	
	rand.Seed(time.Now().Unix())

	for i := 0; i < 8; i++ {
		random := rand.Intn(len(strGenerator))
		token += string(strGenerator[random])
	}

	return token
}

func FailOnError(msg string, err error) {
	if err != nil {
		log.Panicf("%s: %s", msg, err)
	}
}

type CustomClaims struct {
	jwt.RegisteredClaims
	Username  string  `json:"username"`
	Role      string  `json:"role"`
}

func GenerateAccessToken(username, role string) string {
		signingKey := []byte("yasinnetsinx15")

		claims := CustomClaims{
			jwt.RegisteredClaims{
				Issuer: "this is a jwt",
				IssuedAt: jwt.NewNumericDate(time.Now()),
				ExpiresAt: jwt.NewNumericDate(time.Now().Add(30 * time.Minute)),
			},
			username,
			role,
		}

		genToken := jwt.NewWithClaims(jwt.SigningMethodHS512, claims)
		token, _ := genToken.SignedString(signingKey)

		return token
}

func GenerateRefreshToken(username, role string) string {
	signingKey := []byte("adminyasinnetsinx_15")

	claims := CustomClaims{
		jwt.RegisteredClaims{
			Issuer: "this is a jwt",
			IssuedAt: jwt.NewNumericDate(time.Now()),
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(2 * time.Hour)),
		},
		username,
		role,
	}

	genToken := jwt.NewWithClaims(jwt.SigningMethodHS512, claims)
	token, _ := genToken.SignedString(signingKey)

	return token
}

func SetCookies(c echo.Context, name string, value string, time time.Time) {
	cookie := http.Cookie{
		Name: name,
		Value: value,
		Expires: time,
		HttpOnly: true,
		Path: "/",
		SameSite: http.SameSiteStrictMode,
		Secure: true,
	}
	
	c.SetCookie(&cookie)
}
