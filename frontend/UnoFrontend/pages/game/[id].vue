<script setup lang="ts">
import type { Card, Player, Room, uuid } from "~/util/models";
import { useGameLogic } from "~/composables/useGameLogic";
import { useWebSocket } from "~/composables/useWebSocket";
import { useGameStore } from "~/stores/game";
import { useGameState } from "~/composables/useGameState";
import { useGameApi } from "~/composables/useGameApi";
import { useGameMessages } from "~/composables/useGameMessages";
import { useGameEventHandlers } from "~/composables/useGameEventHandlers";

const route = useRoute();
const gameId: string = route.params.id as string;
const shouldStartGame = ref<boolean>(route.query.start === "true");
const lastMessageID = ref<uuid | null>(null);

const roomStore = useRoomStore();
const gameStore = useGameStore();

// Initialize state using the new composable
const { initializePlayer, initializeRoom } = useGameState();
const room = ref<Room | null>(null);
const currentPlayer = ref<Player | null>(null);

// Initialize player and room
try {
  currentPlayer.value = initializePlayer();
  room.value = initializeRoom();
} catch (error) {
  console.error("Failed to initialize game state:", error);
}

const isLoading = ref<boolean>(false);
const errorMessage = ref<string>("");

const { addEventHandlers, isConnected, removeEventHandlers, sendMessage } =
  useWebSocket();
const { startRoom, startGame, loadRoomData: loadRoomDataApi } = useGameApi();
const { createDrawCardMessage, createPlayCardMessage } = useGameMessages();
const { createGameEventHandlers } = useGameEventHandlers(
  currentPlayer,
  lastMessageID
);

/**
 * Fetches room data and initializes game state
 */
onMounted(async (): Promise<void> => {
  if (isConnected()) {
    console.log("adding event handlers");
    setGameHandlers();

    // If this page was loaded via start button, trigger game start API call
    if (shouldStartGame.value) {
      await startRoomViaAPI();
    }
  } else {
    console.error("Unable to connect to the game");
  }

  await loadRoomData();

  await new Promise((r) => setTimeout(r, 1000));

  if (shouldStartGame.value) {
    await startGameViaAPI();
  }
});

onBeforeUnmount((): void => {
  removeEventHandlers(["onGameStarted", "onError", "onAskCard", "onDrawCard"]);
  // Clear any running timer when leaving the page
  gameStore.clearTimer();
});

function setGameHandlers() {
  const gameEventHandlers = createGameEventHandlers();
  addEventHandlers(gameEventHandlers);
}

/**
 * Starts the room via API call (only called by host)
 */
async function startRoomViaAPI(): Promise<void> {
  if (!currentPlayer.value) {
    console.error("No current player available");
    return;
  }

  await startRoom(gameId, currentPlayer.value.id);
}

/**
 * Starts the game via API call (only called by host)
 */
async function startGameViaAPI(): Promise<void> {
  if (!currentPlayer.value) {
    console.error("No current player available");
    return;
  }

  await startGame(gameId, currentPlayer.value.id);

  // Reset the start flag after successful API call
  shouldStartGame.value = false;
}

async function loadRoomData() {
  isLoading.value = true;
  errorMessage.value = "";

  try {
    const data = await loadRoomDataApi(gameId);
    if (data) {
      room.value = data;
      roomStore.updateRoom(data);
    }
  } catch (error) {
    if (error instanceof Error) {
      errorMessage.value = error.message;
    } else {
      console.error("Unknown error:", error);
      errorMessage.value = "Ein unbekannter Fehler ist aufgetreten";
    }
  } finally {
    isLoading.value = false;
  }
}

/**
 * Computes player positions based on count and current player position
 */
const playerPositions = computed(() => {
  if (!room.value?.players || !currentPlayer.value) {
    return [];
  }

  return useGameLogic().getPlayerPositions(
    room.value.players,
    currentPlayer.value.id
  );
});

/**
 * Handles drawing a card from the deck
 */
function handleDrawCard(): void {
  console.log("Karte vom Stapel gezogen");

  const drawCardMessage = createDrawCardMessage(
    currentPlayer.value!,
    lastMessageID.value!
  );
  sendMessage(drawCardMessage);

  // Clear timer when player takes action
  gameStore.clearTimer();
}

/**
 * Handles playing a card from player's hand
 */
function handlePlayCard(card: Card, index: number) {
  console.log(`${card.color} ${card.value} gespielt`);
  console.log("Karte gespielt");

  const playCardMessage = createPlayCardMessage(
    currentPlayer.value!,
    card,
    lastMessageID.value!
  );

  sendMessage(playCardMessage);
  gameStore.removeCardFromHand(index);

  // Clear timer when player takes action
  gameStore.clearTimer();
}
</script>

<template>
  <div class="game-container">
    <!-- Error Message -->
    <div
      v-if="errorMessage"
      class="fixed top-4 left-1/2 transform -translate-x-1/2 bg-red-100 border border-red-400 text-red-700 px-4 py-3 rounded z-50"
    >
      {{ errorMessage }}
    </div>

    <!-- Loading State -->
    <div v-if="isLoading" class="game-board flex items-center justify-center">
      <div class="bg-white p-8 rounded-lg shadow-lg flex items-center gap-4">
        <div
          class="animate-spin rounded-full h-6 w-6 border-b-2 border-blue-500"
        ></div>
        <span>Spiel wird geladen...</span>
      </div>
    </div>

    <!-- Game Content -->
    <div v-else class="game-board">
      <!-- Player positions -->
      <GamePlayerCard
        v-for="playerPos in playerPositions"
        :key="playerPos.player.id"
        :player="playerPos.player"
        :position="playerPos.position"
        :card-count="7"
      />

      <!-- Game center area -->
      <GameCenter
        :top-card="gameStore.getTopCard"
        :deck-count="gameStore.getDeckCount"
        @draw-card="handleDrawCard"
      />

      <!-- Current player's hand -->
      <GamePlayerHand
        :cards="gameStore.getPlayerHand"
        @play-card="handlePlayCard"
      />
    </div>
  </div>
</template>

<style scoped>
.game-container {
  width: 100vw;
  height: 100vh;
  background: linear-gradient(135deg, #2d5a27 0%, #1a3d1a 100%);
  display: flex;
  justify-content: center;
  align-items: center;
  font-family: "Arial", sans-serif;
}

.game-board {
  position: relative;
  width: 90vw;
  height: 90vh;
  max-width: 1200px;
  max-height: 800px;
  border-radius: 20px;
  background: radial-gradient(ellipse at center, #4a7c59 0%, #2d5a27 70%);
  box-shadow: inset 0 0 50px rgba(0, 0, 0, 0.3);
}

/* Responsive adjustments */
@media (max-width: 768px) {
  .game-board {
    width: 95vw;
    height: 95vh;
  }
}
</style>
