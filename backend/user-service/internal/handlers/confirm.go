package handlers

import (
	"net/http"
	"time"

	"github.com/Cocambo/tamle-kazan/backend/user-service/internal/database"
	"github.com/Cocambo/tamle-kazan/backend/user-service/internal/models"
	"github.com/Cocambo/tamle-kazan/backend/user-service/internal/utils"

	"github.com/gin-gonic/gin"
)

// ConfirmEmail — подтверждение email по токену
//
// GET /confirm?token=<token>
// Токен приходит в query параметре.
// Подтверждает email, если токен валиден и не истёк
// После подтверждения сбрасывает токен и помечает email как подтверждённый
// Возвращает 200 OK или ошибку
func ConfirmEmail(c *gin.Context) {
	token := c.Query("token")
	if token == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "token is required"})
		return
	}

	// хэшируем токен для поиска в базе
	tokenHash := utils.HashToken(token)

	// ищем пользователя с таким токеном
	var user models.User
	if err := database.DB.Where("email_token_hash = ?", tokenHash).First(&user).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid or expired token"})
		return
	}

	// проверяем истёк ли токен
	if time.Now().After(user.TokenExpiresAt) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "token expired"})
		return
	}

	// подтверждаем email пользователя
	user.IsEmailConfirmed = true
	user.EmailTokenHash = ""          // инвалидация токена
	user.TokenExpiresAt = time.Time{} // сброс
	if err := database.DB.Save(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to confirm email"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "email confirmed successfully"})
}
