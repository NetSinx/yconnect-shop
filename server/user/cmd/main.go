package main

import (
	"github.com/NetSinx/yconnect-shop/server/user/config"
	// "github.com/NetSinx/yconnect-shop/server/user/rabbitmq"
	"github.com/NetSinx/yconnect-shop/server/user/routes"
)

func main() {
	config.ConnectDB()

	// go rabbitmq.ResponseGetUsernameByID()

	server := routes.APIRoutes()
	server.Logger.Fatal(server.Start(":8082"))
}