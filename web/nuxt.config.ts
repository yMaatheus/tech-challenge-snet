// https://nuxt.com/docs/api/configuration/nuxt-config
export default defineNuxtConfig({
  srcDir: "./src",
  compatibilityDate: "2025-05-15",
  devtools: { enabled: true },
  modules: ["@nuxt/eslint", "@nuxt/icon", "@nuxt/ui", "@nuxtjs/tailwindcss"],
  runtimeConfig: {
    public: {
      apiBase: "http://localhost:8080",
    },
  },
});
