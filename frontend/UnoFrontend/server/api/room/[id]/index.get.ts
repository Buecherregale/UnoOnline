import { Room } from "~/util/models";

/**
 * Server API endpoint to get room information
 * Acts as proxy between frontend and Go backend
 *
 * @route GET /api/room/{id}
 * @param event - Nuxt event handler context
 * @returns Promise<Room> - Room data with current players
 */
export default defineEventHandler(async (event): Promise<Room> => {
    // Get backend API URL from runtime config
    const { apiBase } = useRuntimeConfig().public as { apiBase: string };

    // Extract room ID from URL parameters
    const id = getRouterParam(event, 'id')

    // Initialize room object
    let room = {} as Room;

    try {
        // Fetch room data from Go backend
        const externalResponse: string = await $fetch(`/room/${id}`, {
            method: 'GET',
            baseURL: apiBase,
        });

        // Parse and return room data
        room = JSON.parse(externalResponse);
        return room;
    } catch (error) {
        console.error('Error communicating with external API:', error);
        throw createError({
            statusCode: 500,
            message: 'Failed to communicate with external API',
        });
    }
});