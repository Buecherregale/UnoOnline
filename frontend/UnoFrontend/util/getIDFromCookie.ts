export const getIDFromCookie = () => {
    const playerCookie = useCookie('playerUUID');
    const playerStr: string = playerCookie.value ?? ""
    if (!playerStr || playerStr.length === 0) {
        return undefined;
    }
    const player: Player = JSON.parse(JSON.stringify(playerStr));
    return player.id;
}
