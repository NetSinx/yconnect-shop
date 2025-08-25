package db

import (
	"fmt"
	"os"
	"github.com/NetSinx/yconnect-shop/server/category/errs"
	"github.com/NetSinx/yconnect-shop/server/category/model"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func ConnectDB() *gorm.DB {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
						os.Getenv("DB_USER"),
						os.Getenv("DB_PASS"),
						os.Getenv("DB_HOST"),
						os.Getenv("DB_PORT"),
						os.Getenv("DB_NAME"),
	)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		SkipDefaultTransaction: true,
		PrepareStmt: true,
	})
	errs.PanicError(err)

	err = db.AutoMigrate(&model.Category{})
	errs.PanicError(err)

	return db
}