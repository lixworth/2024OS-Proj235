// https://nuxt.com/docs/api/configuration/nuxt-config
export default defineNuxtConfig({
  compatibilityDate: '2024-04-03',
  devtools: { enabled: true },
  modules: ["@nuxt/ui", "@nuxt/eslint"],
  ssr: false, 
  spaLoadingTemplate: 'spa-loading-template.html',
  router: {
    options: {
      hashMode: true 
    }
  }
})
