import type { Player, PlayerWithPosition } from "~/util/models";

export const useGameLogic = () => {
  /**
   * Arranges players in positions around the game board
   * Current player is always positioned at the bottom
   */
  const getPlayerPositions = (
    players: Player[],
    currentPlayerId: string
  ): PlayerWithPosition[] => {
    if (!players.length) return [];

    const playerCount = players.length;
    const currentPlayerIndex = players.findIndex(
      (p: Player) => p.id === currentPlayerId
    );

    if (currentPlayerIndex === -1) return [];

    // Arrange players with current player always at bottom
    const orderedPlayers: Player[] = [];
    for (let i = 0; i < playerCount; i++) {
      const index = (currentPlayerIndex + i) % playerCount;
      orderedPlayers.push(players[index]!);
    }

    const positions: PlayerWithPosition[] = [];

    switch (playerCount) {
      case 2:
        // 2 players: opponent at top, current player at bottom
        positions.push({ player: orderedPlayers[1]!, position: "top" });
        positions.push({ player: orderedPlayers[0]!, position: "bottom" });
        break;

      case 3:
        // 3 players: Counterclockwise from current at bottom
        positions.push({ player: orderedPlayers[0]!, position: "bottom" });
        positions.push({ player: orderedPlayers[1]!, position: "top" });
        positions.push({ player: orderedPlayers[2]!, position: "right" });
        break;

      case 4:
        // 4 players: Counterclockwise from current at bottom
        positions.push({ player: orderedPlayers[0]!, position: "bottom" });
        positions.push({ player: orderedPlayers[1]!, position: "right" });
        positions.push({ player: orderedPlayers[2]!, position: "top" });
        positions.push({ player: orderedPlayers[3]!, position: "left" });
        break;

      default:
        // For 5+ players, we need more complex positioning
        positions.push({ player: orderedPlayers[0]!, position: "bottom" });
        for (let i = 1; i < playerCount; i++) {
          const angle = (360 / (playerCount - 1)) * (i - 1);
          let position: "top" | "left" | "right";

          if (angle <= 90 || angle >= 270) {
            position = "top";
          } else if (angle > 90 && angle < 180) {
            position = "right";
          } else {
            position = "left";
          }

          positions.push({ player: orderedPlayers[i]!, position });
        }
    }

    return positions;
  };

  /**
   * Validates if a card can be played
   */
  const canPlayCard = (
    card: { color: string; value: string },
    topCard: { color: string; value: string } | null
  ): boolean => {
    if (!topCard) return true;

    return (
      card.color === topCard.color ||
      card.value === topCard.value ||
      card.color === "wild"
    );
  };

  /**
   * Calculates the next player's turn
   */
  const getNextPlayer = (
    players: Player[],
    currentPlayerId: string,
    direction: "clockwise" | "counterclockwise" = "clockwise"
  ): Player | null => {
    const currentIndex = players.findIndex((p) => p.id === currentPlayerId);
    if (currentIndex === -1) return null;

    const increment = direction === "clockwise" ? 1 : -1;
    const nextIndex =
      (currentIndex + increment + players.length) % players.length;

    return players[nextIndex] || null;
  };

  return {
    getPlayerPositions,
    canPlayCard,
    getNextPlayer,
  };
};
