package handlers

import (
	"net/http"
	"strconv"

	"restaurant-service/internal/service"

	"github.com/gin-gonic/gin"
)

func (h *Handlers) GetRestaurants(c *gin.Context) {
	ctx := c.Request.Context()

	filter := service.RestaurantFilter{
		Limit: 50, // default limit
	}

	// Обрабатываем параметры фильтрации из query-параметров
	if search := c.Query("search"); search != "" {
		filter.Search = search
	}

	if cuisine := c.Query("cuisine"); cuisine != "" {
		filter.Cuisine = cuisine
	}

	if minBillStr := c.Query("min_bill"); minBillStr != "" {
		if minBill, err := strconv.ParseFloat(minBillStr, 64); err == nil {
			filter.MinBill = &minBill
		}
	}

	if maxBillStr := c.Query("max_bill"); maxBillStr != "" {
		if maxBill, err := strconv.ParseFloat(maxBillStr, 64); err == nil {
			filter.MaxBill = &maxBill
		}
	}

	if minRatingStr := c.Query("min_rating"); minRatingStr != "" {
		if minRating, err := strconv.ParseFloat(minRatingStr, 64); err == nil {
			filter.MinRating = &minRating
		}
	}

	if limitStr := c.Query("limit"); limitStr != "" {
		if limit, err := strconv.Atoi(limitStr); err == nil && limit > 0 {
			filter.Limit = limit
		}
	}

	if offsetStr := c.Query("offset"); offsetStr != "" {
		if offset, err := strconv.Atoi(offsetStr); err == nil && offset >= 0 {
			filter.Offset = offset
		}
	}

	// Получаем рестораны по фильтру
	restaurants, err := h.service.GetRestaurants(ctx, filter)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, restaurants)
}

// GetRestaurantByID получает ресторан по его ID.
func (h *Handlers) GetRestaurantByID(c *gin.Context) {
	ctx := c.Request.Context()

	// Получаем ID ресторана из параметров пути
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid restaurant id"})
		return
	}

	restaurant, err := h.service.GetRestaurantByID(ctx, uint(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if restaurant == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "restaurant not found"})
		return
	}

	c.JSON(http.StatusOK, restaurant)
}

func (h *Handlers) GetTopRestaurants(c *gin.Context) {
	ctx := c.Request.Context()

	const topLimit = 3

	restaurants, err := h.service.GetTopRestaurants(ctx, topLimit)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, restaurants)
}
