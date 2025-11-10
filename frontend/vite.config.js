import { fileURLToPath, URL } from 'node:url'
import { defineConfig } from 'vite'
import vue from '@vitejs/plugin-vue'
import vuetify from 'vite-plugin-vuetify'

export default defineConfig({
  plugins: [
    vue(),
    vuetify({ autoImport: false })
  ],

  resolve: {
    alias: {
      '@': fileURLToPath(new URL('./src', import.meta.url)),
      '@components': fileURLToPath(new URL('./src/components', import.meta.url)),
      '@styles': fileURLToPath(new URL('./src/assets/styles', import.meta.url)),
    },
  },

  css: {
  preprocessorOptions: {
    scss: {
      additionalData: `
        @use "vuetify/settings" with (
          $body-font-family: 'Lato', sans-serif,
          $heading-font-family: 'Cormorant Garamond', serif
        );
      `
    }
  }
}
})
