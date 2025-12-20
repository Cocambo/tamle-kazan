package service

import (
	"context"
	"errors"
	"restaurant-service/internal/models"
)

// AddFavorite добавляет ресторан в избранное пользователя
func (s *Service) AddFavorite(ctx context.Context, userID, restaurantID uint) error {
	// Проверка, существует ли ресторан
	restaurant, err := s.repo.GetRestaurantByID(ctx, restaurantID)
	if err != nil {
		return err
	}
	if restaurant == nil {
		return errors.New("restaurant not found")
	}

	// Проверка, не добавлен ли уже в избранное
	existing, err := s.repo.GetFavorite(ctx, userID, restaurantID)
	if err != nil {
		return err
	}
	if existing != nil {
		return errors.New("restaurant already in favorites")
	}

	favorite := &models.Favorite{
		UserID:       userID,
		RestaurantID: restaurantID,
	}

	return s.repo.AddFavorite(ctx, favorite)
}

// RemoveFavorite удаляет ресторан из избранного пользователя
func (s *Service) RemoveFavorite(ctx context.Context, userID, restaurantID uint) error {
	return s.repo.RemoveFavorite(ctx, userID, restaurantID)
}

// GetUserFavorites получает все избранные рестораны пользователя
func (s *Service) GetUserFavorites(ctx context.Context, userID uint) ([]models.Restaurant, error) {
	return s.repo.GetUserFavorites(ctx, userID)
}