import { Room } from "~/util/models";

/**
 * Server API endpoint to create a new game rooms
 * Acts as proxy between frontend and Go backend
 *
 * @route POST /api/rooms
 * @param event - Nuxt event handler context
 * @returns Promise<Room> - Created rooms with player as owner
 */
export default defineEventHandler(async (event): Promise<Room> => {
  // Get backend API URL from runtime config
  const { apiBase } = useRuntimeConfig().public as { apiBase: string };

  // Extract player ID from request body
  const body = await readBody(event);
  const { id } = body;
  let room = {} as Room;

  try {
    // Forward rooms creation request to Go backend
    const externalResponse: string = await $fetch("/rooms", {
      method: "POST",
      body: {
        id: id,
      },
      baseURL: apiBase,
    });

    // Parse and return rooms data from backend
    room = JSON.parse(externalResponse);
    return room;
  } catch (error) {
    console.error("Error communicating with external API:", error);
    throw createError({
      statusCode: 500,
      message: "Failed to communicate with external API",
    });
  }
});
