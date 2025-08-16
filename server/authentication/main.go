package main

import "github.com/NetSinx/yconnect-shop/server/authentication/routes"

func main() {
	e := routes.APIRoutes()
	e.Logger.Fatal(e.Start(":8086"))
}