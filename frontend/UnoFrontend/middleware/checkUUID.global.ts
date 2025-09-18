import type { Player } from "~/util/models";
import {loadPlayerFromCookie} from "~/util/playerCookie";

/**
 * Global authentication middleware
 * Redirects unauthenticated users to login page
 */
export default defineNuxtRouteMiddleware((to, from) => {
  // Skip authentication check for login page
  if (to.fullPath !== "/login") {
    // Get player data from cookie
    const player: Player | null = loadPlayerFromCookie();

    // Redirect to login if no valid player ID
    if (!player) {
      return navigateTo("/login");
    }
  }
});
