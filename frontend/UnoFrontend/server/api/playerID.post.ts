import {Player} from "~/util/models";

export default defineEventHandler(async (event): Promise<Player> => {
    const { apiBase } = useRuntimeConfig().public as { apiBase: string };
    const body = await readBody(event);
    const { name } = body;
    let player = {} as Player;
    player.name = name;

    try {
        const externalResponse: string = await $fetch('/player', {
            method: 'POST',
            body: {
                name: name,
            },
            baseURL: apiBase,
        });

        player.id = JSON.parse(externalResponse).id;
        return player;
    } catch (error) {
        console.error('Error communicating with external API:', error);
        throw createError({
            statusCode: 500,
            message: 'Failed to communicate with external API',
        });
    }
});