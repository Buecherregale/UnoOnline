import type {Player} from "~/util/models";

export default defineNuxtRouteMiddleware((to, from) => {
  if (to.fullPath !== "/login") {
      const playerCookie = useCookie('playerUUID');
      const playerStr: string = playerCookie.value ?? ""
      const player: Player = JSON.parse(JSON.stringify(playerStr));
      if (!player.id) {
          return navigateTo("/login");
      }
  }
})