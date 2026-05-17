<template>
  <v-container class="feedback-form d-flex justify-center align-center">
    <VCard class="feedback-form__card pa-8" elevation="3" rounded="0">
      <h1 class="feedback-form__title text-center mb-6">Оставьте отзыв</h1>

      <VForm @submit.prevent="leave">
        <div class="rating-block mb-4">
          <span class="rating-label">Оценка:</span>
          <VRating
            v-model="mark"
            length="5"
            color="primary"
            density="comfortable"
          />
        </div>

        <VTextarea
          v-model="feedback"
          label="Комментарий"
          variant="outlined"
          rounded="0"
          rows="3"
          hide-details
          class="mb-6"
        />

        <div class="text-center">
          <VBtn
            class="register-btn"
            type="submit"
            rounded="0"
            :loading="reviewsStore.loading"
          >
            Отправить
          </VBtn>
        </div>
      </VForm>
    </VCard>
    <v-snackbar v-model="snackbar" :color="snackbarColor" timeout="3000">
      {{ snackbarText }}
    </v-snackbar>
  </v-container>
</template>

<script setup>
import { ref } from "vue";
import { useRoute } from "vue-router";
import { useReviewsStore } from "@/stores/reviewsStore";
import { useRestaurantsStore } from "@/stores/restaurantsStore";

const restaurantsStore = useRestaurantsStore();

const reviewsStore = useReviewsStore();
const route = useRoute();

const mark = ref(0);
const feedback = ref("");

const snackbar = ref(false);
const snackbarText = ref("");
const snackbarColor = ref("");

const showToast = (message, color = "warning") => {
  snackbarText.value = message;
  snackbarColor.value = color;
  snackbar.value = true;
};

const leave = async () => {
  const restaurantId = route.params.id;

  await reviewsStore.createRestaurantReview(restaurantId, {
    rating: mark.value,
    comment: feedback.value,
  });

  await reviewsStore.fetchRestaurantReviews(restaurantId);
  await restaurantsStore.fetchRestaurantById(restaurantId);

  mark.value = 0;
  feedback.value = "";

  showToast("Отзыв успешно отправлен!", "success");
};
</script>

<style scoped>
.feedback-form {
  min-height: 500px;
}

.feedback-form__card {
  border: 2px solid rgb(var(--v-theme-primary));
  width: 450px;
}

.feedback-form__title {
  color: rgb(var(--v-theme-primary));
  font-weight: 700;
  font-size: 40px;
}

.rating-block {
  display: flex;
  align-items: center;
  gap: 15px;
  background-color: rgba(var(--v-theme-primary), 0.15);
  padding: 12px;
  border: 1px solid rgba(var(--v-theme-primary), 0.4);
}

.rating-label {
  color: rgb(var(--v-theme-primary));
  font-weight: 500;
}

.v-textarea {
  background-color: rgba(var(--v-theme-primary), 0.15);
}

.register-btn {
  background-color: rgb(var(--v-theme-primary));
  color: white;
  font-weight: 600;
  font-size: 18px;
  font-family: "Cormorant Garamond", serif;
}
</style>
