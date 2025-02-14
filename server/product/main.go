package main

import (
	"github.com/NetSinx/yconnect-shop/server/product/config"
	"github.com/NetSinx/yconnect-shop/server/product/rabbitmq"
	"github.com/NetSinx/yconnect-shop/server/product/routes"
)

func main() {
	config.ConnectDB()
	
	go rabbitmq.ResponseProductByID()

	server := routes.ApiRoutes()
	server.Logger.Fatal(server.Start(":8081"))
}