package service

import (
	"context"
	"errors"
	"restaurant-service/internal/models"
)

// CreateReview создает новый отзыв и обновляет рейтинг ресторана
func (s *Service) CreateReview(ctx context.Context, review *models.Review) error {
	// Проверка рейтинга
	if review.Rating < 1 || review.Rating > 5 {
		return errors.New("rating must be between 1 and 5")
	}

	// Проверка, существует ли ресторан
	restaurant, err := s.repo.GetRestaurantByID(ctx, review.RestaurantID)
	// Возможна техническая ошибка при получении ресторана, проверяем ее сначала
	if err != nil {
		return err
	}
	if restaurant == nil {
		return errors.New("restaurant not found")
	}

	// Проверка, есть ли уже отзыв от этого пользователя
	existing, err := s.repo.GetReviewByUserAndRestaurant(ctx, review.UserID, review.RestaurantID)
	if err != nil {
		return err
	}
	if existing != nil {
		return errors.New("user already reviewed this restaurant")
	}

	// Создание отзыва
	if err := s.repo.CreateReview(ctx, review); err != nil {
		return err
	}

	// Пересчет рейтинга ресторана
	rating, count, err := s.repo.CalculateRestaurantRating(ctx, review.RestaurantID)
	if err != nil {
		return err
	}

	// Обновление рейтинга ресторана
	return s.repo.UpdateRestaurantRating(ctx, review.RestaurantID, rating, count)
}

// GetReviewsByRestaurantID получает все отзывы для заданного ресторана, просто прокси вызов репозитория
func (s *Service) GetReviewsByRestaurantID(ctx context.Context, restaurantID uint) ([]models.Review, error) {
	return s.repo.GetReviewsByRestaurantID(ctx, restaurantID)
}
