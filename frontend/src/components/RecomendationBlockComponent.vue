<template>
  <section class="recommendation-section">
    <div class="recommendation-section__heading">
      <h2 class="recommendation-section__title">{{ title }}</h2>
      <p class="recommendation-section__subtitle">{{ subtitle }}</p>
    </div>

    <ThreeRestaurantsComponent
      :restaurants="restaurantsStore.recommendationRestaurants"
      :card-width="300"
      :card-height="400"
    />

    <div class="recommendation-section__actions">
      <v-btn
        color="primary"
        class="recommendation-section__button"
        rounded="0"
        to="/restaurants"
      >
        Смотреть все рестораны
      </v-btn>
    </div>
  </section>
</template>

<script setup>
import { computed, onMounted } from "vue";
import { useAuthStore } from "@/stores/authStore";
import { useRestaurantsStore } from "@/stores/restaurantsStore";
import ThreeRestaurantsComponent from "@/components/ThreeRestaurantsComponent.vue";

const authStore = useAuthStore();
const restaurantsStore = useRestaurantsStore();

const title = computed(() =>
  authStore.isAuthenticated ? "Рекомендации для вас" : "Популярные рестораны"
);

const subtitle = computed(() =>
  authStore.isAuthenticated
    ? "Подобрали рестораны по вашим избранным и отзывам. Если данных пока мало, покажем лучшие места города."
    : "Начните с самых популярных мест города. После авторизации мы сможем показывать персональные рекомендации."
);

onMounted(async () => {
  if (authStore.isAuthenticated) {
    await Promise.all([
      restaurantsStore.fetchFavorites(),
      restaurantsStore.fetchRecommendationRestaurants(),
    ]);
    return;
  }

  await restaurantsStore.fetchRecommendationRestaurants();
});
</script>

<style scoped>
.recommendation-section {
  padding: 32px 0 80px;
  background-color: #ebf0e4;
}

.recommendation-section__heading {
  text-align: center;
  margin-bottom: 32px;
}

.recommendation-section__title {
  font-family: "Cormorant Garamond", serif;
  font-size: 60px;
  font-weight: 700;
  margin-bottom: 12px;
}

.recommendation-section__subtitle {
  max-width: 760px;
  margin: 0 auto;
  font-size: 18px;
  line-height: 1.5;
  color: #4d4d4d;
}

.recommendation-section__actions {
  display: flex;
  justify-content: center;
  margin-top: 40px;
}

.recommendation-section__button {
  font-family: "Cormorant Garamond", serif;
}
</style>
