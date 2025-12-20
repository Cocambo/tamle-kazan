package repository

import (
	"context"
	"errors"
	"restaurant-service/internal/models"

	"gorm.io/gorm"
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

func (r *Repository) GetRestaurants(ctx context.Context, filter RestaurantFilter) ([]models.Restaurant, error) {
	var restaurants []models.Restaurant
	// Пока SQL не выполняется, можно безопасно накапливать условия
	query := r.db.WithContext(ctx).Model(&models.Restaurant{})

	// Применяем фильтры
	// Поиск по имени ресторана
	if filter.Search != "" {
		query = query.Where("name ILIKE ?", "%"+filter.Search+"%")
	}
	// Фильтр по кухне
	if filter.Cuisine != "" {
		query = query.Where("cuisine = ?", filter.Cuisine)
	}
	// Фильтры по среднему чеку
	if filter.MinBill != nil {
		query = query.Where("average_bill >= ?", *filter.MinBill)
	}
	if filter.MaxBill != nil {
		query = query.Where("average_bill <= ?", *filter.MaxBill)
	}
	// Фильтр по минимальному рейтингу
	if filter.MinRating != nil {
		query = query.Where("rating >= ?", *filter.MinRating)
	}
	// Пагинация
	if filter.Limit > 0 {
		query = query.Limit(filter.Limit)
	}
	// Смещение
	if filter.Offset > 0 {
		query = query.Offset(filter.Offset)
	}
	// Выполняем запрос с предзагрузкой фотографий

	// TODO: Подумать над оптимизацией предзагрузки, чтобы загружать только основную фотографию Preload("Photos", "is_main = true")
	err := query.Preload("Photos").Find(&restaurants).Error

	// Возвращаем список ресторанов
	return restaurants, err
}

// Получаем ресторан по его ID(с предзагрузкой фотографий)
func (r *Repository) GetRestaurantByID(ctx context.Context, id uint) (*models.Restaurant, error) {
	var restaurant models.Restaurant
	err := r.db.WithContext(ctx).
		Preload("Photos").
		First(&restaurant, id).Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}
	return &restaurant, err
}

// Обновляем рейтинг и количество отзывов ресторана
func (r *Repository) UpdateRestaurantRating(ctx context.Context, restaurantID uint, rating float64, reviewsCount int) error {
	return r.db.WithContext(ctx).
		Model(&models.Restaurant{}).
		Where("id = ?", restaurantID).
		Updates(map[string]interface{}{
			"rating":        rating,
			"reviews_count": reviewsCount,
		}).Error
}
