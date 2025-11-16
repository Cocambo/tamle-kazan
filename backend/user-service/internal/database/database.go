package database

import (
	"fmt"
	"log"

	"github.com/Cocambo/tamle-kazan/backend/user-service/internal/config"
	"github.com/Cocambo/tamle-kazan/backend/user-service/internal/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// Хранит подключение к базе данных
var DB *gorm.DB

// InitDB — подключение к PostgreSQL через GORM
func InitDB() {
	// Формирум DSN (Data Source Name), строку для подключения к БД
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		config.AppConfig.PostgresHost,
		config.AppConfig.PostgresUser,
		config.AppConfig.PostgresPassword,
		config.AppConfig.PostgresDB,
		config.AppConfig.PostgresPort,
	)

	//Открываем соединение с БД
	var err error
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	log.Println("Connect to database succesful")
}

// Migrate выполняет миграцию моделей
func Migrate() {
	err := DB.AutoMigrate(&models.User{})
	if err != nil {
		log.Fatalf("Migration failed %v", err)
	}

	log.Println("Migrations succesful")
}
