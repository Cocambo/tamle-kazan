package handlers

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/Cocambo/tamle-kazan/backend/user-service/internal/config"
	"github.com/Cocambo/tamle-kazan/backend/user-service/internal/database"
	"github.com/Cocambo/tamle-kazan/backend/user-service/internal/models"
	"github.com/Cocambo/tamle-kazan/backend/user-service/internal/utils"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

// Используем jwtSecret для подписи токенов
// var jwtSecret = []byte(config.AppConfig.JwtSecret)

func getJWTSecret() []byte {
	return []byte(config.AppConfig.JwtSecret)
}

// Register — создание нового пользователя
//
// POST /register
func Register(c *gin.Context) { // Используем Gin
	var input struct { // принимаем тело запроса, Gin автоматически проверит обязательность полей и формат email.
		FirstName string `json:"first_name" binding:"required"`
		LastName  string `json:"last_name" binding:"required"`
		Email     string `json:"email" binding:"required,email"`
		Password  string `json:"password" binding:"required,min=6"`
	}
	// Парсинг и валидация JSON-запроса
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Проверка, что email ещё не зарегистрирован
	var existing models.User
	if err := database.DB.Where("email = ?", input.Email).First(&existing).Error; err == nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "email already exists"})
		return
	}

	// Хэшируем пароль
	hash, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to hash password"})
		return
	}

	// Создаем пользователя в базе
	user := models.User{
		FirstName: input.FirstName,
		LastName:  input.LastName,
		Email:     input.Email,
		Password:  string(hash),
		Role:      "user",
	}

	// Сохраняем пользователя
	if err := database.DB.Create(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to create user"})
		return
	}

	// Генерируем токен подтверждения и отправляем письмо
	ttl := 20 // время жизни токена в минутах
	token, expiresAt, err := utils.GenerateToken(ttl)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to generate confirmation token"})
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
		if err := utils.SendConfirmationEmail(emailAddr, token); err != nil {
			log.Printf("failed to send confirmation email to %s: %v", emailAddr, err)
		}
	}(user.Email, token)

	c.JSON(http.StatusCreated, gin.H{"message": "user registered successfully; confirmation email sent"})

}

// Login — проверка email/пароля и выдача JWT
//
// POST /login
func Login(c *gin.Context) {
	fmt.Println("User service JWT secret:", config.AppConfig.JwtSecret)
	var input struct { // принимаем тело запроса, Gin автоматически проверит обязательность полей и формат email.
		Email    string `json:"email" binding:"required,email"`
		Password string `json:"password" binding:"required"`
	}
	// Парсинг и валидация JSON-запроса
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// Проверяем существует ли пользователь
	var user models.User
	if err := database.DB.Where("email = ?", input.Email).First(&user).Error; err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid email or password"})
		return
	}

	// Проверка пароля
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(input.Password)); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid email or password"})
		return
	}

	// Создаём JWT токен
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": user.ID,
		"role":    user.Role,
		"exp":     time.Now().Add(24 * time.Hour).Unix(),
	})

	// Подписываем JWT токен
	tokenString, err := token.SignedString(getJWTSecret())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to generate token"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"token": tokenString,
	})
}
