import { Room } from "~/util/models";

export default defineEventHandler(async (event): Promise<Room> => {
    const roomID = getRouterParam(event, 'id')
    const body = await readBody(event);
    const { id } = body;

    let room = {} as Room;

    try {
        const externalResponse: string = await $fetch(`/room/${roomID}/players`, {
            method: 'POST',
            baseURL: 'http://localhost:8080',
            body: {
                id: id,
            },
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