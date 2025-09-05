import type {Player} from "~/util/models";

/**
 * Extracts and validates player ID from browser cookies
 * @returns Player UUID string or undefined if not found
 */
export const getIDFromCookie = (): string | undefined => {
    // Access player cookie using Nuxt's composable
    const playerCookie = useCookie('playerUUID');
    const playerStr: string = playerCookie.value ?? ""

    // Return undefined for empty/missing cookies
    if (!playerStr || playerStr.length === 0) {
        return undefined;
    }

    // Parse stored player data and extract ID
    const player: Player = JSON.parse(JSON.stringify(playerStr));
    return player.id;
}
