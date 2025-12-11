package proxy

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// UserProfileProxy создаёт прокси-обработчик для получения профиля пользователя
// из user-service, используя user_id из контекста Gin.
// Делает запрос к user-service по пути /users/<id>
// и возвращает ответ клиенту.
func UserProfileProxy(userServiceURL string) gin.HandlerFunc {
	//Берем user_id из контекста Gin, его мы заранее положили туда, когда проверяли JWT
	return func(c *gin.Context) {
		uidAny, exists := c.Get("user_id")
		if !exists {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "user_id not found in token"})
			return
		}

		uid := uidAny.(uint)

		// Делаем запрос на user-service: /users/<id>
		target := fmt.Sprintf("%s/users/%d", userServiceURL, uid)

		req, err := http.NewRequest("GET", target, nil)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to build request"})
			return
		}

		// Пробрасываем заголовки из исходного запроса
		req.Header = c.Request.Header

		// Выполняем запрос к user-service
		client := &http.Client{}
		resp, err := client.Do(req)
		if err != nil {
			c.JSON(http.StatusBadGateway, gin.H{"error": "failed to reach user service"})
			return
		}
		defer resp.Body.Close()

		// Пробрасываем ответ от user-service обратно клиенту
		c.DataFromReader(resp.StatusCode, resp.ContentLength, resp.Header.Get("Content-Type"), resp.Body, nil)
	}
}
