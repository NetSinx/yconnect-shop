package main

import (
	"fmt"

	"github.com/NetSinx/yconnect-shop/server/user/internal/config"
)

func main() {
	viperConfig := config.NewViper()
	log := config.NewLogger(viperConfig)
	db := config.NewDatabase(viperConfig, log)
	app := config.NewEcho()
	rabbitmq := config.NewRabbitMQ(viperConfig, log)
	defer rabbitmq.Close()

	redis := config.NewRedis(viperConfig, log)
	validator := config.NewValidator()

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
