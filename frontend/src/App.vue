<template>
  <v-app>
    <HeaderComponent />
    <v-main>
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
