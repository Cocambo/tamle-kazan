package repository

import (
	"context"
	"restaurant-service/internal/models"
)

func (r *Repository) GetUserFavoriteRestaurants(ctx context.Context, userID uint) ([]models.Restaurant, error) {
	var restaurants []models.Restaurant

	err := r.db.WithContext(ctx).
		Model(&models.Restaurant{}).
		Joins("JOIN favorites ON favorites.restaurant_id = restaurants.id").
		Where("favorites.user_id = ?", userID).
		Preload("Photos", "is_main = true").
		Find(&restaurants).Error

	return restaurants, err
}

func (r *Repository) GetUserHighlyRatedRestaurants(ctx context.Context, userID uint, minRating int) ([]models.Restaurant, error) {
	var restaurants []models.Restaurant

	err := r.db.WithContext(ctx).
		Model(&models.Restaurant{}).
		Joins("JOIN reviews ON reviews.restaurant_id = restaurants.id").
		Where("reviews.user_id = ? AND reviews.rating >= ?", userID, minRating).
		Preload("Photos", "is_main = true").
		Find(&restaurants).Error

	return restaurants, err
}

func (r *Repository) GetCandidateRestaurantsByCuisines(ctx context.Context, cuisines []string, excludeIDs []uint, limit int) ([]models.Restaurant, error) {
	var restaurants []models.Restaurant

	query := r.db.WithContext(ctx).Model(&models.Restaurant{})

	if len(cuisines) > 0 {
		query = query.Where("cuisine IN ?", cuisines)
	}

	if len(excludeIDs) > 0 {
		query = query.Where("restaurants.id NOT IN ?", excludeIDs)
	}

	err := query.
		Order("rating DESC").
		Order("reviews_count DESC").
		Limit(limit).
		Preload("Photos", "is_main = true").
		Find(&restaurants).Error

	return restaurants, err
}

func (r *Repository) GetPopularRestaurantsExcluding(ctx context.Context, excludeIDs []uint, limit int) ([]models.Restaurant, error) {
	var restaurants []models.Restaurant

	query := r.db.WithContext(ctx).Model(&models.Restaurant{})

	if len(excludeIDs) > 0 {
		query = query.Where("restaurants.id NOT IN ?", excludeIDs)
	}

	err := query.
		Order("rating DESC").
		Order("reviews_count DESC").
		Limit(limit).
		Preload("Photos", "is_main = true").
		Find(&restaurants).Error

	return restaurants, err
}
