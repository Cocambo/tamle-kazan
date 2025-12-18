<template>
  <v-app>
    <HeaderComponent v-if="!route.meta.hideHeader && route.name !== 'Home'" />
    <v-main :class="{ 'pa-0': route.meta.hideHeader }">
      <RouterView />
    </v-main>
    <FooterComponent v-if="!route.meta.hideFooter" />
  </v-app>
</template>

<script setup>
import { useRoute } from "vue-router";
import FooterComponent from "@components/FooterComponent.vue";
import HeaderComponent from "@components/HeaderComponent.vue";

const route = useRoute();
</script>

<style>
main {
  background-image: url("@/assets/branch.png");
  background-repeat: no-repeat;
  background-position: left bottom;
  background-size: auto;
}
</style>
