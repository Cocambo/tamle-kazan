package service

import (
	"context"
	"restaurant-service/internal/models"
	"restaurant-service/internal/repository"
)

type RestaurantFilter struct {
	Search    string
	Cuisine   string
	MinBill   *float64
	MaxBill   *float64
	MinRating *float64
	Limit     int
	Offset    int
}

// GetRestaurants передает вызов репозиторию для получения списка ресторанов с учетом фильтров
func (s *Service) GetRestaurants(ctx context.Context, filter RestaurantFilter) ([]models.Restaurant, error) {
	repoFilter := repository.RestaurantFilter{
		Search:    filter.Search,
		Cuisine:   filter.Cuisine,
		MinBill:   filter.MinBill,
		MaxBill:   filter.MaxBill,
		MinRating: filter.MinRating,
		Limit:     filter.Limit,
		Offset:    filter.Offset,
	}
	return s.repo.GetRestaurants(ctx, repoFilter)
}

// GetRestaurantByID передает вызов репозиторию для получения ресторана по его ID
func (s *Service) GetRestaurantByID(ctx context.Context, id uint) (*models.Restaurant, error) {
	return s.repo.GetRestaurantByID(ctx, id)
}

// GetTopRestaurants передает вызов репозиторию для получения топ ресторанов по рейтингу
func (s *Service) GetTopRestaurants(ctx context.Context, limit int) ([]models.Restaurant, error) {
	return s.repo.GetTopRestaurants(ctx, limit)
}
