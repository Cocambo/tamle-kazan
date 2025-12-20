package config

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type Config struct {
	DBHost     string
	DBPort     string
	DBUser     string
	DBPassword string
	DBName     string
	ServerPort string
	PhotosDir  string
}

func Load() *Config {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, using environment variables")
	}

	return &Config{
		DBHost:     getEnv("DB_HOST", "localhost"),
		DBPort:     getEnv("DB_PORT", "5432"),
		DBUser:     getEnv("DB_USER", "postgres"),
		DBPassword: getEnv("DB_PASSWORD", "postgres"),
		DBName:     getEnv("DB_NAME", "restaurant_db"),
		ServerPort: getEnv("SERVER_PORT", "8082"),
		PhotosDir:  getEnv("PHOTOS_DIR", "./storage/photos"),
	}
}

func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}

func (c *Config) GetDBPortInt() int {
	port, err := strconv.Atoi(c.DBPort)
	if err != nil {
		return 5432
	}
	return port
}

func (c *Config) GetServerPortInt() int {
	port, err := strconv.Atoi(c.ServerPort)
	if err != nil {
		return 8082
	}
	return port
}
