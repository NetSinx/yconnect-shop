package config

import (
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func NewLogger(viper *viper.Viper) *logrus.Logger {
	logger := logrus.New()

	logger.SetLevel(logrus.Level(viper.GetInt("log.level")))
	logger.SetFormatter(&logrus.JSONFormatter{})

	return logger
}