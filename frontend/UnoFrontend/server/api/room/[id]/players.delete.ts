export default defineEventHandler(async (event): Promise<void> => {
    const roomID = getRouterParam(event, 'id')
    const body = await readBody(event);
    const { id } = body;

    try {
        await $fetch(`/room/${roomID}/players`, {
            method: 'DELETE',
            baseURL: 'http://localhost:8080',
            body: {
                id: id,
            },
        });
    } catch (error) {
        console.error('Error communicating with external API:', error);
        throw createError({
            statusCode: 500,
            message: 'Failed to communicate with external API',
        });
    }
});