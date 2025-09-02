package main

import (
	"fmt"
	"github.com/NetSinx/yconnect-shop/server/authentication/internal/config"
	"github.com/NetSinx/yconnect-shop/server/authentication/internal/helpers"
)

func main() {
	viperConfig := config.NewViper()
	log := config.NewLogger(viperConfig)
	db := config.NewDatabase(viperConfig, log)
	validator := config.NewValidator()
	app := config.NewEcho()
	redis := config.NewRedis(viperConfig, log)
	tokenUtil := helpers.NewTokenUtil("rahasiadeh", redis)

	config.BootstrapApp(&config.AppConfig{
		Config:      viperConfig,
		DB:          db,
		Log:         log,
		Validator:   validator,
		App:         app,
		RedisClient: redis,
		TokenUtil:   tokenUtil,
	})

	host := viperConfig.GetString("app.host")
	port := viperConfig.GetInt("app.port")

	if err := app.Start(fmt.Sprintf("%s:%d", host, port)); err != nil {
		log.Fatalf("failed to start server: %v", err)
	}
}
