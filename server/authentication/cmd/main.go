package main

import (
	"fmt"
	"github.com/NetSinx/yconnect-shop/server/authentication/internal/config"
)

func main() {
	viperConfig := config.NewViper()
	log := config.NewLogger(viperConfig)
	db := config.NewDatabase(viperConfig, log)
	validator := config.NewValidator()
	app := config.NewEcho()
	redis := config.NewRedis(viperConfig, log)
	rabbitmq := config.NewRabbitMQ(viperConfig, log)
	defer rabbitmq.Close()

	config.BootstrapApp(&config.AppConfig{
		Config:      viperConfig,
		DB:          db,
		Log:         log,
		Validator:   validator,
		App:         app,
		RedisClient: redis,
		RabbitMQ:    rabbitmq,
	})

	host := viperConfig.GetString("app.host")
	port := viperConfig.GetInt("app.port")

	if err := app.Start(fmt.Sprintf("%s:%d", host, port)); err != nil {
		log.Fatalf("failed to start server: %v", err)
	}
}
