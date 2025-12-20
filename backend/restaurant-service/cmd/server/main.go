package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"restaurant-service/internal/config"
	"restaurant-service/internal/database"
	"restaurant-service/internal/handlers"
	"restaurant-service/internal/repository"
	"restaurant-service/internal/router"
	"restaurant-service/internal/seed"
	"restaurant-service/internal/service"
)

func main() {
	cfg := config.Load()

	// Подключение к БД и выполнение миграций
	db, err := database.ConnectAndMigrate(cfg)
	if err != nil {
		log.Fatalf("Failed to connect and migrate database: %v", err)
	}

	// Создание директории для фотографий
	if err := os.MkdirAll(cfg.PhotosDir, 0755); err != nil {
		log.Fatalf("Failed to create photos directory: %v", err)
	}

	// Инициализация репозитория
	repo := repository.NewRepository(db)

	// Загрузка seed данных
	ctx := context.Background()
	if err := seed.LoadSeedData(ctx, db, cfg.PhotosDir); err != nil {
		log.Printf("Warning: Failed to load seed data: %v", err)
	}

	// Инициализация сервиса
	svc := service.NewService(repo)

	// Инициализация handlers
	h := handlers.NewHandlers(svc)

	// Настройка роутера
	r := router.SetupRouter(h)

	// Запуск сервера
	port := fmt.Sprintf(":%d", cfg.GetServerPortInt())
	log.Printf("Server starting on port %s", port)
	if err := r.Run(port); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
