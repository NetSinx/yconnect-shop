package config

import (
	"fmt"
	"os"
	"github.com/NetSinx/yconnect-shop/server/seller/model/entity"
	"github.com/NetSinx/yconnect-shop/server/seller/utils"
	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConfigDB() {
	var seller entity.Seller

	godotenv.Load()

	loadEnv := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
													os.Getenv("DB_USER"),
													os.Getenv("DB_PASS"),
													os.Getenv("DB_HOST"),
													os.Getenv("DB_PORT"),
													os.Getenv("DB_DBNAME"),
												)

	initDb, err := gorm.Open(mysql.Open(loadEnv), &gorm.Config{})
	if err != nil {
		utils.LogPanic(err)
	}

	initDb.AutoMigrate(&seller)
	DB = initDb
}