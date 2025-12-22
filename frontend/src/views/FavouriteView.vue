<template>
  <v-container class="restaurants-page">
    <h1 class="page-title">Хочу пойти</h1>
    <p class="page-subtitle">
      Ваши избранные заведения собраны в одном месте. Возвращайтесь к
      проверенным вкусам и открывайте новые возможности для гастрономических
      открытий.
    </p>

    <ThreeRestaurantsComponent
      :restaurants="store.favorites"
      :card-width="300"
      :card-height="400"
    />
  </v-container>
</template>

<script setup>
import { onMounted } from "vue";
import { useRestaurantsStore } from "@/stores/restaurantsStore";
import ThreeRestaurantsComponent from "@/components/ThreeRestaurantsComponent.vue";

const store = useRestaurantsStore();

onMounted(() => {
  store.fetchFavorites();
});
</script>

<style>
.restaurants-page {
  padding-top: 40px;
}
.page-title {
  font-size: 48px;
  font-weight: 600;
  margin-bottom: 8px;
}
.page-subtitle {
  font-size: 16px;
  color: #555;
  margin-bottom: 32px;
}
.filters {
  display: grid;
  grid-template-columns: 1fr 1fr 1fr 2fr;
  gap: 18px;
  background: #eef4e4;
  padding: 20px;
  margin-bottom: 40px;
}
</style>
