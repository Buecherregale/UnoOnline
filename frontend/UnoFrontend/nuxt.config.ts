// https://nuxt.com/docs/api/configuration/nuxt-config
import {process} from "std-env";

export default defineNuxtConfig({
  compatibilityDate: '2024-11-01',
  devtools: { enabled: true },
  modules: ['@nuxtjs/tailwindcss', '@vueuse/nuxt', '@pinia/nuxt'],
  runtimeConfig: {
    public: {
      apiBase: process.env.API_BASE_URL || 'http://backend:8080'
    }
  },
  nitro: {
      experimental: {
          websocket: true
      },
  }
})