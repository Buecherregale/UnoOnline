import type { Card, uuid, Player } from "~/util/models";
import type { WebSocketEventHandlers } from "~/util/webSocketHelper";

export const useGameEventHandlers = (
  currentPlayer: Ref<Player | null>,
  lastMessageID: Ref<uuid | null>
) => {
  const gameStore = useGameStore();

  const createGameEventHandlers = (): WebSocketEventHandlers => ({
    onGameStarted: (card: Card): void => {
      gameStore.setTopCard(card);
    },
    onDrawCard: (cards: Card[]): void => {
      gameStore.addToPlayerHand(cards);
      console.log("drew Card:", cards[0]);
    },
    onAskCard: (_: Card[], id: uuid): void => {
      if (currentPlayer.value) {
        gameStore.startPlayerTimer(currentPlayer.value.id);
      }
      lastMessageID.value = id;
    },
    onError: (error: Error): void => {
      console.error("WebSocket error:", error);
      // Handle WebSocket errors appropriately
    },
  });

  return {
    createGameEventHandlers,
  };
};
