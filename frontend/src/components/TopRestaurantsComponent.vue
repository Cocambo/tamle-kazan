<template>
  <section class="top-restaurants-section">
    <div class="title-container">
      <h2 class="section-title">Топ 3 ресторана</h2>
      <p class="section-subtitle">Выбор наших клиентов</p>
    </div>

    <ThreeRestaurantsComponent />

    <div class="button-container">
      <v-btn color="primary" class="all-btn" rounded="0" to="/restaurants">
        Все рестораны
      </v-btn>
    </div>

    <img src="@/assets/branch.png" alt="decor branch" class="decor-branch" />
  </section>
</template>

<script setup>
import ThreeRestaurantsComponent from "@/components/ThreeRestaurantsComponent.vue";
</script>

<style>
.top-restaurants-section {
  position: relative;
  padding: 30px 0 80px;
  background-color: #fff;
  overflow: hidden;
}

.title-container {
  text-align: left;
  margin-left: 60px;
  margin-bottom: 30px;
}

.section-title {
  font-family: "Cormorant Garamond", serif;
  font-weight: 700;
  font-size: 64px;
  margin-bottom: 10px;
  color: #000;
}

.section-subtitle {
  color: #4d4d4d;
  font-size: 18px;
}

.cards-container {
  max-width: 1200px;
  margin: 0 auto;
  gap: 24px;
}

.card-wrapper {
  flex: 1 1 30%;
  display: flex;
  justify-content: center;
}

.button-container {
  text-align: center;
  margin-top: 60px;
}

.all-btn {
  background-color: rgb(var(--v-theme-primary));
  color: white;
  padding: 10px 20px;
  font-weight: 600;
  font-size: 16px;
  font-family: "Cormorant Garamond", serif;
}

.decor-branch {
  position: absolute;
  top: 0;
  right: 0;
  width: 600px;
  opacity: 0.9;
  pointer-events: none;
  user-select: none;
  z-index: 0;
  transform: scaleX(-1);
  transform: rotate(180deg);
}
</style>
