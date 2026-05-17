<template>
  <RestaurantInformationComponent />
  <RestaurantAtmosphereComponent />
  <FeedbacksCarouselComponent />
  <LeaveFeedbackComponent />
</template>

<script setup>
import { onMounted } from "vue";
import { useRestaurantsStore } from "@/stores/restaurantsStore";
import { useReviewsStore } from "@/stores/reviewsStore";

import RestaurantInformationComponent from "@/components/RestaurantInformationComponent.vue";
import RestaurantAtmosphereComponent from "@/components/RestaurantAtmosphereComponent.vue";
import FeedbacksCarouselComponent from "@/components/FeedbacksCarouselComponent.vue";
import LeaveFeedbackComponent from "@/components/LeaveFeedbackComponent.vue";

const props = defineProps({
  id: {
    type: String,
    required: true,
  },
});

const restaurantsStore = useRestaurantsStore();
const reviewsStore = useReviewsStore();

onMounted(async () => {
  await restaurantsStore.fetchRestaurantById(props.id);
  await reviewsStore.fetchRestaurantReviews(props.id);
});
</script>
