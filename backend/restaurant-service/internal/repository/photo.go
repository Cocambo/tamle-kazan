package repository

import (
	"context"
	"restaurant-service/internal/models"
)

// Создаем фотографию для ресторана
func (r *Repository) CreatePhoto(ctx context.Context, photo *models.Photo) error {
	return r.db.WithContext(ctx).Create(photo).Error
}

// Получаем все фотографии для конкретного ресторана по его ID, отсортированные так,
// что основная фотография идет первой, затем по дате создания
func (r *Repository) GetPhotosByRestaurantID(ctx context.Context, restaurantID uint) ([]models.Photo, error) {
	var photos []models.Photo
	err := r.db.WithContext(ctx).
		Where("restaurant_id = ?", restaurantID).
		Order("is_main DESC, created_at ASC").
		Find(&photos).Error
	return photos, err
}

// Снимаем отметку основной фотографии для всех фотографий ресторана
func (r *Repository) UnsetMainPhotos(ctx context.Context, restaurantID uint) error {
	return r.db.WithContext(ctx).
		Model(&models.Photo{}).
		Where("restaurant_id = ?", restaurantID).
		Update("is_main", false).Error
}
