package main

import (
	"fmt"
	"os"
	"github.com/NetSinx/yconnect-shop/server/product/errs"
	"github.com/NetSinx/yconnect-shop/server/product/handler/http"
	"github.com/NetSinx/yconnect-shop/server/product/model"
	"github.com/NetSinx/yconnect-shop/server/product/repository"
	"github.com/NetSinx/yconnect-shop/server/product/service"
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
							os.Getenv("DB_NAME"),
				)
	
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		errs.PanicError(err)
	}

	db.AutoMigrate(&model.Product{}, &model.Gambar{})

	productRepository := repository.ProductRepository(db)
	productService := service.ProductService(productRepository)
	productHandler := http.ProductHandler(productService)
	
	e := echo.New()
	http.ApiRoutes(e, productHandler)

	e.Logger.Fatal(e.Start(":8081"))
}