import { defineStore } from "pinia";
import { reviewsApi } from "@/api/reviews.api";

export const useReviewsStore = defineStore("reviews", {
    state: () => ({
        reviews: [],
        loading: false,
    }),

    actions: {
        async fetchRestaurantReviews(restaurantId) {
            try {
                this.loading = true;
                const { data } = await reviewsApi.getRestaurantReviews(restaurantId);
                this.reviews = data;
            } catch (e) {
                console.error("Ошибка загрузки отзывов", e);
            } finally {
                this.loading = false;
            }
        },
        async createRestaurantReview(restaurantId, review) {
            try {
                this.loading = true;
                const { status } = await reviewsApi.createRestaurantsReview(restaurantId, review);

                if (status === 201) {
                    await this.reviews.fetchRestaurantReviews(restaurantId);
                    console.log("Отзыв успешно добавлен")
                }

            } catch (e) {
                if (e.response?.status === 404) {
                    console.warn(e.response.data.error);
                } else if (e.response?.status === 409) {
                    console.warn(e.response.data.error);
                } else {
                    console.error("Ошибка создания отзыва", e);
                }

            } finally {
                this.loading = false;
            }
        }
    },
});
