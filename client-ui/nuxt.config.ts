// https://nuxt.com/docs/api/configuration/nuxt-config
export default defineNuxtConfig({
  compatibilityDate: '2024-04-03',
  devtools: { enabled: false },
  modules: ["@nuxt/eslint","@nuxt/ui",'@vueuse/nuxt'],
  css: ["assets/main.scss"],
  ssr: false, 
  spaLoadingTemplate: 'spa-loading-template.html',
  router: {
    options: {
      hashMode: true 
    }
  }
})
