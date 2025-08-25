package main

import (
	"github.com/NetSinx/yconnect-shop/server/product/db"
	"github.com/NetSinx/yconnect-shop/server/product/handler/http"
	"github.com/NetSinx/yconnect-shop/server/product/repository"
	"github.com/NetSinx/yconnect-shop/server/product/service"
	"github.com/NetSinx/yconnect-shop/server/product/service/rabbitmq"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
)

func main() {
	godotenv.Load()

	db := db.ConnectDB()
	productRepository := repository.NewProductRepository(db)
	productService := service.NewProductService(productRepository)
	productHandler := http.NewProductHandler(productService)
	
	go rabbitmq.ConsumeCategoryEvents()

	e := echo.New()
	http.ApiRoutes(e, productHandler)

	e.Logger.Fatal(e.Start(":8081"))
}