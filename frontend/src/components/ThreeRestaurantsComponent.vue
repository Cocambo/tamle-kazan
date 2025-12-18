<template>
  <VContainer
    class="cards-container pa-8 d-flex align-center justify-space-between flex-wrap"
  >
    <div
      v-for="(restaurant, index) in restaurants"
      :key="index"
      class="card-wrapper"
    >
      <RestaurantCard
        :name="restaurant.name"
        :image="restaurant.image"
        :id="restaurant.id"
        :width="cardWidth"
        :height="cardHeight"
      />
    </div>
  </VContainer>
</template>

<script setup>
const props = defineProps({
  cardWidth: { type: [Number, String], default: 300 },
  cardHeight: { type: [Number, String], default: 400 },
});

import RestaurantCard from "@/components/RestaurantCard.vue";
import olioImg from "@/assets/olio.jpg";
import diaImg from "@/assets/dia2.jpg";
import terraImg from "@/assets/terra.jpg";

const restaurants = [
  { id: 1, name: "Olio", image: olioImg },
  { id: 2, name: "Dia", image: diaImg },
  { id: 3, name: "Terra et Silva", image: terraImg },
];
</script>
