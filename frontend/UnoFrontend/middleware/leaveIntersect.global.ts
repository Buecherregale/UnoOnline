import type { Room } from "~/util/models";
import {clearGameCookies, getRoomFromCookie} from "~/util/cookieHelpers";
import {getIDFromCookie} from "~/util/getIDFromCookie";
import path from "pathe";

/**
 * Global middleware to warn users before leaving active lobbies
 * Prevents accidental navigation away from game rooms
 */
export default defineNuxtRouteMiddleware(async (to, from) => {
    // Skip on server-side rendering
    if (import.meta.server) return;

    // Check if leaving a lobby and navigating to different page
    if (
        from.fullPath.indexOf("/lobby") > -1 &&
        to.fullPath !== from.fullPath &&
        !to.fullPath.includes("/game")
    ) {
        // Show confirmation dialog before leaving lobby
        if (!window.confirm("You are in a Lobby are you Sure you want to leave?")) {
            return abortNavigation();
        }

        const id = getIDFromCookie();
        const currentRoom = getRoomFromCookie();
        const roomID = currentRoom?.id || from.fullPath.split("lobby-")[1];

        try {
            await $fetch(`/api/rooms/${roomID}/players`, {
                method: "DELETE",
                body: {
                    id: id,
                },
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
