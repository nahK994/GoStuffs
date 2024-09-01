package config

import (
	"book-store/pkg/models"
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() {
	POSTGRES_USER := "user"
	POSTGRES_PASSWORD := "password"
	POSTGRES_DB := "book_store"
	POSTGRES_PORT := 5432
	POSTGRES_HOST := "localhost"
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable TimeZone=Asia/Dhaka", POSTGRES_HOST, POSTGRES_USER, POSTGRES_PASSWORD, POSTGRES_DB, POSTGRES_PORT)

	DB, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	DB.AutoMigrate(&models.Book{})
}
