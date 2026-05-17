package service

import (
	"context"
	"math"
	"math/rand"
	"sort"
	"time"

	"restaurant-service/internal/models"
)

const (
	highRatingThreshold        = 4
	recommendationCandidateCap = 20
	recommendationPoolLimit    = 7
	topRestaurantsLimit        = 3
)

type preferenceProfile struct {
	TopCuisines []string
	AverageBill float64
	ExcludeIDs  []uint
}

type scoredRestaurant struct {
	Restaurant models.Restaurant
	Score      float64
}

func (s *Service) GetRecommendationsByUser(ctx context.Context, userID uint, limit int) ([]models.Restaurant, error) {
	favorites, err := s.repo.GetUserFavoriteRestaurants(ctx, userID)
	if err != nil {
		return nil, err
	}

	highRated, err := s.repo.GetUserHighlyRatedRestaurants(ctx, userID, highRatingThreshold)
	if err != nil {
		return nil, err
	}

	likedRestaurants := mergeUniqueRestaurants(favorites, highRated)
	if len(likedRestaurants) == 0 {
		return s.getColdStartRecommendations(ctx, nil, limit)
	}

	profile := buildPreferenceProfile(likedRestaurants)

	candidates, err := s.repo.GetCandidateRestaurantsByCuisines(
		ctx,
		profile.TopCuisines,
		profile.ExcludeIDs,
		recommendationCandidateCap,
	)
	if err != nil {
		return nil, err
	}

	if len(candidates) < limit {
		popular, err := s.repo.GetPopularRestaurantsExcluding(ctx, profile.ExcludeIDs, recommendationCandidateCap)
		if err != nil {
			return nil, err
		}
		candidates = mergeUniqueRestaurants(candidates, popular)
	}

	if len(candidates) == 0 {
		return s.getColdStartRecommendations(ctx, profile.ExcludeIDs, limit)
	}

	scored := make([]scoredRestaurant, 0, len(candidates))
	for _, restaurant := range candidates {
		scored = append(scored, scoredRestaurant{
			Restaurant: restaurant,
			Score:      scoreRestaurant(restaurant, profile),
		})
	}

	sort.Slice(scored, func(i, j int) bool {
		if scored[i].Score == scored[j].Score {
			return scored[i].Restaurant.ID < scored[j].Restaurant.ID
		}
		return scored[i].Score > scored[j].Score
	})

	result := make([]models.Restaurant, 0, limit)
	for _, item := range scored {
		result = append(result, item.Restaurant)
		if len(result) == limit {
			break
		}
	}

	return result, nil
}

func (s *Service) getColdStartRecommendations(ctx context.Context, excludeIDs []uint, limit int) ([]models.Restaurant, error) {
	pool, err := s.repo.GetPopularRestaurantsExcluding(ctx, excludeIDs, recommendationPoolLimit)
	if err != nil {
		return nil, err
	}

	if len(pool) == 0 {
		return []models.Restaurant{}, nil
	}

	selectionPool := pool
	if len(pool) > topRestaurantsLimit {
		selectionPool = pool[topRestaurantsLimit:]
	}

	if len(selectionPool) < limit {
		selectionPool = pool
	}

	return pickRandomRestaurants(selectionPool, limit), nil
}

func buildPreferenceProfile(restaurants []models.Restaurant) preferenceProfile {
	cuisineCount := make(map[string]int)
	excludeMap := make(map[uint]struct{})

	var totalBill float64
	for _, restaurant := range restaurants {
		cuisineCount[restaurant.Cuisine]++
		totalBill += restaurant.AverageBill
		excludeMap[restaurant.ID] = struct{}{}
	}

	topCuisines := getTopCuisines(cuisineCount, 2)

	excludeIDs := make([]uint, 0, len(excludeMap))
	for id := range excludeMap {
		excludeIDs = append(excludeIDs, id)
	}

	return preferenceProfile{
		TopCuisines: topCuisines,
		AverageBill: totalBill / float64(len(restaurants)),
		ExcludeIDs:  excludeIDs,
	}
}

func scoreRestaurant(restaurant models.Restaurant, profile preferenceProfile) float64 {
	score := restaurant.Rating*2 + math.Log1p(float64(restaurant.ReviewsCount))

	if len(profile.TopCuisines) > 0 && restaurant.Cuisine == profile.TopCuisines[0] {
		score += 5
	} else if len(profile.TopCuisines) > 1 && restaurant.Cuisine == profile.TopCuisines[1] {
		score += 3
	}

	billDiff := math.Abs(restaurant.AverageBill - profile.AverageBill)
	switch {
	case billDiff <= 300:
		score += 3
	case billDiff <= 700:
		score += 1.5
	}

	return score
}

func mergeUniqueRestaurants(groups ...[]models.Restaurant) []models.Restaurant {
	unique := make(map[uint]models.Restaurant)

	for _, group := range groups {
		for _, restaurant := range group {
			unique[restaurant.ID] = restaurant
		}
	}

	result := make([]models.Restaurant, 0, len(unique))
	for _, restaurant := range unique {
		result = append(result, restaurant)
	}

	return result
}

func getTopCuisines(cuisineCount map[string]int, limit int) []string {
	type cuisineStat struct {
		Name  string
		Count int
	}

	stats := make([]cuisineStat, 0, len(cuisineCount))
	for name, count := range cuisineCount {
		stats = append(stats, cuisineStat{Name: name, Count: count})
	}

	sort.Slice(stats, func(i, j int) bool {
		if stats[i].Count == stats[j].Count {
			return stats[i].Name < stats[j].Name
		}
		return stats[i].Count > stats[j].Count
	})

	result := make([]string, 0, limit)
	for _, item := range stats {
		result = append(result, item.Name)
		if len(result) == limit {
			break
		}
	}

	return result
}

func pickRandomRestaurants(restaurants []models.Restaurant, limit int) []models.Restaurant {
	if len(restaurants) <= limit {
		return restaurants
	}

	shuffled := append([]models.Restaurant(nil), restaurants...)
	randomizer := rand.New(rand.NewSource(time.Now().UnixNano()))
	randomizer.Shuffle(len(shuffled), func(i, j int) {
		shuffled[i], shuffled[j] = shuffled[j], shuffled[i]
	})

	return shuffled[:limit]
}
