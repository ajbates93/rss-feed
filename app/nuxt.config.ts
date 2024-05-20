// https://nuxt.com/docs/api/configuration/nuxt-config
export default defineNuxtConfig({
  devtools: { enabled: true },
  modules: ["@nuxt/ui", "@pinia/nuxt", "@nuxtjs/google-fonts"],
  ssr: true,

  googleFonts: {
    families: {
      "DM Sans": [400, 500, 700],
    },
  },

  css: ["@/assets/css/main.css"],
});
