import {getIDFromCookie} from "~/util/getIDFromCookie";

/**
 * Server API endpoint to start a room
 * Acts as proxy between frontend and Go backend
 *
 * @route Post /api/room/{id}
 * @param event - Nuxt event handler context
 * @returns boolean - if starting the room was successful
 */
export default defineEventHandler(async (event): Promise<boolean> => {
    // Get backend API URL from runtime config
    const { apiBase } = useRuntimeConfig().public as { apiBase: string };

    // Extract room ID from URL parameters
    const id = getRouterParam(event, "id");
    // get Player from Cookies
    const player = getIDFromCookie();

    try {
        // Fetch room data from Go backend
        const externalResponse: string = await $fetch(`/room/${id}`, {
            method: "POST",
            baseURL: apiBase,
            body: {
                id: player
            }
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
