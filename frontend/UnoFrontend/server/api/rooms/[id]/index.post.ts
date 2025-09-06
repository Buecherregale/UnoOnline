import { getIDFromCookie } from "~/util/getIDFromCookie";

/**
 * Server API endpoint to start a rooms
 * Acts as proxy between frontend and Go backend
 *
 * @route Post /api/rooms/{id}
 * @param event - Nuxt event handler context
 * @returns boolean - if starting the rooms was successful
 */
export default defineEventHandler(async (event): Promise<boolean> => {
  // Get backend API URL from runtime config
  const { apiBase } = useRuntimeConfig().public as { apiBase: string };

  // Extract rooms ID from URL parameters
  const id = getRouterParam(event, "id");
  // get Player from Cookies
  const player = getIDFromCookie();

  try {
    // Fetch rooms data from Go backend
    const externalResponse: string = await $fetch(`/rooms/${id}`, {
      method: "POST",
      baseURL: apiBase,
      body: {
        id: player,
      },
    });

    // check Response Code
    return externalResponse === "200";
  } catch (error) {
    console.error("Error communicating with external API:", error);
    throw createError({
      statusCode: 500,
      message: "Failed to communicate with external API",
    });
  }
});
