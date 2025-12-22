import { defineStore } from "pinia";
import { restaurantsApi } from "@/api/restaurants.api";

export const useRestaurantsStore = defineStore("restaurants", {
  state: () => ({
    restaurants: [],
    currentRestaurant: null,
    favorites: [],
    topRestaurants: [],
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

    async fetchFavorites() {
      try {
        this.loading = true;
        const { data } = await restaurantsApi.getFavorites();
        this.favorites = data.restaurants || [];
      } catch (e) {
        console.error("Ошибка загрузки избранного", e);
      } finally {
        this.loading = false;
      }
    },

    async addRestaurantInFavorites(id) {
      try {
        this.loading = true;
        const { data, status } = await restaurantsApi.addToFavorites(id);
        if (status === 201) {
          this.favorites.push({ id });
          console.log(data.message);
        }
      } catch (e) {
        if (e.response?.status === 409) {
          console.warn(e.response.data.error);
        } else {
          console.error("Ошибка добавления в избранное", e);
        }
      } finally {
        this.loading = false;
      }
    },

    async removeRestaurantFromFavorites(id) {
      try {
        this.loading = true;
        const { data } = await restaurantsApi.removeFromFavorites(id);
        this.favorites = this.favorites.filter(r => r.id !== id);
        console.log(data.message);
      } catch (e) {
        if (e.response?.status === 404) {
          console.warn(e.response.data.error);
        } else {
          console.error("Ошибка удаления из избранного", e);
        }
      } finally {
        this.loading = false;
      }
    },
    
    async fetchTopRestaurants() {
      try {
        this.loading = true;
        const { data } = await restaurantsApi.getTopRestaurants();
        this.topRestaurants = data;
      } catch (e) {
        console.error("Ошибка загрузки топ-ресторанов", e);
      } finally {
        this.loading = false;
      }
    },
  },
});
