<template>
  <div
    class="restaurant-card"
    :style="{
      backgroundImage: `url(${image})`,
      width: `${props.width}px`,
      height: `${props.height}px`,
    }"
  >
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

const props = defineProps({
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
  console.log(`Переход к ресторану: ${props.name}`);
}
</script>

<style>
.restaurant-card {
  position: relative;
  background-size: cover;
  background-position: center;
  border-radius: 0;
  cursor: pointer;
  overflow: hidden;
  display: flex;
  justify-content: center;
  align-items: center;
  transition: transform 0.3s ease;
}

.restaurant-card:hover {
  transform: scale(1.02);
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
