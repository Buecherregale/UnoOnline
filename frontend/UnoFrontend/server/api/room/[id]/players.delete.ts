export default defineEventHandler(async (event): Promise<void> => {
    const { apiBase } = useRuntimeConfig().public as { apiBase: string };
    const roomID = getRouterParam(event, 'id')
    const body = await readBody(event);
    const { id } = body;

    try {
        await $fetch(`/room/${roomID}/players`, {
            method: 'DELETE',
            baseURL: apiBase,
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