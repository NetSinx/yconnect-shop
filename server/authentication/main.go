package main

import (
	"github.com/NetSinx/yconnect-shop/server/authentication/handler/http"
	"github.com/NetSinx/yconnect-shop/server/authentication/service"
	"github.com/labstack/echo/v4"
)

func main() {
	authService := service.AuthServ()
	authHandler := http.AuthHandler(authService)

	e := echo.New()
	http.APIRoutes(e, authHandler)

	e.Logger.Fatal(e.Start(":8086"))
}