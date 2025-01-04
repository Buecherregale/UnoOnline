import { Room } from "~/util/models";

export default defineEventHandler(async (event): Promise<Room> => {
    const { apiBase } = useRuntimeConfig().public as { apiBase: string };
    const id = getRouterParam(event, 'id')

    let room = {} as Room;

    try {
        const externalResponse: string = await $fetch(`/room/${id}`, {
            method: 'GET',
            baseURL: apiBase,
        });

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