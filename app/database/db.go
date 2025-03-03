package database

import (
	"fmt"
	"log"

	"jwt_auth/app/config"
	"jwt_auth/app/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect(config *config.Config) {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		config.PostgresHost,
		config.PostgresUser,
		config.PostgresPass,
		config.PostgresDB,
		config.PostgresPort,
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	DB = db
	DB.AutoMigrate(&models.User{}, &models.Session{}, &models.Properties{}, &models.Access{})
}
