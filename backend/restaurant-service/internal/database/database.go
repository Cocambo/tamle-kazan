package database

import (
	"context"
	"fmt"
	"log"
	"restaurant-service/internal/config"
	"restaurant-service/internal/repository"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// Connect позволяет подключаться к базе данных
func Connect(cfg *config.Config) (*gorm.DB, error) {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		cfg.DBHost, cfg.DBUser, cfg.DBPassword, cfg.DBName, cfg.DBPort)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %w", err)
	}

	return db, nil
}

// Migrate выполняет миграции базы данных
func Migrate(ctx context.Context, db *gorm.DB) error {
	repo := repository.NewRepository(db)

	if err := repo.Migrate(ctx); err != nil {
		return fmt.Errorf("failed to run migrations: %w", err)
	}

	if err := repo.CreateIndexes(ctx); err != nil {
		log.Printf("Warning: Failed to create indexes: %v", err)
	}

	return nil
}

// ConnectAndMigrate подключается к БД и выполняет миграции
func ConnectAndMigrate(cfg *config.Config) (*gorm.DB, error) {
	db, err := Connect(cfg)
	if err != nil {
		return nil, err
	}

	ctx := context.Background()
	if err := Migrate(ctx, db); err != nil {
		return nil, err
	}

	return db, nil
}
