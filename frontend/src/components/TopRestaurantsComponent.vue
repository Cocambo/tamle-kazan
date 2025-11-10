<template>
  <section class="top-restaurants-section">
    <div class="title-container">
      <h2 class="section-title">Топ 3 ресторана</h2>
      <p class="section-subtitle">Выбор наших клиентов</p>
    </div>

    <v-container fluid class="cards-container">
      <v-row justify="center" align="stretch" dense>
        <v-col
          v-for="(restaurant, index) in restaurants"
          :key="index"
          cols="12"
          md="4"
        >
          <RestaurantCard
            :name="restaurant.name"
            :image="restaurant.image"
            :id="restaurant.id"
          />
        </v-col>
      </v-row>

      <div class="button-container">
        <v-btn color="olive" class="all-btn" rounded="lg" variant="flat">
          Все рестораны
        </v-btn>
      </div>
    </v-container>
  </section>
</template>

<script setup>
import RestaurantCard from "@/components/RestaurantCard.vue";

const restaurants = [
  { id: 1, name: "Olio", image: "@/assets/olio.jpg" },
  { id: 2, name: "Dia", image: "@/assets/dia.jpg" },
  { id: 3, name: "Terra et Silva", image: "@/assets/terra.jpg" },
];
</script>

<style scoped>
.top-restaurants-section {
  padding: 80px 0;
  background-color: #fff;
  position: relative;
}

.title-container {
  text-align: left;
  margin-left: 60px;
  margin-bottom: 40px;
}

.section-title {
  font-family: "Cormorant Garamond", serif;
  font-weight: 700;
  font-size: 56px;
  margin-bottom: 10px;
  color: #1a1a1a;
}

.section-subtitle {
  color: #666;
  font-size: 18px;
  font-family: "Cormorant Garamond", serif;
}

.cards-container {
  max-width: 1200px;
  margin: 0 auto;
}

.button-container {
  text-align: center;
  margin-top: 40px;
}

.all-btn {
  background-color: #3b4a1d !important;
  color: white !important;
  padding: 12px 32px;
  font-size: 18px;
  font-family: "Cormorant Garamond", serif;
}
</style>
