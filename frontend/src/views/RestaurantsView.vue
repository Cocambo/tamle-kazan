<template>
  <v-container class="restaurants-page">
    <h1 class="page-title">Рестораны</h1>
    <p class="page-subtitle">
      Откройте для себя богатый выбор гастрономических заведений города.
      Используйте фильтры, чтобы быстро найти ресторан вашей мечты.
    </p>

    <div class="filters">
      <v-select
        v-model="filters.cuisine"
        :items="cuisines"
        label="Кухня"
        variant="outlined"
        hide-details
        class="filter-item"
        rounded="0"
        clearable
        @update:model-value="updateRestaurants"
      />

      <v-select
        v-model="filters.min_bill"
        :items="bills"
        label="Мин. чек"
        variant="outlined"
        hide-details
        class="filter-item"
        rounded="0"
        clearable
        @update:model-value="updateRestaurants"
      />

      <!-- <v-select
        v-model="filters.max_bill"
        :items="bills"
        label="Макс. чек"
        variant="outlined"
        hide-details
        class="filter-item"
        rounded="0"
        clearable
        @update:model-value="updateRestaurants"
      /> -->

      <v-select
        v-model="filters.min_rating"
        :items="ratings"
        label="Рейтинг"
        variant="outlined"
        hide-details
        class="filter-item"
        rounded="0"
        clearable
        @update:model-value="updateRestaurants"
      />

      <v-text-field
        v-model="filters.search"
        label="Поиск ресторана"
        variant="outlined"
        hide-details
        append-inner-icon="mdi-magnify"
        class="filter-search"
        rounded="0"
        @update:model-value="updateRestaurants"
      />
    </div>

    <ThreeRestaurantsComponent
      :restaurants="store.restaurants"
      :card-width="300"
      :card-height="400"
    />
  </v-container>
</template>

<script setup>
import { reactive, onMounted } from "vue";
import { useRestaurantsStore } from "@/stores/restaurantsStore";
import ThreeRestaurantsComponent from "@/components/ThreeRestaurantsComponent.vue";

const filters = reactive({
  search: "",
  cuisine: null,
  min_bill: null,
  max_bill: null,
  min_rating: null,
  limit: 20,
});

const cuisines = [
  "Татарская",
  "Фьюжн",
  "Итальянская",
  "Паназиатская",
  "Греческая",
  "Грузинская",
  "Индийская",
];
const bills = [500, 1000, 1500, 2000, 3000];
const ratings = [0, 4, 4.5, 5];

const store = useRestaurantsStore();

onMounted(() => {
  updateRestaurants();
});

function updateRestaurants() {
  const params = {
    search: filters.search || undefined,
    cuisine: filters.cuisine || undefined,
    min_bill: filters.min_bill || 0,
    max_bill: filters.max_bill || 10000,
    min_rating: filters.min_rating || 0,
    limit: filters.limit,
  };

  store.fetchRestaurants(params);
}
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
