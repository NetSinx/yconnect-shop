package main

import (
	"github.com/NetSinx/yconnect-shop/server/seller/config"
	"github.com/NetSinx/yconnect-shop/server/seller/routes"
)

func main() {
	config.ConfigDB()

	server := routes.APIRoutes()
	server.Logger.Fatal(server.Start(":8084"))
}