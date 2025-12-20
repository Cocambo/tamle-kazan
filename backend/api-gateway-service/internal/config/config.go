package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	ServerPort           string
	JwtSecret            string
	UserServiceURL       string
	RestaurantServiceURL string
}

func LoadConfig() (*Config, error) {
	if err := godotenv.Load(); err != nil {
		log.Println(".env not found, using environment variables")
	}

	cfg := &Config{
		ServerPort:           getEnv("SERVER_PORT", "8080"),
		JwtSecret:            getEnv("JWT_SECRET", ""),
		UserServiceURL:       getEnv("USER_SERVICE_URL", "http://localhost:8081"),
		RestaurantServiceURL: getEnv("RESTAURANT_SERVICE_URL", "http://localhost:8082"),
	}

	return cfg, nil
}

func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}
