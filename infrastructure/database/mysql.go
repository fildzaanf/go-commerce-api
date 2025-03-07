package database

import (
	"fmt"
	"go-commerce-api/infrastructure/config"
	"time"

	"github.com/sirupsen/logrus"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func ConnectMySQL() *gorm.DB {
	config, err := config.LoadConfig()
	if err != nil {
		logrus.WithError(err).Fatal("failed to load MySQL configuration")
	}

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local&allowPublicKeyRetrieval=true",
		config.MYSQL.MYSQL_USER,
		config.MYSQL.MYSQL_PASS,
		config.MYSQL.MYSQL_HOST,
		config.MYSQL.MYSQL_PORT,
		config.MYSQL.MYSQL_NAME,
	)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info), 
	})
	if err != nil {
		logrus.WithError(err).Fatal("failed to connect to MySQL")
	}

	sqlDB, err := db.DB()
	if err != nil {
		logrus.WithError(err).Fatal("failed to get database instance")
	}

	sqlDB.SetMaxOpenConns(10)                  
	sqlDB.SetMaxIdleConns(5)                   
	sqlDB.SetConnMaxLifetime(30 * time.Minute) 

	Migration(db)

	logrus.Info("Connected to MySQL successfully")

	return db
}
