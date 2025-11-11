<template>
  <v-footer class="footer">
    <v-container class="footer-content" fluid>
      <v-row align="center">
        <v-col md="4" class="footer-left">
          <div class="footer-title">
            <span>Tamle</span><br />
            <span>Kazan</span>
          </div>
        </v-col>

        <v-col cmd="4" class="footer-center">
          <v-list class="footer-links">
            <v-list-item>
              <RouterLink to="/" class="footer-link">Главная</RouterLink>
            </v-list-item>
            <v-list-item>
              <RouterLink to="/restaurants" class="footer-link"
                >Рестораны</RouterLink
              >
            </v-list-item>
            <v-list-item>
              <RouterLink to="/favourite" class="footer-link"
                >Хочу пойти</RouterLink
              >
            </v-list-item>
          </v-list>
        </v-col>

        <v-col md="4" class="footer-right">
          <div class="footer-contacts">
            <div class="footer-contact-title">Контакты</div>
            <div>8 917 220 86 66</div>
            <div class="footer-link">info@tamlekzn.com</div>
          </div>
          <img
            src="@/assets/rosemary.png"
            alt="Декор ветка"
            class="footer-branch"
          />
        </v-col>
      </v-row>
    </v-container>
  </v-footer>
</template>

<style scoped>
.footer {
  background-color: rgb(var(--v-theme-primary));
  color: #ffffff;
  position: relative;
  overflow: hidden;
  padding: 20px 0;
  max-height: 160px;
}

.footer-content {
  position: relative;
  max-width: 1200px;
  margin: 0 auto;
}

.footer-title {
  font-family: "Cormorant Garamond", serif;
  font-size: 36px;
  font-weight: 600;
  line-height: 1.2;
}

.footer-links {
  background: transparent;
  padding: 0;
  margin: 0;
}

.footer-links .v-list-item {
  min-height: 30px;
  padding: 0;
  font-weight: 300;
}

.footer-link {
  text-decoration: none;
  color: #ffffff;
  font-size: 16px;
  transition: opacity 0.3s;
}

.footer-link:hover {
  opacity: 0.8;
}

.footer-contact-title {
  font-weight: 600;
  margin-bottom: 5px;
}

.footer-branch {
  position: absolute;
  right: -160px;
  height: 160px;
  width: auto;
}

.footer-center,
.footer-right,
.footer-left {
  display: flex;
  flex-direction: column;
  justify-content: center;
}
</style>
