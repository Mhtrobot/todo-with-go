package config

import (
	"log"
	"os"
	"todo-app/models"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func InitDB() *gorm.DB {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}

	dsn := os.Getenv("DB_URL")

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),// Disable Logger
	})
	if err != nil {
		log.Fatal("Failed to connect to database")
	}

	db.AutoMigrate(&models.User{}, &models.Todo{})

	return db
}