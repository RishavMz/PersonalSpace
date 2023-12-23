package config

import (
	"TodoService/models"
	"log"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var conn *gorm.DB

func CreateConnection() *gorm.DB {
    database := os.Getenv("DATABASE_URL")
	conn, err := gorm.Open(mysql.Open(database), &gorm.Config{})
    if err != nil {
        log.Fatal("[SERVER]: Failed to connect to the database")
    } else {
		log.Println("[SERVER]: Successfully connected to database")
	}

	conn.AutoMigrate(&models.User{})
	conn.AutoMigrate(&models.Todo{})

	return conn
}

