package config

import (
	"fmt"
	"os"
	"github.com/NetSinx/yconnect-shop/server/user/model/entity"
	"github.com/NetSinx/yconnect-shop/server/user/utils"
	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() {
	var users entity.User
	var alamat entity.Alamat

	godotenv.Load()

	initDb := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
													os.Getenv("DB_USER"),
													os.Getenv("DB_PASS"),
													os.Getenv("DB_HOST"),
													os.Getenv("DB_PORT"),
													os.Getenv("DB_NAME"),
												)

	db, err := gorm.Open(mysql.Open(initDb), &gorm.Config{})
	if err != nil {
		utils.LogPanic(err)
	}

	db.AutoMigrate(&users, &alamat)
	DB = db
}