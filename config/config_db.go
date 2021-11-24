package config

import (
	"fmt"
	"os"

	"github.com/helyus1412/messaging-api/models"
	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	Database *gorm.DB
)

func Init_DB() *gorm.DB {
	err := godotenv.Load()

	if err != nil {
		fmt.Println(err.Error())
	}

	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASS")
	dbHost := os.Getenv("DB_HOST")
	dbName := os.Getenv("DB_NAME")
	dbPort := os.Getenv("DB_PORT")
	dbMigrate := os.Getenv("DB_AUTO_MIGRATE")

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", dbUser, dbPass, dbHost, dbPort, dbName)

	if dbMigrate == "true" {
		db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

		if err != nil {
			fmt.Println(err)
		}

		Database = db

		db.AutoMigrate(&models.User{})
		db.AutoMigrate(&models.Messages{})
	}

	return Database
}
