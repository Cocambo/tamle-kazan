import { fileURLToPath, URL } from 'node:url'
import { defineConfig } from 'vite'
import vue from '@vitejs/plugin-vue'
import vuetify from 'vite-plugin-vuetify'

// https://vitejs.dev/config/
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
    },

    server: {
        host: true,          // позволяет подключаться извне (через localhost)
        port: 5173,          // явно указываем порт
        watch: {
            usePolling: true,  // нужно для Docker + Windows/WSL, чтобы hot reload работал
            interval: 100,     // (опционально) уменьшает задержку реакции на изменения
        },
    },
})
