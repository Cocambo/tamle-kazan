package seed

import (
	"context"
	"fmt"
	"os"
	"path/filepath"
	"restaurant-service/internal/models"
	"sort"
	"strings"
	"time"

	"gorm.io/gorm"
)

func LoadSeedData(ctx context.Context, db *gorm.DB, photosDir string) error {
	var count int64
	if err := db.WithContext(ctx).Model(&models.Restaurant{}).Count(&count).Error; err != nil {
		return err
	}
	if count > 0 {
		return nil
	}

	restaurants := []models.Restaurant{
		{
			Name:        "Туган Авылым",
			Description: "Татарская кухня в атмосфере традиционного татарского села",
			Address:     "ул. Тукая, 120",
			Cuisine:     "Татарская",
			AverageBill: 1500,
		},
		{
			Name:        "Дом Татарской Кулинарии",
			Description: "Аутентичная татарская кухня в центре Казани",
			Address:     "ул. Баумана, 58",
			Cuisine:     "Татарская",
			AverageBill: 1200,
		},
		{
			Name:        "Чайхана",
			Description: "Узбекская и татарская кухня, большие порции",
			Address:     "ул. Баумана, 77",
			Cuisine:     "Среднеазиатская",
			AverageBill: 1000,
		},
		{
			Name:        "Казанская кухня",
			Description: "Современная интерпретация татарских блюд",
			Address:     "ул. Профсоюзная, 29",
			Cuisine:     "Татарская",
			AverageBill: 1800,
		},
		{
			Name:        "Паста & Пинца",
			Description: "Итальянская кухня в центре Казани",
			Address:     "ул. Баумана, 58/8",
			Cuisine:     "Итальянская",
			AverageBill: 1400,
		},
	}

	for i := range restaurants {
		restaurant := &restaurants[i]

		if err := db.WithContext(ctx).Create(restaurant).Error; err != nil {
			return fmt.Errorf("failed to create restaurant %s: %w", restaurant.Name, err)
		}

		restaurantIndex := i + 1
		prefix := fmt.Sprintf("restaurant_%d_", restaurantIndex)

		entries, err := os.ReadDir(photosDir)
		if err != nil {
			return fmt.Errorf("failed to read photos dir: %w", err)
		}

		var photoFiles []string
		for _, entry := range entries {
			if entry.IsDir() {
				continue
			}

			name := entry.Name()
			if strings.HasPrefix(name, prefix) &&
				(strings.HasSuffix(name, ".jpg") || strings.HasSuffix(name, ".jpeg") || strings.HasSuffix(name, ".png")) {
				photoFiles = append(photoFiles, name)
			}
		}

		if len(photoFiles) == 0 {
			return fmt.Errorf("no photos found for restaurant %d (%s)", restaurantIndex, restaurant.Name)
		}

		// стабильный порядок
		sort.Strings(photoFiles)

		for idx, fileName := range photoFiles {
			filePath := filepath.Join(photosDir, fileName)
			if _, err := os.Stat(filePath); err != nil {
				return fmt.Errorf("photo file not found: %s", filePath)
			}

			photo := models.Photo{
				RestaurantID: restaurant.ID,
				URL:          "/photos/" + fileName,
				IsMain:       idx == 0,
				CreatedAt:    time.Now(),
			}

			if err := db.WithContext(ctx).Create(&photo).Error; err != nil {
				return fmt.Errorf("failed to create photo %s: %w", fileName, err)
			}
		}
	}

	return nil
}
