import { Room } from "~/util/models";

export default defineEventHandler(async (event): Promise<Room> => {
    const { apiBase } = useRuntimeConfig().public as { apiBase: string };
    const roomID = getRouterParam(event, 'id')
    const body = await readBody(event);
    const { id } = body;

    let room = {} as Room;

    try {
        const externalResponse: string = await $fetch(`/room/${roomID}/players`, {
            method: 'POST',
            baseURL: apiBase,
            body: {
                id: id,
            },
        });

        room = JSON.parse(externalResponse);
        return room;
    } catch (error: any) {
        if (error?.response?.status === 409) {
            throw error;
        }
        console.error('Error communicating with external API:', error);
        throw createError({
            statusCode: 500,
            message: 'Failed to communicate with external API',
        });
    }
});