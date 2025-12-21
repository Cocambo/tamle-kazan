package models

import "time"

// Restaurant хранит информацию о ресторане. Включает связи с отзывами, избранным и фотографиями.
type Restaurant struct {
	ID           uint      `gorm:"primaryKey" json:"id"`
	Name         string    `gorm:"not null" json:"name"`
	Description  string    `json:"description"`
	Address      string    `gorm:"not null" json:"address"`
	Cuisine      string    `gorm:"not null" json:"cuisine"`
	AverageBill  float64   `gorm:"not null" json:"average_bill"`
	Rating       float64   `gorm:"default:0" json:"rating"`
	ReviewsCount int       `gorm:"default:0" json:"reviews_count"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`

	// Связи
	// 1 ресторан - много отзывов, foreignKey:RestaurantID  - поле в Review, omitempty - если отзывов нет, поле не вернётся в JSON
	Reviews []Review `gorm:"foreignKey:RestaurantID" json:"reviews,omitempty"`
	// 1 ресторан - много избранных, foreignKey:RestaurantID - поле в Favorite
	Favorites []Favorite `gorm:"foreignKey:RestaurantID" json:"favorites,omitempty"`
	// 1 ресторан - много фотографий, foreignKey:RestaurantID - поле в Photo
	Photos []Photo `gorm:"foreignKey:RestaurantID" json:"photos,omitempty"`
}

type Review struct {
	ID uint `gorm:"primaryKey" json:"id"`

	// Уникальный индекс по паре (UserID, RestaurantID) для предотвращения дублирующих отзывов от одного пользователя на один ресторан
	RestaurantID uint `gorm:"not null;index;uniqueIndex:idx_reviews_user_restaurant" json:"restaurant_id"`
	UserID       uint `gorm:"not null;index;uniqueIndex:idx_reviews_user_restaurant" json:"user_id"`

	Rating    int       `gorm:"not null;check:rating >= 1 AND rating <= 5" json:"rating"`
	Comment   string    `json:"comment"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`

	// Связь с рестораном
	Restaurant Restaurant `gorm:"foreignKey:RestaurantID;references:ID" json:"restaurant,omitempty"`
}

type Favorite struct {
	ID uint `gorm:"primaryKey" json:"id"`

	// Уникальный индекс по паре (UserID, RestaurantID) для предотвращения дублирующих избранных от одного пользователя на один ресторан
	UserID       uint `gorm:"not null;index;uniqueIndex:idx_favorites_user_restaurant" json:"user_id"`
	RestaurantID uint `gorm:"not null;index;uniqueIndex:idx_favorites_user_restaurant" json:"restaurant_id"`

	CreatedAt time.Time `json:"created_at"`

	// Связь с рестораном
	Restaurant Restaurant `gorm:"foreignKey:RestaurantID" json:"restaurant,omitempty"`
}

type Photo struct {
	ID           uint      `gorm:"primaryKey" json:"id"`
	RestaurantID uint      `gorm:"not null;index" json:"restaurant_id"`
	URL          string    `gorm:"not null" json:"url"`
	IsMain       bool      `gorm:"default:false" json:"is_main"`
	CreatedAt    time.Time `json:"created_at"`

	// Связь с рестораном
	// Restaurant Restaurant `gorm:"foreignKey:RestaurantID" json:"restaurant,omitempty"`
	Restaurant Restaurant `json:"-"`
}
