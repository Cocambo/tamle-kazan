package router

import (
	"restaurant-service/internal/handlers"

	"github.com/gin-gonic/gin"
)

func SetupRouter(h *handlers.Handlers) *gin.Engine {
	router := gin.Default()

	// Public endpoints
	public := router.Group("/restaurants")
	{
		public.GET("", h.GetRestaurants)
		public.GET("/:id", h.GetRestaurantByID)
		public.GET("/:id/reviews", h.GetReviews)
	}

	// Authorized endpoints (user_id comes from X-User-ID header via API Gateway)
	authorized := router.Group("/restaurants")
	{
		authorized.GET("/favorites", h.GetUserFavorites)
		authorized.POST("/:id/reviews", h.CreateReview)
		authorized.POST("/:id/favorite", h.AddFavorite)
		authorized.DELETE("/:id/favorite", h.RemoveFavorite)
	}

	return router
}
