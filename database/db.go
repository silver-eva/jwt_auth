package database

import (
	"fmt"
	"log"

	"jwt_auth/config"
	"jwt_auth/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

var DB *gorm.DB

func Connect(config *config.Config) {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable",
		config.PostgresHost,
		config.PostgresUser,
		config.PostgresPass,
		config.PostgresDB,
		config.PostgresPort,
	)

	db, err := gorm.Open(
		postgres.New(
			postgres.Config{
				DSN: dsn,
				PreferSimpleProtocol: true,
			},
		), 
		&gorm.Config{
			NamingStrategy: schema.NamingStrategy{
				TablePrefix: "auth.",
				SingularTable: false,
			},
		},
	)
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	DB = db
	DB.AutoMigrate(&models.User{}, &models.Session{}, &models.Properties{}, &models.Access{})
}
