package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	ServerPort      string
	JwtSecret       string
	UserServiceURL  string
	OrderServiceURL string
}

func LoadConfig() (*Config, error) {
	if err := godotenv.Load(); err != nil {
		log.Println(".env not found, using environment variables")
	}

	cfg := &Config{
		ServerPort:      getEnv("SERVER_PORT", "8080"),
		JwtSecret:       getEnv("JWT_SECRET", "d8f3b4e7a01c2f8f64ad3e7b9c61d39a5f27cf8f9272b9a18c3a77d4a2b3e9cf"),
		UserServiceURL:  getEnv("USER_SERVICE_URL", "http://localhost:8081"),
		OrderServiceURL: getEnv("ORDER_SERVICE_URL", "http://localhost:8082"),
	}

	return cfg, nil
}

func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}
