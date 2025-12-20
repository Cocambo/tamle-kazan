package handlers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// AddFavorite добавляет ресторан в избранное для пользователя
func (h *Handlers) AddFavorite(c *gin.Context) {
	ctx := c.Request.Context()

	// Получаем user_id из заголовка (API Gateway передает его)
	userIDStr := c.GetHeader("X-User-ID")
	if userIDStr == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "user_id is required"})
		return
	}

	// Парсим user_id из строки в uint64
	userID, err := strconv.ParseUint(userIDStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid user_id"})
		return
	}

	// Получаем ID ресторана из параметров пути
	idStr := c.Param("id")
	restaurantID, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid restaurant id"})
		return
	}

	// Добавляем ресторан в избранное
	if err := h.service.AddFavorite(ctx, uint(userID), uint(restaurantID)); err != nil {
		if err.Error() == "restaurant not found" {
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}
		if err.Error() == "restaurant already in favorites" {
			c.JSON(http.StatusConflict, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "restaurant added to favorites"})
}

// RemoveFavorite удаляет ресторан из избранного для пользователя
func (h *Handlers) RemoveFavorite(c *gin.Context) {
	ctx := c.Request.Context()

	// Получаем user_id из заголовка (API Gateway передает его)
	userIDStr := c.GetHeader("X-User-ID")
	if userIDStr == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "user_id is required"})
		return
	}

	// Парсим user_id из строки в uint64
	userID, err := strconv.ParseUint(userIDStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid user_id"})
		return
	}

	// Получаем ID ресторана из параметров пути
	idStr := c.Param("id")
	restaurantID, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid restaurant id"})
		return
	}

	// Удаляем ресторан из избранного
	if err := h.service.RemoveFavorite(ctx, uint(userID), uint(restaurantID)); err != nil {
		if err.Error() == "favorite not found" {
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "restaurant removed from favorites"})
}

// GetUserFavorites получает все избранные рестораны пользователя
func (h *Handlers) GetUserFavorites(c *gin.Context) {
	ctx := c.Request.Context()

	// Получаем user_id из заголовка (API Gateway передает его)
	userIDStr := c.GetHeader("X-User-ID")
	if userIDStr == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "user_id is required"})
		return
	}

	// Парсим user_id из строки в uint64
	userID, err := strconv.ParseUint(userIDStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid user_id"})
		return
	}

	// Получаем все избранные рестораны пользователя
	restaurants, err := h.service.GetUserFavorites(ctx, uint(userID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"restaurants": restaurants})
}
