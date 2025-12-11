<template>
  <section
    class="hero-section d-flex flex-column justify-center align-start text-start"
  >
    <div class="overlay"></div>

    <div class="hero-content align-start">
      <h1 class="hero-title">
        Ваш <br />
        гастрономический <br />компас в Казани
      </h1>
      <p class="hero-subtitle">Воплощаем искусство комфортной гастрономии</p>
    </div>
  </section>
  <TopRestaurantsComponent />
  <RecomendationBlockComponent />
</template>

<script setup>
import TopRestaurantsComponent from "@components/TopRestaurantsComponent.vue";
import RecomendationBlockComponent from "@components/RecomendationBlockComponent.vue";
</script>

<style>
.hero-section {
  position: relative;
  height: 100vh;
  width: 100%;
  background-image: url("@/assets/restaurant.png");
  background-size: cover;
  background-position: center;
  background-repeat: no-repeat;
  color: white;
  overflow: hidden;
}

.overlay {
  position: absolute;
  inset: 0;
  background: rgba(0, 0, 0, 0.4);
  z-index: 1;
}

.hero-content {
  position: relative;
  z-index: 2;
  max-width: 800px;
  padding: 0px 80px;
  margin-top: 0px;
}

.hero-title {
  font-family: "Cormorant Garamond", serif;
  font-weight: 700;
  font-size: 75px;
  line-height: 1.2;
  margin-bottom: 20px;
  color: white;
}

.hero-subtitle {
  font-size: 24px;
  font-weight: 500;
  font-family: "Cormorant Garamond", serif;
}
</style>
