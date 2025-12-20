package repository

import (
	"context"
	"errors"
	"restaurant-service/internal/models"

	"gorm.io/gorm"
)

// Добавляем ресторан в избранное
func (r *Repository) AddFavorite(ctx context.Context, favorite *models.Favorite) error {
	return r.db.WithContext(ctx).Create(favorite).Error
}

// Удаляем ресторан из избранного пользователя по userID и restaurantID
func (r *Repository) RemoveFavorite(ctx context.Context, userID, restaurantID uint) error {
	result := r.db.WithContext(ctx).
		Where("user_id = ? AND restaurant_id = ?", userID, restaurantID).
		Delete(&models.Favorite{})

	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return errors.New("favorite not found")
	}

	return nil
}

// Получаем избранный ресторан по userID и restaurantID
func (r *Repository) GetFavorite(ctx context.Context, userID, restaurantID uint) (*models.Favorite, error) {
	var favorite models.Favorite
	err := r.db.WithContext(ctx).
		Where("user_id = ? AND restaurant_id = ?", userID, restaurantID).
		First(&favorite).Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}
	return &favorite, err
}

// Получаем все избранные рестораны пользователя
func (r *Repository) GetUserFavorites(ctx context.Context, userID uint) ([]models.Restaurant, error) {
	var restaurants []models.Restaurant
	err := r.db.WithContext(ctx).
		Table("restaurants").
		Joins("INNER JOIN favorites ON restaurants.id = favorites.restaurant_id").
		Where("favorites.user_id = ?", userID).
		Preload("Photos").
		Find(&restaurants).Error

	return restaurants, err
}
