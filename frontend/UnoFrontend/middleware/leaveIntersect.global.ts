import type { LeaveRoomRequest, Room } from "~/util/models";
import { clearGameCookies, loadRoomFromCookie } from "~/util/roomCookie";
import { loadPlayerFromCookie } from "~/util/playerCookie";

/**
 * Global middleware to warn users before leaving active lobbies
 * Prevents accidental navigation away from game rooms
 */
export default defineNuxtRouteMiddleware(async (to, from) => {
  // Skip on server-side rendering
  if (import.meta.server) return;

  // Check if leaving a lobby and navigating to different page except game
  let toGame =
    from.fullPath.includes("/lobby") &&
    to.fullPath !== from.fullPath &&
    !to.fullPath.includes("/game");
  if (toGame) {
    // Show confirmation dialog before leaving lobby
    if (!window.confirm("You are in a Lobby are you Sure you want to leave?")) {
      return abortNavigation();
    }

    const player = loadPlayerFromCookie();
    const id = player?.id;
    const currentRoom = loadRoomFromCookie();
    const roomID = currentRoom?.id || from.fullPath.split("lobby-")[1];

    const requestBody: LeaveRoomRequest = {
      id: id!,
    };

    try {
      await $fetch(`/api/rooms/${roomID}/players`, {
        method: "DELETE",
        body: requestBody,
      });

      useState<Room | null>("rooms").value = null;

      clearGameCookies();
    } catch (error) {
      console.error("Error communicating with internal API:", error);
      throw createError({
        statusCode: 500,
        message: "Failed to communicate with internal API",
      });
    }
  }
});
