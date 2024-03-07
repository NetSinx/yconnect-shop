package main

import (
	"github.com/NetSinx/yconnect-shop/server/product/app/config"
	"github.com/NetSinx/yconnect-shop/server/product/app/routes"
)

func main() {
	config.ConnectDB()

	server := routes.ApiRoutes()
	server.Logger.Fatal(server.Start(":8081"))
}