package router

import (
	"github.com/Cocambo/tamle-kazan/backend/api-gateway-service/internal/handlers"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	api := r.Group("/api")
	{
		api.POST("/register", handlers.Register)
		api.POST("/login", handlers.Login)
		api.GET("/user/:id", handlers.GetUser)
	}

	return r
}
