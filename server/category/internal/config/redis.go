package config

import (
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

type RedisConfig struct {
	Config *viper.Viper
	Log    *logrus.Logger
}

func NewRedis(config *viper.Viper, log *logrus.Logger)