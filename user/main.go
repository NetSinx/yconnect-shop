package main

import (
	"github.com/NetSinx/yconnect-shop/user/app/config"
	"github.com/NetSinx/yconnect-shop/user/app/routes"
)

func main() {
	config.ConnectDB()

	server := routes.ApiRoutes()
	server.Logger.Fatal(server.Start(":8082"))
}