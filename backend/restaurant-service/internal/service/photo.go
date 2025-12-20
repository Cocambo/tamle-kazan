package service

import (
	"context"
	"restaurant-service/internal/models"
)

// CreatePhoto создает новое фото для ресторана
func (s *Service) CreatePhoto(ctx context.Context, photo *models.Photo) error {
	// Если это главное фото, снимаем флаг is_main с других фото
	if photo.IsMain {
		if err := s.repo.UnsetMainPhotos(ctx, photo.RestaurantID); err != nil {
			return err
		}
	}

	return s.repo.CreatePhoto(ctx, photo)
}

// GetPhotosByRestaurantID получает все фото для заданного ресторана
func (s *Service) GetPhotosByRestaurantID(ctx context.Context, restaurantID uint) ([]models.Photo, error) {
	return s.repo.GetPhotosByRestaurantID(ctx, restaurantID)
}
