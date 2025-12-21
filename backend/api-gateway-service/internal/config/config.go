package config

import (
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
)

type Config struct {
	ServerPort           string
	JwtSecret            string
	UserServiceURL       string
	RestaurantServiceURL string
	CORS                 CORSConfig
}

type CORSConfig struct {
	AllowOrigins     []string
	AllowMethods     []string
	AllowHeaders     []string
	ExposeHeaders    []string
	AllowCredentials bool
	MaxAge           time.Duration
}

func LoadConfig() (*Config, error) {
	if err := godotenv.Load(); err != nil {
		log.Println(".env not found, using environment variables")
	}

	cfg := &Config{
		ServerPort:           getEnv("SERVER_PORT", "8080"),
		JwtSecret:            getEnv("JWT_SECRET", ""),
		UserServiceURL:       getEnv("USER_SERVICE_URL", "http://localhost:8081"),
		RestaurantServiceURL: getEnv("ORDER_SERVICE_URL", "http://localhost:8082"),
		CORS: CORSConfig{
			AllowOrigins:     []string{"http://localhost:5173", "http://localhost:8081", "http://localhost:8082"},
			AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
			AllowHeaders:     []string{"Origin", "Content-Type", "Accept", "Authorization", "Cookie"},
			ExposeHeaders:    []string{"Content-Length", "Set-Cookie"},
			AllowCredentials: true,
			MaxAge:           12 * time.Hour,
		},
	}

	return cfg, nil
}

func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}
