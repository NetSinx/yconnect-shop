package main

import (
	"fmt"
	"os"
	"github.com/NetSinx/yconnect-shop/server/category/errs"
	"github.com/NetSinx/yconnect-shop/server/category/handler/http"
	"github.com/NetSinx/yconnect-shop/server/category/model"
	"github.com/NetSinx/yconnect-shop/server/category/repository"
	"github.com/NetSinx/yconnect-shop/server/category/service"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	godotenv.Load()

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
						os.Getenv("DB_USER"),
						os.Getenv("DB_PASS"),
						os.Getenv("DB_HOST"),
						os.Getenv("DB_PORT"),
						os.Getenv("DB_DBNAME"),
	)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	errs.PanicError(err)

	db.AutoMigrate(&model.Category{})

	categoryRepository := repository.CategoryRepository(db)
	categoryService := service.CategoryService(categoryRepository)
	categoryHandler := http.CategoryHandler(categoryService)
	
	e := echo.New()
	http.ApiRoutes(e, categoryHandler)
	
	e.Logger.Fatal(e.Start(":8080"))
}