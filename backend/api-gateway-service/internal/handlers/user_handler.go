package handlers

import (
	"bytes"
	"io"
	"log"
	"net/http"

	"github.com/Cocambo/tamle-kazan/backend/api-gateway-service/internal/config"
	"github.com/gin-gonic/gin"
)

// Проксирование регистрации
func Register(c *gin.Context) {
	forwardRequest(c, "/register")
}

// Проксирование логина
func Login(c *gin.Context) {
	forwardRequest(c, "/login")
}

// Проксирование получения пользователя
func GetUser(c *gin.Context) {
	userID := c.Param("id")
	forwardRequest(c, "/user/"+userID)
}

// Универсальная функция для пересылки запроса
func forwardRequest(c *gin.Context, path string) {
	targetURL := config.AppConfig.UserServiceURL + path

	// Читаем тело запроса
	bodyBytes, _ := io.ReadAll(c.Request.Body)

	// Создаём новый HTTP-запрос
	req, err := http.NewRequest(c.Request.Method, targetURL, bytes.NewBuffer(bodyBytes))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to create request"})
		return
	}
	// Копируем заголовки
	req.Header = c.Request.Header.Clone()
	client := &http.Client{}
	// Отправляем запрос
	resp, err := client.Do(req)
	if err != nil {
		log.Printf("Error forwarding request: %v", err)
		c.JSON(http.StatusBadGateway, gin.H{"error": "user-service unavailable"})
		return
	}
	defer resp.Body.Close()

	// Возвращаем ответ клиенту
	respBody, _ := io.ReadAll(resp.Body)
	c.Data(resp.StatusCode, resp.Header.Get("Content-Type"), respBody)
}
