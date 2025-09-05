import { Room } from "~/util/models";

/**
 * Server API endpoint to add a player to an existing room
 * Acts as proxy between frontend and Go backend
 *
 * @route POST /api/room/{id}/players
 * @param event - Nuxt event handler context
 * @returns Promise<Room> - Updated room data with new player
 */
export default defineEventHandler(async (event): Promise<Room> => {
  // Get backend API URL from runtime config
  const { apiBase } = useRuntimeConfig().public as { apiBase: string };

  // Extract room ID from URL parameters
  const roomID = getRouterParam(event, "id");

  // Extract player ID from request body
  const body = await readBody(event);
  const { id } = body;

  // Initialize room object
  let room = {} as Room;

  try {
    // Send join request to Go backend
    const externalResponse: string = await $fetch(`/room/${roomID}/players`, {
      method: "POST",
      baseURL: apiBase,
      body: {
        id: id,
      },
    });

    // Parse and return updated room data
    room = JSON.parse(externalResponse);
    return room;
  } catch (error: any) {
    // Pass through 409 Conflict (player already in room)
    if (error?.response?.status === 409) {
      throw error;
    }

    // Log other errors and return 500 status
    console.error("Error communicating with external API:", error);
    throw createError({
      statusCode: 500,
      message: "Failed to communicate with external API",
    });
  }
});
