package main

import (
	"fmt"
	"github.com/NetSinx/yconnect-shop/server/category/internal/config"
)

func main() {
	viperConfig := config.NewViper()
	log := config.NewLogger(viperConfig)
	db := config.NewDatabase(viperConfig, log)
	validator := config.NewValidator()
	rabbitmq := config.NewRabbitMQ(viperConfig, log)
	app := config.NewEcho()

	config.NewAppBootstrap(&config.AppBootstrap{
		DB:        db,
		App:       app,
		Config:    viperConfig,
		Log:       log,
		Validator: validator,
		RabbitMQ:  rabbitmq,
	})

	host := viperConfig.GetString("app.host")
	port := viperConfig.GetInt("app.port")

	if err := app.Start(fmt.Sprintf("%s:%d", host, port)); err != nil {
		log.Fatalf("failed to start server: %v", err)
	}
}
