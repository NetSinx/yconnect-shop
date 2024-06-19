package main

import (
	"github.com/NetSinx/yconnect-shop/server/category/config"
	"github.com/NetSinx/yconnect-shop/server/category/routes"
)

func main() {
	config.ConnectDB()

	server := routes.ApiRoutes()

	server.Logger.Fatal(server.Start(":8080"))
}