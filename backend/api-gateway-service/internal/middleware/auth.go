package middleware

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

// JWTMiddleware(jwtSecret string) возвращает gin.HandlerFunc,
// и проверяет JWT в заголовке Authorization и пропускает запрос дальше только если токен валиден.
//
// При ошибке возвращает 401 и завершает обработку.
func JWTMiddleware(jwtSecret string) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Извлечение токена из заголовка Authorization, в случае отсутствия, возвращаем 401
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header required"})
			c.Abort()
			return
		}
		// Проверка формата Bearer <token>, ожидается два элемента: Bearer и сам токен
		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid Authorization header format"})
			c.Abort()
			return
		}
		// Извлечение самого токена
		tokenString := parts[1]
		// Проверка подписи и парсинг токена
		token, err := jwt.ParseWithClaims(
			tokenString,
			&jwt.RegisteredClaims{},
			func(token *jwt.Token) (interface{}, error) {
				// Проверяем ожидаемый алгоритм подписи
				if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
					return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
				}
				// Возвращается секрет для проверки подписи
				return []byte(jwtSecret), nil
			})
		if err != nil || !token.Valid {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid or expired token"})
			c.Abort()
			return
		}

		c.Next()
	}
}
