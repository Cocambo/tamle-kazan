package repository

import (
	"context"
	"restaurant-service/internal/models"

	"gorm.io/gorm"
)

type Repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{db: db}
}

// Миграция схемы базы данных
func (r *Repository) Migrate(ctx context.Context) error {
	return r.db.WithContext(ctx).AutoMigrate(
		&models.Restaurant{},
		&models.Review{},
		&models.Favorite{},
		&models.Photo{},
	)
}

// Создаем уникальные индексы, чтобы предотвратить дублирование отзывов и избранного
func (r *Repository) CreateIndexes(ctx context.Context) error {
	if err := r.db.WithContext(ctx).Exec(`
		CREATE UNIQUE INDEX IF NOT EXISTS idx_reviews_user_restaurant 
		ON reviews(user_id, restaurant_id)
	`).Error; err != nil {
		return err
	}

	if err := r.db.WithContext(ctx).Exec(`
		CREATE UNIQUE INDEX IF NOT EXISTS idx_favorites_user_restaurant 
		ON favorites(user_id, restaurant_id)
	`).Error; err != nil {
		return err
	}

	return nil
}
