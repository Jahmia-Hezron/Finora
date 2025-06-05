package database

import (
	"finora/database/model"
	"fmt"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB
var err error

func DBConnect() {
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_PORT"),
	)

	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("❌ Failed to connect to the database: %v\n", err)
		os.Exit(1)
	} else {
		log.Println("✅ Database connection established successfully")
	}
}

func DBMigrate() {
	if err := DB.AutoMigrate(
		&model.Expense{},
		&model.Income{},
		&model.Budget{},
		&model.ExpenseCategory{},
		&model.IncomeSource{},
		&model.User{},
		); err != nil {
		log.Fatalf("❌ Failed to migrate database models: %v\n", err)
		os.Exit(1)
	} else {
		log.Println("✅ Database models migrated successfully")
	}
}