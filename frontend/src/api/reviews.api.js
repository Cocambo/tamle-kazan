import { http } from "@/utils/http";

export const reviewsApi = {
  createRestaurantsReview(restaurantId, review) {
    return http.post(`/restaurants/${restaurantId}/reviews`, review );
  },

  getRestaurantReviews(restaurantId) {
    return http.get(`/restaurants/${restaurantId}/reviews`);
  },
};
