<template>
  <v-container
    class="cards-container d-flex align-center justify-space-between flex-wrap"
    fluid
  >
    <div
      v-for="restaurant in restaurants"
      :key="restaurant.id"
      class="card-wrapper"
    >
      <RestaurantCard
        :id="restaurant.id"
        :name="restaurant.name"
        :image="getImage(restaurant)"
        :width="cardWidth"
        :height="cardHeight"
      />
    </div>
  </v-container>
</template>

<script setup>
import RestaurantCard from "@/components/RestaurantCard.vue";
import { getMainPhoto } from "@/utils/getMainPhoto.js";

const props = defineProps({
  restaurants: { type: Array, required: true },
  cardWidth: { type: [Number, String], default: 300 },
  cardHeight: { type: [Number, String], default: 400 },
});

function getImage(restaurant) {
  return getMainPhoto(restaurant) || "";
}
</script>
