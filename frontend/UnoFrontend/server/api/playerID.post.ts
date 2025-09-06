import { Player } from "~/util/models";

/**
 * Server API endpoint to create a new player
 * Acts as proxy between frontend and Go backend
 *
 * @route POST /api/playerID
 * @param event - Nuxt event handler context
 * @returns Promise<Player> - Created player with UUID
 */
export default defineEventHandler(async (event): Promise<Player> => {
  // Get backend API URL from runtime config
  const { apiBase } = useRuntimeConfig().public as { apiBase: string };

  // Extract player name from request body
  const body = await readBody(event);
  const { name } = body;

  // Initialize player object
  let player = {} as Player;
  player.name = name;

  try {
    // Forward request to Go backend API
    const externalResponse: string = await $fetch("/players", {
      method: "POST",
      body: {
        name: name,
      },
      baseURL: apiBase,
    });

    // Parse backend response and extract player ID
    player.id = JSON.parse(externalResponse).id;
    return player;
  } catch (error) {
    // Log error and return 500 status
    console.error("Error communicating with external API:", error);
    throw createError({
      statusCode: 500,
      message: "Failed to communicate with external API",
    });
  }
});
