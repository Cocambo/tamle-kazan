package handlers

import (
	"net/http"
	"strconv"

	"restaurant-service/internal/models"

	"github.com/gin-gonic/gin"
)

// CreateReviewRequest представляет тело запроса для создания отзыва.
type CreateReviewRequest struct {
	Rating  int    `json:"rating" binding:"required"`
	Comment string `json:"comment"`
}

// GetReviews получает все отзывы для конкретного ресторана
func (h *Handlers) GetReviews(c *gin.Context) {
	ctx := c.Request.Context()

	// Получаем ID ресторана из параметров пути
	idStr := c.Param("id")
	restaurantID, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid restaurant id"})
		return
	}

	// Получаем отзывы по ID ресторана
	reviews, err := h.service.GetReviewsByRestaurantID(ctx, uint(restaurantID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, reviews)
}

// CreateReview создает новый отзыв для конкретного ресторана
func (h *Handlers) CreateReview(c *gin.Context) {
	ctx := c.Request.Context()

	// Получаем user_id из заголовка (API Gateway передает его)
	userIDStr := c.GetHeader("X-User-ID")
	if userIDStr == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "user_id is required"})
		return
	}

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

	// Парсим тело запроса
	var req CreateReviewRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	review := &models.Review{
		RestaurantID: uint(restaurantID),
		UserID:       uint(userID),
		Rating:       req.Rating,
		Comment:      req.Comment,
	}

	// Создаем отзыв
	if err := h.service.CreateReview(ctx, review); err != nil {
		if err.Error() == "restaurant not found" {
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}
		if err.Error() == "user already reviewed this restaurant" {
			c.JSON(http.StatusConflict, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	//c.JSON(http.StatusCreated, review)
	// Возвращаем только необходимые поля, чтобы избежать циклических ссылок
	c.JSON(http.StatusCreated, gin.H{
		"id":            review.ID,
		"restaurant_id": review.RestaurantID,
		"user_id":       review.UserID,
		"rating":        review.Rating,
		"comment":       review.Comment,
		"created_at":    review.CreatedAt,
		"updated_at":    review.UpdatedAt,
	})
}
