<template>
  <section class="feedbacks-section flex-column align-center text-center">
    <div class="feedbacks__overlay"></div>

    <div class="feedbacks__content">
      <h1 class="feedbacks-content__title">Отзывы</h1>

      <VCarousel
        v-model="activeSlide"
        class="feedbacks-carousel"
        height="450"
        hide-delimiter-background
        :show-arrows="false"
      >
        <v-carousel-item v-for="review in goodReviews" :key="review.id">
          <FeedbacksComponent :review="review" />
        </v-carousel-item>
      </VCarousel>
    </div>
  </section>
</template>

<script setup>
import { ref, watch, computed } from "vue";
import { useReviewsStore } from "@/stores/reviewsStore";
import FeedbacksComponent from "@/components/FeedbacksComponent.vue";

const reviewsStore = useReviewsStore();
const activeSlide = ref(0);

const goodReviews = computed(() => {
  return reviewsStore.reviews
    .filter((r) => r.rating >= 2)
    .slice(-5)
    .reverse();
});

watch(
  () => goodReviews.value.length,
  (newLength, oldLength) => {
    if (newLength > oldLength) {
      activeSlide.value = 0;
    }
  },
);
</script>

<style>
.feedbacks-section {
  position: relative;
  width: 100%;
  height: 600px;
  padding: 35px 0;

  background-image: url("@/assets/olio.jpg");
  background-size: cover;
  background-position: center;
  background-repeat: no-repeat;

  color: white;
  overflow: hidden;
}

.feedbacks__overlay {
  position: absolute;
  inset: 0;
  background: rgba(35, 48, 0, 0.8);
  z-index: 1;
}

.feedbacks__content {
  position: relative;
  z-index: 2;
  width: 100%;
}

.feedbacks-content__title {
  color: white;
}

h1 {
  font-size: 48px;
  margin-bottom: 40px;
}

.feedbacks-carousel .v-carousel__controls .v-btn {
  width: 10px;
  height: 10px;
  min-width: 0;
  padding: 0;

  border-radius: 50%;
  border: 2px solid white;
  background: transparent;

  box-shadow: none;
  overflow: hidden;
}

.feedbacks-carousel .v-carousel__controls .v-btn .v-ripple__container {
  display: none;
}

.feedbacks-carousel .v-carousel__controls .v-btn--active {
  background: white;
  border-color: white;
}
</style>
