export default defineNuxtRouteMiddleware((from) => {
  if (from.name !== 'login' && !useAuth().value) {
    return navigateTo('login')
  }
})
