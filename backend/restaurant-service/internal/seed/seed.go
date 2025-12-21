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

// LoadSeedData загружает начальные данные ресторанов и их фотографий в базу данных.
func LoadSeedData(ctx context.Context, db *gorm.DB, photosDir string) error {
	// Проверяем, есть ли уже рестораны в базе данных
	var count int64
	if err := db.WithContext(ctx).Model(&models.Restaurant{}).Count(&count).Error; err != nil {
		return err
	}
	// Если рестораны уже есть, пропускаем загрузку начальных данных
	if count > 0 {
		return nil
	}

	// Определяем начальные данные ресторанов
	restaurants := []models.Restaurant{
		{
			Name:        "Olio первый ресторан",
			Description: "Татарская кухня в атмосфере традиционного татарского села",
			Address:     "ул. Тукая, 120",
			Cuisine:     "Татарская",
			AverageBill: 1500,
		},
		{
			Name:        "Izakaya bereg второй ресторан",
			Description: "Аутентичная татарская кухня в центре Казани",
			Address:     "ул. Баумана, 58",
			Cuisine:     "Татарская",
			AverageBill: 1200,
		},
		// {
		// 	Name:        "Чайхана",
		// 	Description: "Узбекская и татарская кухня, большие порции",
		// 	Address:     "ул. Баумана, 77",
		// 	Cuisine:     "Среднеазиатская",
		// 	AverageBill: 1000,
		// },
		// {
		// 	Name:        "Казанская кухня",
		// 	Description: "Современная интерпретация татарских блюд",
		// 	Address:     "ул. Профсоюзная, 29",
		// 	Cuisine:     "Татарская",
		// 	AverageBill: 1800,
		// },
		// {
		// 	Name:        "Паста & Пинца",
		// 	Description: "Итальянская кухня в центре Казани",
		// 	Address:     "ул. Баумана, 58/8",
		// 	Cuisine:     "Итальянская",
		// 	AverageBill: 1400,
		// },
	}

	// Читаем папку с изображениями один раз
	entries, err := os.ReadDir(photosDir)
	if err != nil {
		return fmt.Errorf("failed to read photos dir: %w", err)
	}

	// Создаем мапу для хранения фотографий по ресторанам
	// В виде: restaurantIndex -> []photoFileNames
	photosByRestaurant := make(map[int][]string)

	for _, entry := range entries {
		// Пропускаем директории
		if entry.IsDir() {
			continue
		}

		// Получаем имя файла
		name := entry.Name()

		// Фильтруем по расширению
		if !(strings.HasSuffix(name, ".jpg") ||
			strings.HasSuffix(name, ".jpeg") ||
			strings.HasSuffix(name, ".png")) {
			continue
		}

		// вытаскиваем индекс ресторана из имени
		// restaurant_1_0.jpg -> index = 1
		var restaurantIndex int
		if _, err := fmt.Sscanf(name, "restaurant_%d_", &restaurantIndex); err == nil {
			photosByRestaurant[restaurantIndex] = append(
				photosByRestaurant[restaurantIndex],
				name,
			)
		}
	}

	// Загружаем рестораны и их фотографии в базу данных в рамках одной транзакции
	return db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		for i := range restaurants {
			restaurant := &restaurants[i]

			// Создаем ресторан в базе данных
			if err := tx.Create(restaurant).Error; err != nil {
				return fmt.Errorf("failed to create restaurant %s: %w", restaurant.Name, err)
			}

			restaurantIndex := i + 1
			photoFiles := photosByRestaurant[restaurantIndex]

			if len(photoFiles) == 0 {
				return fmt.Errorf(
					"no photos found for restaurant %d (%s)",
					restaurantIndex,
					restaurant.Name,
				)
			}

			// Гарантируем стабильный порядок
			sort.Strings(photoFiles)

			// Создаем записи фотографий в базе данных
			for idx, fileName := range photoFiles {
				filePath := filepath.Join(photosDir, fileName)
				if _, err := os.Stat(filePath); err != nil {
					return fmt.Errorf("photo file not found: %s", filePath)
				}

				photo := models.Photo{
					RestaurantID: restaurant.ID,
					URL:          "/api/photos/" + fileName,
					IsMain:       idx == 0, // первая фотография - главная
					CreatedAt:    time.Now(),
				}

				if err := tx.Create(&photo).Error; err != nil {
					return fmt.Errorf("failed to create photo %s: %w", fileName, err)
				}
			}
		}
		return nil
	})
}
