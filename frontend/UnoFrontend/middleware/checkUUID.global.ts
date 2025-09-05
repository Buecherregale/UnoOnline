import type {Player} from "~/util/models";

/**
 * Global authentication middleware
 * Redirects unauthenticated users to login page
 */
export default defineNuxtRouteMiddleware((to, from) => {
  // Skip authentication check for login page
  if (to.fullPath !== "/login") {
      // Get player data from cookie
      const playerCookie = useCookie('playerUUID');
      const playerStr: string = playerCookie.value ?? ""
      const player: Player = JSON.parse(JSON.stringify(playerStr));

      // Redirect to login if no valid player ID
      if (!player.id) {
          return navigateTo("/login");
      }
  }
})