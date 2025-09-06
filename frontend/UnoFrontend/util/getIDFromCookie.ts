import type { Player } from "~/util/models";

/**
 * Extracts and validates player ID from browser cookies
 * @returns Player UUID string or undefined if not found
 */
export const getIDFromCookie = (): string => {
  // Access player cookie using Nuxt's composable
  const playerCookie = useCookie("playerUUID");
  const playerStr: string = playerCookie.value ?? "";

  // Return undefined for empty/missing cookies
  if (!playerStr || playerStr.length === 0) {
    //instead of return redirect player to home page
    navigateTo("/login");
    return "";
  }

  // Parse stored player data and extract ID
  const player: Player = JSON.parse(JSON.stringify(playerStr));
  return player.id;
};
