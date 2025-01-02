import { Room } from "~/util/models";

export default defineEventHandler(async (event): Promise<Room> => {
    const body = await readBody(event);
    const { id } = body;
    let room = {} as Room;

    try {
        const externalResponse: string = await $fetch('/room', {
            method: 'POST',
            body: {
                id: id,
            },
            baseURL: 'http://localhost:8080'
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