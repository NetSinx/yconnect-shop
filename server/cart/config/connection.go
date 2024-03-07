package config

import (
	"fmt"
	"log"
	"os"
	"github.com/NetSinx/yconnect-shop/server/cart/model"
	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func DBConfig() {
	var cart model.Cart

	if err := godotenv.Load(); err != nil {
		initDb := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", 
			os.Getenv("DB_USER"),
			os.Getenv("DB_PASSWORD"),
			os.Getenv("DB_HOST"),
			os.Getenv("DB_PORT"),
			os.Getenv("DB_DATABASE"),
		)

		db, err := gorm.Open(mysql.Open(initDb), &gorm.Config{})
		if err != nil {
			log.Fatal(err)
		}

		db.AutoMigrate(&cart)

		DB = db

		return
	}

	initDb := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", 
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_DATABASE"),
	)

	db, err := gorm.Open(mysql.Open(initDb), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	db.AutoMigrate(&cart)

	DB = db
}