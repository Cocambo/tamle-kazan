<template>
  <VContainer fluid class="profile pa-0 d-flex justify-center">
    <VCard class="profile-card" rounded="0">
      <div
        class="profile-card__header h-40 ma-0 pa-16 d-flex align-center justify-space-between"
      >
        <div class="profile-photo"></div>
        <div class="profile-name">
          <h1 class="profile-name__name">Мазитова Алина</h1>
          <p class="profile-name__status">Ценитель прекрасного</p>
          <div class="profile-email d-flex align-center">
            <p class="profile-email__email">mazitova1402@gmail.com</p>
            <VBtn
              class="profile-email__btn"
              size="sm"
              rounded="0"
              color="primary"
              >Подтвердить</VBtn
            >
          </div>
        </div>
      </div>
      <div class="profile-restaurants pa-8 justify-center">
        <h1 class="text-center">Топ ваших ресторанов</h1>
        <ThreeRestaurantsComponent :cardWidth="230" :cardHeight="300" />
      </div>
      <div class="profile-logout d-flex justify-center align-center">
        <VBtn class="profile-logout__btn" rounded="0" color="primary"
          >Выйти</VBtn
        >
      </div>
    </VCard>
  </VContainer>
</template>

<script setup>
import ThreeRestaurantsComponent from "@/components/ThreeRestaurantsComponent.vue";
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
