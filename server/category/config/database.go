package config

import (
	"fmt"
	"log"
	"os"
	"github.com/NetSinx/yconnect-shop/server/category/app/model"
	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() {
	var category model.Category

	godotenv.Load()

	initDB := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
												 os.Getenv("DB_USER"),
												 os.Getenv("DB_PASS"),
												 os.Getenv("DB_HOST"),
												 os.Getenv("DB_PORT"),
												 os.Getenv("DB_DBNAME"))

	db, err := gorm.Open(mysql.Open(initDB), &gorm.Config{})

	if err != nil {
		log.Panic(err)
	}

	db.AutoMigrate(&category)

	DB = db
}