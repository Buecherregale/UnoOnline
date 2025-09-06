import type { Room } from "~/util/models";
import { getIDFromCookie } from "~/util/getIDFromCookie";

/**
 * Middleware for lobby routes with automatic room joining
 * Validates room existence and adds player if not already a member
 */
export default defineNuxtRouteMiddleware(async (to, from) => {
  // Skip execution on server-side rendering
  if (import.meta.server) return;

  // Check if navigating to a lobby route
  if (to.fullPath.indexOf("/lobby") > -1) {
    // Extract room ID from URL path
    const roomID = to.fullPath.split("/lobby-")[1];

    // Validate room ID exists
    if (!roomID) {
      window.alert("Invalid room ID.");
      return navigateTo("/hostOrJoin");
    }

    let room: Room | null = null;

    try {
      // Fetch room data to verify it exists
      room = await $fetch<Room>(`/api/room/${roomID}`);
    } catch (error) {
      // Room not found - redirect to join page
      window.alert("no room found with ID: " + roomID);
      return navigateTo("/hostOrJoin");
    }

    // Get current player's ID from cookie
    const playerID = getIDFromCookie();

    // Check if player is not already in the room
    if (room && !room?.players.some((player) => player.id === playerID)) {
      try {
        // Add player to room
        const responseRoom: Room = await $fetch<Room>(
          `api/room/${roomID}/players`,
          {
            method: "POST",
            body: {
              id: playerID,
            },
          }
        );
        console.log(responseRoom);

        // Update global room state
        useState("room", () => responseRoom);

        // Navigate to lobby
        navigateTo(`/lobby-${roomID}`);
      } catch (error: any) {
        // Handle 409 Conflict (player already in room)
        if (error?.response?.status === 409) {
          return; // Ignore and continue
        }

        // Other errors - redirect to main page
        console.error("Error communicating with internal API:", error);
        return navigateTo("/hostOrJoin");
      }
    }
  }
});
