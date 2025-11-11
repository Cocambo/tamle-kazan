package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	PostgresHost     string
	PostgresPort     string
	PostgresUser     string
	PostgresPassword string
	PostgresDB       string
	ServerPort       string
	JwtSecret        string
}

var AppConfig Config

func LoadConfig() {
	// Загружаем .env
	if err := godotenv.Load(); err != nil {
		log.Println(".env file is not found")
	}

	// if err := godotenv.Load(".env"); err != nil {
	// 	log.Println(".env file is not found")
	// }

	AppConfig = Config{
		PostgresHost:     getEnv("POSTGRES_HOST", "localhost"),
		PostgresPort:     getEnv("POSTGRES_PORT", "5432"),
		PostgresUser:     getEnv("POSTGRES_USER", "postgres"),
		PostgresPassword: getEnv("POSTGRES_PASSWORD", "password"),
		PostgresDB:       getEnv("POSTGRES_DB", "user_service"),
		ServerPort:       getEnv("SERVER_PORT", "8080"),
		JwtSecret:        getEnv("JWT_SECRET", "secret"),
	}
}

func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}
