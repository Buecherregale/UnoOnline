<script setup lang="ts">
import type { Player, Room } from "~/util/models";
import { loadPlayerFromCookie } from "~/util/playerCookie";
import { useGameLogic } from "~/composables/useGameLogic";
import { handleApiError, validatePlayerSession } from "~/util/errorUtils";

const route = useRoute();
const gameId: string = route.params.id as string;
const room = ref<Room | null>(null);
const players = ref<Player[]>([]);
const currentPlayerId = ref<string>("");
const isLoading = ref<boolean>(false);
const errorMessage = ref<string>("");

// Mock game state - will be replaced with real WebSocket data later
const playerHand = ref([
  { color: "red", value: "5" },
  { color: "blue", value: "7" },
  { color: "green", value: "skip" },
  { color: "yellow", value: "2" },
  { color: "red", value: "draw2" },
  { color: "wild", value: "wild" },
  { color: "blue", value: "9" },
]);

const topCard = ref({ color: "red", value: "8" });
const deckCount = ref(76);

const { getPlayerPositions } = useGameLogic();

definePageMeta({
  middleware: ["check-join"],
});

/**
 * Fetches room data and initializes game state
 */
onMounted(async (): Promise<void> => {
  isLoading.value = true;
  errorMessage.value = "";

  try {
    const data: Room = await $fetch<Room>(`/api/rooms/${gameId}`);
    if (data) {
      room.value = data;
      players.value = data.players || [];

      const currentPlayer: Player | null = loadPlayerFromCookie();
      validatePlayerSession(currentPlayer);
      currentPlayerId.value = currentPlayer.id;
    }
  } catch (error) {
    if (error instanceof Error) {
      errorMessage.value = error.message;
    } else {
      handleApiError(error);
    }
  } finally {
    isLoading.value = false;
  }
});

/**
 * Computes player positions based on count and current player position
 */
const playerPositions = computed(() => {
  return getPlayerPositions(players.value, currentPlayerId.value);
});

/**
 * Handles drawing a card from the deck
 */
function handleDrawCard(): void {
  console.log("Karte vom Stapel gezogen");
  // TODO: Implement WebSocket communication
}

/**
 * Handles playing a card from player's hand
 */
function handlePlayCard(card: any, index: number): void {
  console.log(`${card.color} ${card.value} gespielt`);
  // TODO: Validate card play and send to server
  playerHand.value.splice(index, 1);
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
        :top-card="topCard"
        :deck-count="deckCount"
        @draw-card="handleDrawCard"
      />

      <!-- Current player's hand -->
      <GamePlayerHand :cards="playerHand" @play-card="handlePlayCard" />
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
