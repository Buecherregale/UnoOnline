import {Player} from "~/util/models";
import {pl} from "cronstrue/dist/i18n/locales/pl";

export default defineEventHandler(async (event): Promise<Player> => {
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
            baseURL: 'http://localhost:8080'
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