<template>
  <VContainer fluid class="profile pa-0 d-flex justify-center">
    <VCard class="profile-card" rounded="0">
      <div
        class="profile-card__header h-40 ma-0 pa-16 d-flex align-center justify-space-between"
      >
        <div class="profile-photo"></div>
        <div class="profile-name">
          <h1 class="profile-name__name">
            {{ user?.first_name }} {{ user?.last_name }}
          </h1>
          <p class="profile-name__status">Ценитель прекрасного</p>
          <div class="profile-email d-flex align-center">
            <p class="profile-email__email">{{ user?.email }}</p>
            <VBtn
              v-if="!authStore.user?.is_email_confirmed"
              @click="resendConfirmation"
              class="profile-email__btn"
              size="sm"
              rounded="0"
              color="primary"
              >Подтвердить</VBtn
            >
            <span v-else> Почта подтверждена </span>
          </div>
        </div>
      </div>
      <div class="profile-restaurants pa-8 justify-center">
        <h1 class="text-center">Топ ваших ресторанов</h1>
        <ThreeRestaurantsComponent :cardWidth="230" :cardHeight="300" />
      </div>
      <div class="profile-logout d-flex justify-center align-center">
        <VBtn
          class="profile-logout__btn"
          rounded="0"
          color="primary"
          @click="logout"
          >Выйти</VBtn
        >
      </div>
    </VCard>
    <VSnackbar v-model="snackbar" :color="snackbarColor" timeout="3000">
      {{ snackbarText }}
    </VSnackbar>
  </VContainer>
</template>

<script setup>
import { computed, ref } from "vue";
import { VSnackbar } from "vuetify/components";
import { useAuthStore } from "@/stores/authStore";
import { useRouter } from "vue-router";
import ThreeRestaurantsComponent from "@/components/ThreeRestaurantsComponent.vue";

const authStore = useAuthStore();
const router = useRouter();
const user = computed(() => authStore.user);

const snackbar = ref(false);
const snackbarText = ref("");
const snackbarColor = ref("");

const showToast = (message, color = "warning") => {
  snackbarText.value = message;
  snackbarColor.value = color;
  snackbar.value = true;
};

const logout = async () => {
  await authStore.logout();
  router.push("/auth");
};

const resendConfirmation = async () => {
  try {
    await authStore.resendEmailConfirmation();
    showToast("Письмо подтверждения отправлено на почту", "success");
    return;
  } catch {
    showToast("Ошибка отправки письма", "warning");
    return;
  }
};
</script>

<style scoped>
.profile {
  margin: 40px auto;
}

.profile-card {
  border: 2px solid rgb(var(--v-theme-primary));
  max-width: 900px;
  width: 100%;
}

.profile-card__header {
  border-radius: 0px;
  background-color: #ebf0e4;
}

.profile-photo {
  background-size: cover;
  background-position: center;
  background-image: url("@/assets/cat1.png");
  border: 3px solid rgb(var(--v-theme-primary));
  width: 180px;
  height: 180px;
  margin-left: 20px;
}
.profile-name {
  padding-right: 20px;
}

.profile-name__name {
  font-size: 38px;
}

.profile-name__status {
  color: rgb(var(--v-theme-secondary));
}

.profile-email {
  margin-top: 16px;
  display: flex;
  align-items: center;
  gap: 14px;
}

.profile-email__btn {
  font-size: 14px;
  font-family: "Cormorant Garamond", serif;
  padding: 6px 12px;
  border: none;
  border-radius: 0;
}

.profile-logout {
  height: 70px;
  background-color: #ebf0e4;
}

.profile-logout__btn {
  font-family: "Cormorant Garamond", serif;
}
</style>
