package repository

import (
	"context"
	"errors"
	"restaurant-service/internal/models"

	"gorm.io/gorm"
)

// Создаем отзыв для ресторана
func (r *Repository) CreateReview(ctx context.Context, review *models.Review) error {
	return r.db.WithContext(ctx).Create(review).Error
}

// Получаем все отзывы для конкретного ресторана по его ID, отсортированные по дате создания (новые первыми)
func (r *Repository) GetReviewsByRestaurantID(ctx context.Context, restaurantID uint) ([]models.Review, error) {
	var reviews []models.Review
	err := r.db.WithContext(ctx).
		Preload("Restaurant").
		Where("restaurant_id = ?", restaurantID).
		Order("created_at DESC").
		Find(&reviews).Error
	return reviews, err
}

// Получаем отзыв пользователя для конкретного ресторана по userID и restaurantID
// Если отзыв не найден, возвращаем nil без ошибки
// Это позволяет легко проверить, оставил ли пользователь отзыв
func (r *Repository) GetReviewByUserAndRestaurant(ctx context.Context, userID, restaurantID uint) (*models.Review, error) {
	var review models.Review
	err := r.db.WithContext(ctx).
		Where("user_id = ? AND restaurant_id = ?", userID, restaurantID).
		First(&review).Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}
	return &review, err
}

// Вычисляем средний рейтинг и количество отзывов для ресторана
func (r *Repository) CalculateRestaurantRating(ctx context.Context, restaurantID uint) (float64, int, error) {
	var result struct {
		AverageRating float64
		Count         int
	}
	// Используем COALESCE, чтобы вернуть 0, если нет отзывов
	err := r.db.WithContext(ctx).
		Model(&models.Review{}).
		Select("COALESCE(AVG(rating), 0) as average_rating, COUNT(*) as count").
		Where("restaurant_id = ?", restaurantID).
		Scan(&result).Error

	return result.AverageRating, result.Count, err
}
