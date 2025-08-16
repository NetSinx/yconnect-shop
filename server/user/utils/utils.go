package utils

import (
	"context"
	"fmt"
	"log"
	"math/rand"
	"time"
	"github.com/golang-jwt/jwt/v5"
	"github.com/redis/go-redis/v9"
)

type CustomClaims struct {
	jwt.RegisteredClaims
	Username  string  `json:"username"`
	Role      string  `json:"role"`
}

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

	ctx, cancel := context.WithTimeout(context.Background(), 5 * time.Second)
	defer cancel()

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
