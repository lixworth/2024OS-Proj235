export default defineNuxtRouteMiddleware((from,to) => {
  if(from.name != 'login' && !useAuth().value){
    return navigateTo('login');
  }
})
