import { defineStore } from "pinia";
import { useAuthStore } from "@/stores/authStore";
import { restaurantsApi } from "@/api/restaurants.api";

const normalizeRestaurantId = (id) => {
  const normalizedId = Number(id);
  return Number.isNaN(normalizedId) ? id : normalizedId;
};

export const useRestaurantsStore = defineStore("restaurants", {
  state: () => ({
    restaurants: [],
    currentRestaurant: null,
    favorites: [],
    favoriteIds: new Set(),
    favoritesLoaded: false,
    favoritesPromise: null,
    topRestaurants: [],
    topUserRestaurants: [],
    recommendationRestaurants: [],
    loading: false,
  }),

  actions: {
    resetFavoritesState() {
      this.favorites = [];
      this.favoriteIds = new Set();
      this.favoritesLoaded = false;
      this.favoritesPromise = null;
    },

    getKnownRestaurantById(id) {
      const normalizedId = normalizeRestaurantId(id);
      const restaurant =
        this.restaurants.find((item) => normalizeRestaurantId(item.id) === normalizedId) ||
        this.topRestaurants.find((item) => normalizeRestaurantId(item.id) === normalizedId) ||
        this.topUserRestaurants.find((item) => normalizeRestaurantId(item.id) === normalizedId) ||
        this.recommendationRestaurants.find((item) => normalizeRestaurantId(item.id) === normalizedId) ||
        (this.currentRestaurant &&
        normalizeRestaurantId(this.currentRestaurant.id) === normalizedId
          ? this.currentRestaurant
          : null);

      return restaurant ? { ...restaurant } : null;
    },

    addFavoriteLocally(id) {
      const normalizedId = normalizeRestaurantId(id);
      const nextFavoriteIds = new Set(this.favoriteIds);
      nextFavoriteIds.add(normalizedId);
      this.favoriteIds = nextFavoriteIds;

      const hasRestaurant = this.favorites.some(
        (restaurant) => normalizeRestaurantId(restaurant.id) === normalizedId
      );

      if (!hasRestaurant) {
        const restaurant = this.getKnownRestaurantById(normalizedId) || { id: normalizedId };
        this.favorites = [...this.favorites, restaurant];
      }
    },

    removeFavoriteLocally(id) {
      const normalizedId = normalizeRestaurantId(id);
      this.favoriteIds = new Set(
        [...this.favoriteIds].filter((favoriteId) => favoriteId !== normalizedId)
      );
      this.favorites = this.favorites.filter(
        (restaurant) => normalizeRestaurantId(restaurant.id) !== normalizedId
      );
    },

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

    async fetchFavorites(options = {}) {
      const { force = false } = options;
      const authStore = useAuthStore();

      if (!authStore.isAuthenticated) {
        this.resetFavoritesState();
        return [];
      }

      if (force) {
        this.favoritesLoaded = false;
        this.favoritesPromise = null;
      }

      if (this.favoritesLoaded) return this.favorites;
      if (this.favoritesPromise) return this.favoritesPromise;

      try {
        this.loading = true;
        this.favoritesPromise = restaurantsApi.getFavorites().then(({ data }) => {
          this.favorites = data.restaurants || [];
          this.favoriteIds = new Set(
            this.favorites.map((restaurant) => normalizeRestaurantId(restaurant.id))
          );
          this.favoritesLoaded = true;

          return this.favorites;
        });

        return await this.favoritesPromise;
      } catch (e) {
        console.error("Ошибка загрузки избранного", e);
        return [];
      } finally {
        this.favoritesPromise = null;
        this.loading = false;
      }
    },

    async addRestaurantInFavorites(id) {
      const authStore = useAuthStore();
      const normalizedId = normalizeRestaurantId(id);

      if (!authStore.isAuthenticated) {
        return false;
      }

      const previousFavorites = [...this.favorites];
      const previousFavoriteIds = new Set(this.favoriteIds);

      this.addFavoriteLocally(normalizedId);

      try {
        this.loading = true;
        const { data, status } = await restaurantsApi.addToFavorites(normalizedId);
        if (status === 201) {
          console.log(data.message);
        }

        return true;
      } catch (e) {
        if (e.response?.status === 409) {
          console.warn(e.response.data.error);
          this.favoritesLoaded = false;
          return true;
        }

        this.favorites = previousFavorites;
        this.favoriteIds = previousFavoriteIds;
        console.error("Ошибка добавления в избранное", e);
        return false;
      } finally {
        this.loading = false;
      }
    },

    async removeRestaurantFromFavorites(id) {
      const authStore = useAuthStore();
      const normalizedId = normalizeRestaurantId(id);

      if (!authStore.isAuthenticated) {
        return false;
      }

      const previousFavorites = [...this.favorites];
      const previousFavoriteIds = new Set(this.favoriteIds);

      this.removeFavoriteLocally(normalizedId);

      try {
        this.loading = true;
        const { data } = await restaurantsApi.removeFromFavorites(normalizedId);
        console.log(data.message);
        return true;
      } catch (e) {
        if (e.response?.status === 404) {
          console.warn(e.response.data.error);
          this.favoritesLoaded = false;
          return true;
        }

        this.favorites = previousFavorites;
        this.favoriteIds = previousFavoriteIds;
        console.error("Ошибка удаления из избранного", e);
        return false;
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

    async fetchTopUserRestaurants() {
      try {
        this.loading = true;
        const { data } = await restaurantsApi.getTopUserRestaurants();
        this.topUserRestaurants = data;
      } catch (e) {
        console.error("Ошибка загрузки топ-ресторанов пользователя", e);
      } finally {
        this.loading = false;
      }
    },

    async fetchRecommendationRestaurants() {
      const authStore = useAuthStore();

      try {
        this.loading = true;

        if (authStore.isAuthenticated) {
          const { data } = await restaurantsApi.getRecommendations();
          this.recommendationRestaurants = data;
          return data;
        }

        const { data } = await restaurantsApi.getTopRestaurants();
        this.recommendationRestaurants = data;
        return data;
      } catch (e) {
        console.error("Ошибка загрузки рекомендаций", e);
        this.recommendationRestaurants = [];
        return [];
      } finally {
        this.loading = false;
      }
    },
  },
});
