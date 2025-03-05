package database

import (
	"fmt"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DATABASE *gorm.DB

func initDB() error {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_DATABASE"),
		os.Getenv("DB_PORT"),
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		return err
	}

	DATABASE = db
	return nil
}

func Connect() {
	err := initDB()
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Database connected")
}

func Close() {
	if DATABASE == nil {
		return
	}

	sqlDB, err := DATABASE.DB()
	if err != nil {
		fmt.Printf("Error getting sql.DB: %v", err)
		return
	}

	log.Println("Database connection closed")
	sqlDB.Close()
}
