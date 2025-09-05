/**
 * Server API endpoint to remove a player from a room
 * Acts as proxy between frontend and Go backend
 *
 * @route DELETE /api/room/{id}/players
 * @param event - Nuxt event handler context
 * @returns Promise<void> - No content returned
 */
export default defineEventHandler(async (event): Promise<void> => {
    // Get backend API URL from runtime config
    const { apiBase } = useRuntimeConfig().public as { apiBase: string };

    // Extract room ID from URL parameters
    const roomID = getRouterParam(event, 'id')

    // Extract player ID from request body
    const body = await readBody(event);
    const { id } = body;

    try {
        // Send leave request to Go backend
        await $fetch(`/room/${roomID}/players`, {
            method: 'DELETE',
            baseURL: apiBase,
            body: {
                id: id,
            },
        });
    } catch (error) {
        // Log error and return 500 status
        console.error('Error communicating with external API:', error);
        throw createError({
            statusCode: 500,
            message: 'Failed to communicate with external API',
        });
    }
});