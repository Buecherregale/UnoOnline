import type { Room } from "~/util/models";

/**
 * Global middleware to warn users before leaving active lobbies
 * Prevents accidental navigation away from game rooms
 */
export default defineNuxtRouteMiddleware((to, from) => {
  // Skip on server-side rendering
  if (import.meta.server) return;

  // Check if leaving a lobby and navigating to different page
  if (from.fullPath.indexOf("/lobby") > -1 && to.fullPath !== from.fullPath && !to.fullPath.includes("/game")) {
    // Show confirmation dialog before leaving lobby
    if (!window.confirm("You are in a Lobby are you Sure you want to leave?")) {
      return abortNavigation();
    }
    // Clear room state when leaving confirmed
    useState<Room | null>("room").value = null;
  }
});
