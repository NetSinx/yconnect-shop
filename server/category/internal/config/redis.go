package config

import (
	"fmt"
	"github.com/NetSinx/yconnect-shop/server/category/internal/helpers"
	"github.com/go-redis/redis"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

type RedisConfig struct {
	Config *viper.Viper
	Log    *logrus.Logger
}

func NewRedis(config *viper.Viper, log *logrus.Logger) *redis.Client {
	if !config.GetBool("redis.enabled") {
		log.Info("Redis client is disabled")
		return nil
	}

	username := config.GetString("redis.username")
	password := config.GetString("redis.password")
	host := config.GetString("redis.host")
	port := config.GetInt("redis.port")
	db := config.GetString("redis.database")

	url := fmt.Sprintf("redis://%s:%s@%s:%d/%s", username, password, host, port, db)
	opt, err := redis.ParseURL(url)
	helpers.FatalError(log, err, "failed to parse url rabbitmq")

	client := redis.NewClient(opt)

	return client
}