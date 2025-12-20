package seed

import (
	"context"
	"fmt"
	"log"
	"restaurant-service/internal/models"
	"time"

	"gorm.io/gorm"
)

// LoadSeedData загружает начальные данные в БД
func LoadSeedData(ctx context.Context, db *gorm.DB, photosDir string) error {
	// Проверяем, есть ли уже данные
	var count int64
	db.WithContext(ctx).Model(&models.Restaurant{}).Count(&count)
	if count > 0 {
		return nil // Данные уже загружены
	}

	restaurants := []models.Restaurant{
		{
			Name:        "Туган Авылым",
			Description: "Татарская кухня в атмосфере традиционного татарского села",
			Address:     "ул. Тукая, 120",
			Cuisine:     "Татарская",
			AverageBill: 1500.0,
			Rating:      0.0,
			ReviewsCount: 0,
		},
		{
			Name:        "Дом Татарской Кулинарии",
			Description: "Аутентичная татарская кухня в центре Казани",
			Address:     "ул. Баумана, 58",
			Cuisine:     "Татарская",
			AverageBill: 1200.0,
			Rating:      0.0,
			ReviewsCount: 0,
		},
		{
			Name:        "Чайхана",
			Description: "Узбекская и татарская кухня, большие порции",
			Address:     "ул. Баумана, 77",
			Cuisine:     "Среднеазиатская",
			AverageBill: 1000.0,
			Rating:      0.0,
			ReviewsCount: 0,
		},
		{
			Name:        "Казанская кухня",
			Description: "Современная интерпретация татарских блюд",
			Address:     "ул. Профсоюзная, 29",
			Cuisine:     "Татарская",
			AverageBill: 1800.0,
			Rating:      0.0,
			ReviewsCount: 0,
		},
		{
			Name:        "Паста & Пинца",
			Description: "Итальянская кухня в центре Казани",
			Address:     "ул. Баумана, 58/8",
			Cuisine:     "Итальянская",
			AverageBill: 1400.0,
			Rating:      0.0,
			ReviewsCount: 0,
		},
		{
			Name:        "Суши Мастер",
			Description: "Японская кухня, роллы и суши",
			Address:     "ул. Пушкина, 12",
			Cuisine:     "Японская",
			AverageBill: 1600.0,
			Rating:      0.0,
			ReviewsCount: 0,
		},
		{
			Name:        "Бургер Кинг",
			Description: "Американская кухня, бургеры и картофель фри",
			Address:     "пр. Победы, 100",
			Cuisine:     "Американская",
			AverageBill: 800.0,
			Rating:      0.0,
			ReviewsCount: 0,
		},
		{
			Name:        "Вкусно и Точка",
			Description: "Русская кухня, домашние блюда",
			Address:     "ул. Кремлевская, 35",
			Cuisine:     "Русская",
			AverageBill: 900.0,
			Rating:      0.0,
			ReviewsCount: 0,
		},
		{
			Name:        "Кофеин",
			Description: "Европейская кухня, кофейня-ресторан",
			Address:     "ул. Баумана, 75",
			Cuisine:     "Европейская",
			AverageBill: 1100.0,
			Rating:      0.0,
			ReviewsCount: 0,
		},
		{
			Name:        "Грузинская кухня",
			Description: "Аутентичная грузинская кухня, хачапури и хинкали",
			Address:     "ул. Московская, 21",
			Cuisine:     "Грузинская",
			AverageBill: 1300.0,
			Rating:      0.0,
			ReviewsCount: 0,
		},
	}

	// Создаем рестораны
	for i := range restaurants {
		if err := db.WithContext(ctx).Create(&restaurants[i]).Error; err != nil {
			return fmt.Errorf("failed to create restaurant %s: %w", restaurants[i].Name, err)
		}

		// Создаем заглушку для главного фото (в реальности здесь был бы файл)
		photoURL := fmt.Sprintf("/photos/restaurant_%d_main.jpg", restaurants[i].ID)
		photo := models.Photo{
			RestaurantID: restaurants[i].ID,
			URL:          photoURL,
			IsMain:       true,
			CreatedAt:    time.Now(),
		}
		if err := db.WithContext(ctx).Create(&photo).Error; err != nil {
			log.Printf("Warning: failed to create photo for restaurant %d: %v", restaurants[i].ID, err)
		}

		// Создаем еще одно дополнительное фото
		photo2URL := fmt.Sprintf("/photos/restaurant_%d_2.jpg", restaurants[i].ID)
		photo2 := models.Photo{
			RestaurantID: restaurants[i].ID,
			URL:          photo2URL,
			IsMain:       false,
			CreatedAt:    time.Now(),
		}
		if err := db.WithContext(ctx).Create(&photo2).Error; err != nil {
			log.Printf("Warning: failed to create photo for restaurant %d: %v", restaurants[i].ID, err)
		}
	}

	return nil
}

