package main

import (
	"github.com/NetSinx/yconnect-shop/server/order/config"
	"github.com/NetSinx/yconnect-shop/server/order/routes"
)

func main() {
	config.ConnectDB()

	server := routes.RoutesAPI()
	server.Logger.Fatal(server.Start(":8084"))
}