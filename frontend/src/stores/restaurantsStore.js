import { defineStore } from "pinia";
import { restaurantsApi } from "@/api/restaurants.api";

export const useRestaurantsStore = defineStore("restaurants", {
  state: () => ({
    restaurants: [],
    currentRestaurant: null,
    loading: false,
  }),

  actions: {
    async fetchRestaurants(filters = {}) {
      try {
        this.loading = true;
        const { data } = await restaurantsApi.getAll(filters);
        this.restaurants = data;
      } catch (e) {
        console.error("Ошибка загрузки ресторанов", e);
      } finally {
        this.loading = false;
      }
    },

    async fetchRestaurantById(id) {
      try {
        this.loading = true;
        const { data } = await restaurantsApi.getById(id);
        this.currentRestaurant = data;
      } catch (e) {
        console.error("Ошибка загрузки ресторана", e);
      } finally {
        this.loading = false;
      }
    },
  },
});
