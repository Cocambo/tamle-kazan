<template>
  <v-app-bar flat absolute>
    <v-container
      class="d-flex align-center justify-space-between header pr-10 pl-10"
      fluid
    >
      <RouterLink to="/" class="text-decoration-none logo-text">
        TamleKazan
      </RouterLink>

      <div class="d-flex align-center nav-links">
        <RouterLink to="/restaurants" class="nav-link"> Рестораны </RouterLink>
        <RouterLink to="/favourite" class="nav-link ml-6">
          Хочу пойти
        </RouterLink>
      </div>

      <v-btn
        variant="outlined"
        color="white"
        class="register-btn text-none"
        rounded="0"
        to="/login"
      >
        Зарегистрироваться
      </v-btn>
    </v-container>
  </v-app-bar>
</template>

<style>
.header {
  background-color: rgb(var(--v-theme-primary));
}

.logo-text {
  font-size: 28px;
  font-weight: 600;
  font-family: "Cormorant Garamond", serif;
  color: white;
}

.nav-links {
  font-size: 16px;
  font-weight: 300;
  color: white;
}

.nav-link {
  text-decoration: none;
  transition: opacity 0.3s;
  color: white;
}

.nav-link:hover {
  opacity: 0.7;
}

.register-btn {
  border: 1px solid white;
  font-size: 15px;
  text-transform: none;
  letter-spacing: 0;
}
</style>
