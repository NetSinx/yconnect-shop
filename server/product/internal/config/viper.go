package config

import "github.com/spf13/viper"

func NewViper() *viper.Viper {
	config := viper.New()

	config.SetConfigName("config")
	config.SetConfigType("json")
	config.AddConfigPath("./../")
	config.AddConfigPath("./")

	err := config.ReadInConfig()
	if err != nil {
		panic(err)
	}

	return config
}
