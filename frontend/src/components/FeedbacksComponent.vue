<template>
  <section class="feedback">
    <div class="feedback__content">
      <VAvatar size="90">
        <img src="@/assets/cat1.png" />
      </VAvatar>
      <p class="feedback__text">
        {{ review.comment }}
      </p>

      <VRating
        :model-value="review.rating"
        color="accent"
        density="comfortable"
        readonly
      />

      <h2 class="feedback__author">{{ fullName }}</h2>
    </div>
  </section>
</template>

<script setup>
import { onMounted, computed } from "vue";
import { useUsersStore } from "@/stores/usersStore";

const props = defineProps({
  review: Object,
});

const usersStore = useUsersStore();

onMounted(() => {
  usersStore.fetchUser(props.review.user_id);
});

console.log(usersStore.users);

const fullName = computed(() => {
  const user = usersStore.getUserById(props.review.user_id);
  if (!user) return "Загрузка...";
  return `${user.first_name} ${user.last_name}`;
});
</script>

<style scoped>
.feedback {
  display: flex;
  justify-content: center;
  text-align: center;
  max-width: 700px;
  margin: auto;
}

.feedback__content {
  display: flex;
  flex-direction: column;
  align-items: center;
  gap: 16px;
}

.feedback__text {
  font-size: 16px;
  line-height: 1.6;
  max-width: 600px;
}

.feedback__author {
  color: white;
  font-size: 26px;
  font-weight: 600;
}
</style>
