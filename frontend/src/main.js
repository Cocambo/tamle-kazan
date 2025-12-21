import { createApp } from 'vue'
import { createPinia } from 'pinia'
import App from './App.vue'
import router from './router/router'

import { createVuetify } from 'vuetify'

import '@mdi/font/css/materialdesignicons.css'
import '@/style.css'
import 'vuetify/styles'
import * as components from 'vuetify/components'
import * as directives from 'vuetify/directives'


const vuetify = createVuetify({
  components,
  directives,
  icons: {
    defaultSet: 'mdi',
  },
  theme: {
    defaultTheme: 'light',
    themes: {
      light: {
        dark: false,
        colors: {
          primary: '#233000',
          secondary: '#5E6600',
          accent: '#9CAA00',
          text: '#000000',
          link: '#233000',

          'on-primary': '#FFFFFF',
        },
      },
    },
  },
})

const app = createApp(App)
app.use(vuetify)
app.use(createPinia())
app.use(router)
app.mount('#app')
