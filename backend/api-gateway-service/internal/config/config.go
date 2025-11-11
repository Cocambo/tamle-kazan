package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	ServerPort     string
	UserServiceURL string
}

var AppConfig Config

func LoadConfig() {
	// Загружаем .env если есть
	if err := godotenv.Load(); err != nil {
		log.Println(".env file not found, using system environment variables")
	}

	// if err := godotenv.Load(".env"); err != nil {
	// 	log.Println(".env file not found, using system environment variables")
	// }

	AppConfig = Config{
		ServerPort:     getEnv("SERVER_PORT", "8081"),
		UserServiceURL: getEnv("USER_SERVICE_URL", "http://user-service:8080"),
	}

	log.Printf("Loaded config: port=%s, user-service=%s", AppConfig.ServerPort, AppConfig.UserServiceURL)
}

func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}
