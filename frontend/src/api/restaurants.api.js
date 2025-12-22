import { http } from "@/utils/http";

export const restaurantsApi = {
  getAll(params = {}) {
    return http.get("/restaurants", { params });
  },

  getById(id) {
    return http.get(`/restaurants/${id}`);
  },

  addToFavorites(id) {
    return http.post(`/restaurants/${id}/favorite`);
  },

  getFavorites() {
    return http.get("/restaurants/favorites");
  },

  removeFromFavorites(id) {
    return http.delete(`/restaurants/${id}/favorite`);
  },
  
  getTopRestaurants() {
    return http.get("/restaurants/top");
  },
};
