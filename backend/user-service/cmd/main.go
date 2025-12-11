package main

import (
	"log"

	"github.com/Cocambo/tamle-kazan/backend/user-service/internal/config"
	"github.com/Cocambo/tamle-kazan/backend/user-service/internal/database"
	"github.com/Cocambo/tamle-kazan/backend/user-service/internal/handlers"

	"github.com/gin-gonic/gin"
)

func main() {
	// Загружаем конфигурацию
	config.LoadConfig()
	log.Println("The configuration is loaded")

	//Подключаемся к базе данных
	database.InitDB()
	database.Migrate()

	//Создаём Gin router
	r := gin.Default()

	// Публичные маршруты — не требуют авторизации
	r.POST("/register", handlers.Register)
	r.POST("/login", handlers.Login)

	r.GET("/confirm-email", handlers.ConfirmEmail)
	r.POST("/resend-confirmation", handlers.ResendConfirmation)

	// Приватные маршруты — требуют JWT-токен
	// auth := r.Group("/")
	// auth.GET("/users/:id", handlers.GetUser)
	// auth.PUT("/users/:id", handlers.UpdateUser)

	userRoutes := r.Group("/users")
	{
		userRoutes.GET("/:id", handlers.GetUser)
		userRoutes.PUT("/:id", handlers.UpdateUser)
	}

	// Только для администраторов
	//auth.PUT("/users/:id/role", middleware.AdminMiddleware(), handlers.ChangeRole)
	// auth.Use(middleware.AuthMiddleware())
	// {
	// 	auth.GET("/users/:id", handlers.GetUser)
	// 	auth.PUT("/users/:id", handlers.UpdateUser)

	// 	// Только для администраторов
	// 	auth.PUT("/users/:id/role", middleware.AdminMiddleware(), handlers.ChangeRole)
	// }

	// Запуск HTTP-сервера
	port := config.AppConfig.ServerPort
	if port == "" {
		port = "8081"
	}

	log.Printf("Server start on port %s", port)
	r.Run(":" + port)
}
