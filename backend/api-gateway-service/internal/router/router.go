package router

import (
	"github.com/Cocambo/tamle-kazan/backend/api-gateway-service/internal/config"
	"github.com/Cocambo/tamle-kazan/backend/api-gateway-service/internal/middleware"
	"github.com/Cocambo/tamle-kazan/backend/api-gateway-service/internal/proxy"
	"github.com/gin-gonic/gin"
)

func SetupRouter(cfg *config.Config) *gin.Engine {
	r := gin.Default()

	userProxy := proxy.NewProxy(cfg.UserServiceURL, "/api/user")

	// Открытые маршруты user-service
	publicUser := r.Group("/api/user")
	{
		publicUser.POST("/register", userProxy)
		publicUser.POST("/login", userProxy)
	}
	// Закрытые маршруты с JWT аутентификацией
	auth := r.Group("/api")
	auth.Use(middleware.JWTMiddleware(cfg.JwtSecret))
	{
		auth.GET("/user/profile", userProxy)

	}

	return r
}
