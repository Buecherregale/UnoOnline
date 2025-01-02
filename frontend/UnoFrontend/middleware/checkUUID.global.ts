export default defineNuxtRouteMiddleware((to, from) => {
  if (to.fullPath !== "/login") {
      const playerUUIDCookie = useCookie('playerUUID');
      const playerUUID = playerUUIDCookie.value
      if (!playerUUID) {
          alert("No UUID found.");
          return navigateTo("/login");
      }
  }
})