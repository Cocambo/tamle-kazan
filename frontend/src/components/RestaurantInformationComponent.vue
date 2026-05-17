<template>
  <div v-if="restaurant" class="restaurant__header justify-center">
    <img src="../assets/Tomato.png" class="decor tomato" />
    <img src="../assets/Avocado.png" class="decor avocado" />
    <div class="image-wrapper">
      <img class="restaurant__image" :src="mainPhoto" :alt="restaurant.name" />
    </div>

    <div class="restaurant__text">
      <h1 class="title">{{ restaurant.name }}</h1>
      <p class="type">{{ restaurant.cuisine }} кухня</p>

      <p class="description">
        {{ restaurant.description }}
      </p>

      <div class="rating">
        <span class="rating-text">Рейтинг:</span>
        <v-rating
          v-if="restaurant.rating > 0"
          :model-value="Number(restaurant.rating)"
          length="5"
          readonly
          color="accent"
          empty-icon="mdi-star-outline"
          full-icon="mdi-star"
          half-increments
          size="28"
        />

        <span v-else class="no-rating"> Нет оценок </span>
      </div>

      <p class="price">Средний чек: от {{ restaurant.average_bill }}₽</p>
      <p class="address">{{ restaurant.address }}</p>
    </div>
  </div>
</template>

<script setup>
import { computed } from "vue";
import { useRestaurantsStore } from "@/stores/restaurantsStore";
import { getMainPhoto } from "@/utils/getMainPhoto";

const restaurantsStore = useRestaurantsStore();

const restaurant = computed(() => restaurantsStore.currentRestaurant);

const mainPhoto = computed(() => {
  return getMainPhoto(restaurant.value);
});
</script>

<style scoped>
.restaurant__header {
  display: flex;
  align-items: center;
  gap: 70px;
  padding: 80px 80px;
  position: relative;
}

.decor {
  position: absolute;
  pointer-events: none;
  opacity: 0.8;
}

.tomato {
  top: 0;
  right: 0;
  width: 250px;
}

.avocado {
  bottom: 0;
  left: 0;
  width: 250px;
}

.image-wrapper {
  margin-top: 15px;
  border: 6px solid rgb(var(--v-theme-primary));
  width: 300px;
  height: 400px;
  overflow: hidden;
}

.restaurant__image {
  width: 100%;
  height: 100%;
  display: block;
  object-fit: cover;
  object-position: center;
}

.restaurant__text {
  max-width: 520px;
}

.title {
  font-size: 80px;
  font-weight: 600;
  color: rgb(var(--v-theme-primary));
  margin-bottom: 0px;
}

.type {
  color: #666;
  font-size: 20px;
  margin-bottom: 20px;
}

.description {
  font-size: 18px;
  line-height: 1.6;
  color: rgb(var(--v-theme-primary));
  margin-bottom: 20px;
}

.rating {
  display: flex;
  align-items: center;
  gap: 12px;
  margin-bottom: 10px;
}

.rating-text {
  font-size: 18px;
}

.no-rating {
  color: #666;
  font-size: 18px;
}

.price {
  font-size: 18px;
  margin-bottom: 10px;
}

.address {
  color: #666;
  font-size: 18px;
}
</style>
