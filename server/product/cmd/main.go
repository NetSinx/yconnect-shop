package main

import (
	"fmt"

	"github.com/NetSinx/yconnect-shop/server/product/internal/config"
)

func main() {
	viperConfig := config.NewViper()
	log := config.NewLogger(viperConfig)
	db := config.NewDatabase(viperConfig, log)
	redis := config.NewRedis(viperConfig, log)
	rabbitmq := config.NewRabbitMQ(viperConfig, log)
	validator := config.NewValidator()
	app := config.NewEcho()

	config.NewAppBootstrap(&config.AppBootstrap{
		DB: db,
		App: app,
		Config: viperConfig,
		Log: log,
		RedisClient: redis,
		Validator: validator,
		RabbitMQ: rabbitmq,
	})

	host := viperConfig.GetString("host")
	port := viperConfig.GetInt("port")
	if err := app.Start(fmt.Sprintf("%s:%d", host, port)); err != nil {
		log.Fatalf("failed to start server: %v", err)
	}
}