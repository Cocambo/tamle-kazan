<template>
  <v-container class="register d-flex align-center justify-center fill-height">
    <v-card class="register-card pa-8" elevation="3" rounded="0">
      <h1 class="text-center register-title mb-6">Регистрация</h1>

      <v-form @submit.prevent="register">
        <v-text-field
          v-model="firstName"
          variant="outlined"
          label="Введите имя"
          class="mb-4"
          hide-details
          rounded="0"
        />

        <v-text-field
          v-model="lastName"
          variant="outlined"
          label="Введите фамилию"
          class="mb-4"
          hide-details
          rounded="0"
        />

        <v-text-field
          v-model="email"
          type="email"
          variant="outlined"
          label="Введите почту"
          class="mb-4"
          hide-details
          rounded="0"
        />

        <v-text-field
          v-model="password"
          :type="showPassword ? 'text' : 'password'"
          variant="outlined"
          label="Введите пароль"
          class="mb-4"
          hide-details
          rounded="0"
        >
          <template #append-inner>
            <v-icon
              color="primary"
              @click="showPassword = !showPassword"
              style="cursor: pointer"
            >
              {{ showPassword ? "mdi-eye-off" : "mdi-eye" }}
            </v-icon>
          </template>
        </v-text-field>

        <v-text-field
          v-model="confirmPassword"
          :type="showConfirmPassword ? 'text' : 'password'"
          variant="outlined"
          label="Повторите пароль"
          class="mb-2"
          hide-details
          rounded="0"
        >
          <template #append-inner>
            <v-icon
              color="primary"
              @click="showConfirmPassword = !showConfirmPassword"
              style="cursor: pointer"
            >
              {{ showConfirmPassword ? "mdi-eye-off" : "mdi-eye" }}
            </v-icon>
          </template>
        </v-text-field>

        <div class="d-flex justify-end mb-4">
          <RouterLink to="/auth" class="text-link"
            >Есть аккаунт? Войти</RouterLink
          >
        </div>

        <div class="text-center">
          <v-btn color="primary" class="register-btn" type="submit" rounded="0">
            Зарегистрироваться
          </v-btn>
        </div>
      </v-form>
    </v-card>

    <v-snackbar v-model="snackbar" :color="snackbarColor" timeout="3000">
      {{ snackbarText }}
    </v-snackbar>
  </v-container>
</template>

<script setup>
import { ref } from "vue";

const firstName = ref("");
const lastName = ref("");
const email = ref("");
const password = ref("");
const confirmPassword = ref("");

const showPassword = ref(false);
const showConfirmPassword = ref(false);

const snackbar = ref(false);
const snackbarText = ref("");
const snackbarColor = ref("");

const emailRegex = /^[^\s@]+@[^\s@]+\.[^\s@]+$/;

const showToast = (message, color = "warning") => {
  snackbarText.value = message;
  snackbarColor.value = color;
  snackbar.value = true;
};

const register = () => {
  if (
    !firstName.value ||
    !lastName.value ||
    !email.value ||
    !password.value ||
    !confirmPassword.value
  ) {
    showToast("Заполните все поля", "warning");
    return;
  }

  if (!emailRegex.test(email.value)) {
    showToast("Укажите корректную почту", "warning");
    return;
  }

  if (password.value !== confirmPassword.value) {
    showToast("Пароли не совпадают", "warning");
    return;
  }

  showToast("Регистрация прошла успешно!", "success");

  console.log("Регистрация:", { firstName, lastName, email, password });
};
</script>

<style scoped>
.register {
  background-image: url("@/assets/branch.png");
  background-repeat: no-repeat;
  background-position: left bottom;
  background-size: auto;
  min-height: 100vh;
  background-color: #ffffff;
}

.register-card {
  border: 1px solid rgb(var(--v-theme-primary));
  max-width: 700px;
}

.register-title {
  color: rgb(var(--v-theme-primary));
  font-weight: 700;
  font-size: 40px;
}

.v-text-field {
  width: 400px;
  background-color: rgba(var(--v-theme-primary), 0.15);
  color: rgb(var(--v-theme-primary));
}

.v-text-field input {
  color: rgb(var(--v-theme-primary));
}

.register-btn {
  background-color: rgb(var(--v-theme-primary));
  color: white;
  padding: 10px 20px;
  font-weight: 600;
  font-size: 16px;
  font-family: "Cormorant Garamond", serif;
}

.text-link {
  text-decoration: none;
  font-size: 14px;
}

.text-link:hover {
  text-decoration: none;
}
</style>
