package config

import (
	"fmt"
	"os"
	"github.com/NetSinx/yconnect-shop/server/product/app/model"
	"github.com/NetSinx/yconnect-shop/server/product/utils"
	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() {
	var products model.Product
	var images model.Image

	godotenv.Load()

	initDb := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
													os.Getenv("DB_USER"),
													os.Getenv("DB_PASS"),
													os.Getenv("DB_HOST"),
													os.Getenv("DB_PORT"),
													os.Getenv("DB_DBNAME"),
	)
	
	db, err := gorm.Open(mysql.Open(initDb), &gorm.Config{})
	if err != nil {
		utils.PanicError(err)
	}

	db.AutoMigrate(&products, &images)

	DB = db
}