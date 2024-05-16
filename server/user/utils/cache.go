package utils

import (
	"context"
	"fmt"
	"time"
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