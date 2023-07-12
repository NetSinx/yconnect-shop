package main

import (
	"github.com/NetSinx/yconnect-shop/cart/config"
	"github.com/NetSinx/yconnect-shop/cart/routes"
)

func main() {
	config.DBConfig()

	server := routes.ApiRoutes()
	server.Logger.Fatal(server.Start(":8083"))
}