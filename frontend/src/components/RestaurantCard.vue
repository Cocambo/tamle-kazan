<template>
  <div
    class="restaurant-card"
    :style="{
      width: `${width}px`,
      height: `${height}px`,
    }"
  >
    <img class="card-image" :src="image" @error="onImageError" />

    <div class="overlay"></div>

    <div class="card-content pa-4">
      <v-btn
        icon
        variant="text"
        class="favorite-btn"
        @click.stop="toggleFavorite"
      >
        <v-icon color="white" size="30">
          {{ isFavorite ? "mdi-heart" : "mdi-heart-outline" }}
        </v-icon>
      </v-btn>

      <div class="card-bottom">
        <div class="restaurant-name">{{ name }}</div>
        <v-btn
          icon
          variant="text"
          color="white"
          class="arrow-btn"
          @click.stop="goToRestaurant"
        >
          <v-icon size="30">mdi-arrow-right</v-icon>
        </v-btn>
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref } from "vue";
import { useRouter } from "vue-router";
import fallbackImg from "@/assets/no-image.png";

defineProps({
  name: String,
  image: String,
  id: [Number, String],
  width: { type: [Number, String], default: 300 },
  height: { type: [Number, String], default: 400 },
});

const isFavorite = ref(false);
const router = useRouter();

function toggleFavorite() {
  isFavorite.value = !isFavorite.value;
}

function goToRestaurant() {
  router.push(`/restaurants/${id}`);
}

function onImageError(e) {
  e.target.src = fallbackImg;
}
</script>

<style>
.restaurant-card {
  position: relative;
  overflow: hidden;
  cursor: pointer;
  transition: transform 0.3s ease;
}

.restaurant-card:hover {
  transform: scale(1.02);
}

.card-image {
  position: absolute;
  inset: 0;
  width: 100%;
  height: 100%;
  object-fit: cover;
}

.overlay {
  position: absolute;
  inset: 0;
  background: rgba(0, 0, 0, 0.35);
  z-index: 1;
}

.card-content {
  position: relative;
  z-index: 2;
  width: 100%;
  height: 100%;
  display: flex;
  flex-direction: column;
  justify-content: space-between;
}

.favorite-btn {
  align-self: flex-end;
}

.card-bottom {
  display: flex;
  justify-content: space-between;
  align-items: center;
  color: white;
}

.restaurant-name {
  font-family: "Cormorant Garamond", serif;
  font-size: 32px;
  font-weight: 600;
  text-shadow: 0 2px 4px rgba(0, 0, 0, 0.6);
  padding-left: 12px;
}

.arrow-btn {
  background: transparent;
}
</style>
