package main

import (
	"github.com/NetSinx/yconnect-shop/server/cart/config"
	"github.com/NetSinx/yconnect-shop/server/cart/routes"
)

func main() {
	config.DBConfig()

	server := routes.ApiRoutes()
	server.Logger.Fatal(server.Start(":8083"))
}