package config

import (
	"fmt"
	"time"
	"github.com/NetSinx/yconnect-shop/server/authentication/internal/entity"
	"github.com/NetSinx/yconnect-shop/server/authentication/internal/helpers"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func NewDatabase(viper *viper.Viper, log *logrus.Logger) *gorm.DB {
	username := viper.GetString("database.username")
	password := viper.GetString("database.password")
	host := viper.GetString("database.host")
	port := viper.GetInt("database.port")
	name := viper.GetString("database.name")
	idleConnection := viper.GetInt("database.pool.idle")
	maxConnection := viper.GetInt("database.pool.max")
	lifetimeConnection := viper.GetInt("database.pool.lifetime")

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local", username, password, host, port, name)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		SkipDefaultTransaction: true,
		PrepareStmt:            true,
		Logger: logger.New(&logrusWritter{Logger: log}, logger.Config{
			SlowThreshold:             (3 * time.Second),
			LogLevel:                  logger.Info,
			IgnoreRecordNotFoundError: true,
			ParameterizedQueries:      true,
			Colorful:                  false,
		}),
	})
	helpers.FatalError(log, err, "failed to initialize database")

	connection, err := db.DB()
	helpers.FatalError(log, err, "failed to connect database")

	connection.SetMaxIdleConns(idleConnection)
	connection.SetMaxOpenConns(maxConnection)
	connection.SetConnMaxLifetime(time.Duration(lifetimeConnection) * time.Second)

	db.AutoMigrate(&entity.Authentication{})

	return db
}

type logrusWritter struct {
	Logger *logrus.Logger
}

func (l *logrusWritter) Printf(message string, args ...any) {
	l.Logger.Tracef(message, args...)
}