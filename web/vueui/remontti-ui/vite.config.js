import { defineConfig } from 'vite'
import vue from '@vitejs/plugin-vue'
import { resolve } from 'path'

// https://vitejs.dev/config/
export default defineConfig({
  plugins: [vue()],
  build: {
    rollupOptions: {
      input: {
        main: resolve(__dirname, 'index.html'),
        login: resolve(__dirname, '/login/index.html'),
        companies: resolve(__dirname, '/companies/index.html'),
        permissions: resolve(__dirname, '/permissions/index.html'),
      }
    }
  }
})
