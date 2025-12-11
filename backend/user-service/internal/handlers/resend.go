package handlers

import (
	"net/http"
	"time"

	"github.com/Cocambo/tamle-kazan/backend/user-service/internal/database"
	"github.com/Cocambo/tamle-kazan/backend/user-service/internal/models"
	"github.com/Cocambo/tamle-kazan/backend/user-service/internal/utils"

	"github.com/gin-gonic/gin"
)

// ResendConfirmation — повторная отправка письма с подтверждением email
//
// POST /resend-confirmation
// Тело запроса: { "email": "<email>" }.
// Если пользователь с таким email существует и не подтверждён — отправляет новое письмо с токеном.
// Ограничение по частоте: 1 письмо в минуту.
// Возвращает 200 OK независимо от результата (чтобы не раскрывать информацию о существовании email)
func ResendConfirmation(c *gin.Context) {
	// Парсим тело запроса
	var input struct {
		Email string `json:"email" binding:"required,email"`
	}
	// Парсинг и валидация JSON-запроса
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// Ищем пользователя по email
	var user models.User
	if err := database.DB.Where("email = ?", input.Email).First(&user).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "user not found"})
		return
	}
	// Проверяем, подтверждён ли email
	if user.IsEmailConfirmed {
		c.JSON(http.StatusBadRequest, gin.H{"error": "email already confirmed"})
		return
	}

	// Ограничение повторной отправки: например, 1 письмо в 1 минуту
	cooldown := time.Minute
	if time.Since(user.LastConfirmSentAt) < cooldown {
		c.JSON(http.StatusTooManyRequests, gin.H{"error": "too many requests, try later"})
		return
	}

	// Генерируем новый токен подтверждения и сохраняем его в базе
	ttl := 20
	token, expiresAt, err := utils.GenerateToken(ttl)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to generate token"})
		return
	}

	// Сохраняем хэш токена и время истечения в базе
	user.EmailTokenHash = utils.HashToken(token)
	user.TokenExpiresAt = expiresAt
	user.LastConfirmSentAt = time.Now()
	if err := database.DB.Save(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to save confirmation token"})
		return
	}
	// Отправляем письмо (асинхронно в goroutine — чтобы не блокировать запрос)
	go func(emailAddr, token string) {
		_ = utils.SendConfirmationEmail(emailAddr, token)
	}(user.Email, token)

	c.JSON(http.StatusOK, gin.H{"message": "confirmation email resent if user exists"})
}
