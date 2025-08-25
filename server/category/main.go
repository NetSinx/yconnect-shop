package main

import (
	"github.com/NetSinx/yconnect-shop/server/category/db"
	"github.com/NetSinx/yconnect-shop/server/category/handler/http"
	"github.com/NetSinx/yconnect-shop/server/category/repository"
	"github.com/NetSinx/yconnect-shop/server/category/service"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
)

func main() {
	godotenv.Load()

	db := db.ConnectDB()
	categoryRepository := repository.CategoryRepository(db)
	categoryService := service.CategoryService(categoryRepository)
	categoryHandler := http.CategoryHandler(categoryService)
	
	e := echo.New()
	http.ApiRoutes(e, categoryHandler)
	
	e.Logger.Fatal(e.Start(":8080"))
}