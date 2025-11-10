<template>
  <v-card class="restaurant-card" flat rounded="xl">
    <v-img :src="image" class="restaurant-image" height="350" cover>
      <v-btn
        icon
        size="small"
        class="favorite-btn"
        @click.stop="toggleFavorite"
      >
        <v-icon :color="isFavorite ? 'red' : 'white'">
          {{ isFavorite ? "mdi-heart" : "mdi-heart-outline" }}
        </v-icon>
      </v-btn>

      <div class="card-bottom">
        <div class="restaurant-name">{{ name }}</div>
        <v-btn
          icon
          size="small"
          color="white"
          class="arrow-btn"
          @click="goToRestaurant"
        >
          <v-icon>mdi-arrow-right</v-icon>
        </v-btn>
      </div>
    </v-img>
  </v-card>
</template>

<script setup>
import { ref } from "vue";
import { useRouter } from "vue-router";

const props = defineProps({
  name: String,
  image: String,
  id: [Number, String],
});

const isFavorite = ref(false);
const router = useRouter();

function toggleFavorite() {
  isFavorite.value = !isFavorite.value;
}

function goToRestaurant() {
  // Можно позже заменить на реальный роут
  console.log(`Переход к ресторану: ${props.name}`);
  // router.push(`/restaurant/${props.id}`);
}
</script>

<style scoped>
.restaurant-card {
  position: relative;
  overflow: hidden;
}

.restaurant-image {
  position: relative;
  border-radius: 20px;
}

.favorite-btn {
  position: absolute;
  top: 10px;
  right: 10px;
  z-index: 2;
}

.card-bottom {
  position: absolute;
  bottom: 10px;
  left: 20px;
  right: 20px;
  display: flex;
  justify-content: space-between;
  align-items: center;
  color: white;
}

.restaurant-name {
  font-family: "Cormorant Garamond", serif;
  font-size: 24px;
  font-weight: 600;
}

.arrow-btn {
  background: transparent;
}
</style>
