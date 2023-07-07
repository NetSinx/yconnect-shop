package main

import (
	"github.com/NetSinx/yconnect-shop/category/app/config"
	"github.com/NetSinx/yconnect-shop/category/app/routes"
)

func main() {
	config.ConnectDB()

	server := routes.ApiRoutes()

	server.Logger.Fatal(server.Start(":8080"))
}