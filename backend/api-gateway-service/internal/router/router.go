package router

import (
	"github.com/Cocambo/tamle-kazan/backend/api-gateway-service/internal/config"
	"github.com/Cocambo/tamle-kazan/backend/api-gateway-service/internal/middleware"
	"github.com/Cocambo/tamle-kazan/backend/api-gateway-service/internal/proxy"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func SetupRouter(cfg *config.Config) *gin.Engine {
	r := gin.Default()

	r.Use(cors.New(cors.Config{
		AllowOrigins:     cfg.CORS.AllowOrigins,
		AllowMethods:     cfg.CORS.AllowMethods,
		AllowHeaders:     cfg.CORS.AllowHeaders,
		ExposeHeaders:    cfg.CORS.ExposeHeaders,
		AllowCredentials: cfg.CORS.AllowCredentials,
		MaxAge:           cfg.CORS.MaxAge,
	}))

	userProxy := proxy.NewProxy(cfg.UserServiceURL, "/api/user")

	// Открытые маршруты user-service
	publicUser := r.Group("/api/user")
	{
		publicUser.POST("/register", userProxy)
		publicUser.POST("/login", userProxy)

		publicUser.POST("/refresh", userProxy)
		publicUser.POST("/logout", userProxy)

		publicUser.GET("/confirm-email", userProxy)
		publicUser.POST("/resend-confirmation", userProxy)
	}
	// Закрытые маршруты с JWT аутентификацией
	auth := r.Group("/api/user")
	auth.Use(middleware.JWTMiddleware(cfg.JwtSecret))
	{
		auth.GET("/profile", proxy.UserProfileProxy(cfg.UserServiceURL))
		auth.GET("/profile/top", proxy.NewProxy(cfg.RestaurantServiceURL, "/api/user/"))

	}

	restaurantProxy := proxy.NewProxy(cfg.RestaurantServiceURL, "/api")

	// Открытые маршруты restaurant-service
	publicRestaurant := r.Group("/api/restaurants")
	{
		publicRestaurant.GET("", restaurantProxy)
		publicRestaurant.GET("/top", restaurantProxy)
		publicRestaurant.GET("/:id", restaurantProxy)
		publicRestaurant.GET("/:id/reviews", restaurantProxy)
	}

	authorizedRestaurant := r.Group("/api/restaurants")
	authorizedRestaurant.Use(middleware.JWTMiddleware(cfg.JwtSecret))
	{
		authorizedRestaurant.GET("/favorites", restaurantProxy)
		authorizedRestaurant.POST("/:id/reviews", restaurantProxy)
		authorizedRestaurant.POST("/:id/favorite", restaurantProxy)
		authorizedRestaurant.DELETE("/:id/favorite", restaurantProxy)
	}

	photosProxy := proxy.NewProxy(cfg.RestaurantServiceURL, "/api")
	r.GET("/api/photos/*path", photosProxy)

	return r
}
