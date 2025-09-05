import { Room } from "~/util/models";

/**
 * Server API endpoint to create a new game room
 * Acts as proxy between frontend and Go backend
 *
 * @route POST /api/room
 * @param event - Nuxt event handler context
 * @returns Promise<Room> - Created room with player as owner
 */
export default defineEventHandler(async (event): Promise<Room> => {
    // Get backend API URL from runtime config
    const { apiBase } = useRuntimeConfig().public as { apiBase: string };

    // Extract player ID from request body
    const body = await readBody(event);
    const { id } = body;
    let room = {} as Room;

    try {
        // Forward room creation request to Go backend
        const externalResponse: string = await $fetch('/room', {
            method: 'POST',
            body: {
                id: id,
            },
            baseURL: apiBase,
        });

        // Parse and return room data from backend
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