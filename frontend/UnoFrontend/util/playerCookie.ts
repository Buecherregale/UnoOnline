import type {Player, Room} from "~/util/models";

const COOKIE_OPTIONS = {
    maxAge: 60 * 60 * 24, // 24 Stunden
    httpOnly: false,
};

/**
 * saves a Player to a cookie
 */
export function savePlayerToCookie(player: Player): void {
    try {
        const playerCookie = useCookie<Player | null>("player", {
            ...COOKIE_OPTIONS,
            default: () => null,
        });

        playerCookie.value = player;
        console.log("Player saved to cookie:", { name: player.name });
    } catch (error) {
        console.error("Failed to save rooms to cookie:", error);
    }
}

/**
 * loads a Player from a cookie
 */
export function loadPlayerFromCookie(): Player | null {
    try {
        const playerCookie = useCookie<Player | null>("player", {
            ...COOKIE_OPTIONS,
            default: () => null,
        });

        const cookieValue = playerCookie.value;
        console.log("Player loaded from cookie:", cookieValue);

        return cookieValue;
    } catch (error) {
        console.error("Failed to load rooms from cookie:", error);
        return null;
    }
}
