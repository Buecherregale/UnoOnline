<script setup lang="ts">
import type { Room, Player } from "~/util/models";
import { getIDFromCookie } from "~/util/getIDFromCookie";

const route = useRoute();
const gameId = route.params.id;
const room = ref<Room | null>(null);
const players = ref<Player[]>([]);
const currentPlayerId = ref<string>("");

definePageMeta({
  middleware: ["check-join"],
});

// Fetch room data and initialize game
onMounted(async () => {
  try {
    const data = await $fetch(`/api/room/${gameId}`);
    if (data) {
      room.value = data;
      players.value = data.players || [];

      // Handle getIDFromCookie properly
      const playerIdResult = getIDFromCookie();
      if (typeof playerIdResult === 'string') {
        currentPlayerId.value = playerIdResult;
      } else {
        console.error("Could not get player ID from cookie");
        // Handle redirect case - the function might have already navigated
      }
    }
  } catch (error) {
    console.error("Error fetching room data:", error);
  }
});

// Position players based on count and current player position
const getPlayerPositions = computed(() => {
  if (!players.value.length) return [];

  const playerCount = players.value.length;
  const currentPlayerIndex = players.value.findIndex(p => p.id === currentPlayerId.value);

  // Arrange players with current player always at bottom
  const orderedPlayers = [];
  for (let i = 0; i < playerCount; i++) {
    const index = (currentPlayerIndex + i) % playerCount;
    orderedPlayers.push(players.value[index]);
  }

  const positions = [];

  if (playerCount === 2) {
    // 2 players: opponent at top, current player at bottom
    positions.push({ player: orderedPlayers[1], position: 'top' });
    positions.push({ player: orderedPlayers[0], position: 'bottom' });
  } else if (playerCount === 3) {
    // 3 players: clockwise from current at bottom
    positions.push({ player: orderedPlayers[0], position: 'bottom' }); // current player
    positions.push({ player: orderedPlayers[1], position: 'top' });
    positions.push({ player: orderedPlayers[2], position: 'right' });
  } else if (playerCount === 4) {
    // 4 players: clockwise from current at bottom
    positions.push({ player: orderedPlayers[0], position: 'bottom' }); // current player
    positions.push({ player: orderedPlayers[1], position: 'right' });
    positions.push({ player: orderedPlayers[2], position: 'top' });
    positions.push({ player: orderedPlayers[3], position: 'left' });
  }

  return positions;
});
</script>

<template>
  <div class="game-container">
    <div class="game-board">
      <!-- Player positions -->
      <div
        v-for="playerPos in getPlayerPositions"
        :key="playerPos.player.id"
        class="player-area"
        :class="[`player-${playerPos.position}`]"
      >
        <div class="player-info">
          <div class="player-name">{{ playerPos.player.name }}</div>
          <div class="player-cards">
            <!-- Placeholder for player cards count -->
            <span class="cards-count">7 Karten</span>
          </div>
        </div>
      </div>

      <!-- Game center area -->
      <div class="game-center">
        <div class="deck-area">
          <div class="draw-pile">
            <div class="card card-back">Deck</div>
          </div>
          <div class="discard-pile">
            <div class="card card-front">
              <!-- Current top card placeholder -->
              <span>Aktuelle Karte</span>
            </div>
          </div>
        </div>
      </div>

      <!-- Current player's hand (always at bottom) -->
      <div class="player-hand">
        <div class="hand-cards">
          <!-- Placeholder for current player's cards -->
          <div class="card card-front" v-for="n in 7" :key="n">
            Karte {{ n }}
          </div>
        </div>
      </div>
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
  font-family: 'Arial', sans-serif;
}

.game-board {
  position: relative;
  width: 90vw;
  height: 90vh;
  max-width: 1200px;
  max-height: 800px;
  border-radius: 20px;
  background: radial-gradient(ellipse at center, #4a7c59 0%, #2d5a27 70%);
  box-shadow: inset 0 0 50px rgba(0,0,0,0.3);
}

/* Player positioning */
.player-area {
  position: absolute;
  display: flex;
  align-items: center;
  justify-content: center;
}

.player-top {
  top: 20px;
  left: 50%;
  transform: translateX(-50%);
}

.player-bottom {
  bottom: 120px;
  left: 50%;
  transform: translateX(-50%);
}

.player-left {
  left: 20px;
  top: 50%;
  transform: translateY(-50%) rotate(-90deg);
}

.player-right {
  right: 20px;
  top: 50%;
  transform: translateY(-50%) rotate(90deg);
}

.player-info {
  background: rgba(255, 255, 255, 0.9);
  padding: 12px 20px;
  border-radius: 25px;
  box-shadow: 0 4px 12px rgba(0,0,0,0.2);
  text-align: center;
  border: 3px solid #1a3d1a;
}

.player-name {
  font-weight: bold;
  font-size: 16px;
  color: #1a3d1a;
  margin-bottom: 5px;
}

.cards-count {
  font-size: 12px;
  color: #666;
  background: #f0f0f0;
  padding: 4px 8px;
  border-radius: 10px;
}

/* Game center */
.game-center {
  position: absolute;
  top: 50%;
  left: 50%;
  transform: translate(-50%, -50%);
  display: flex;
  align-items: center;
  gap: 30px;
}

.deck-area {
  display: flex;
  gap: 20px;
  align-items: center;
}

.card {
  width: 80px;
  height: 120px;
  border-radius: 12px;
  display: flex;
  align-items: center;
  justify-content: center;
  font-weight: bold;
  cursor: pointer;
  transition: transform 0.2s;
  box-shadow: 0 4px 8px rgba(0,0,0,0.2);
}

.card:hover {
  transform: translateY(-5px);
}

.card-back {
  background: linear-gradient(45deg, #d32f2f 0%, #b71c1c 100%);
  color: white;
  border: 2px solid #b71c1c;
}

.card-front {
  background: white;
  color: #333;
  border: 2px solid #ddd;
  font-size: 12px;
}

/* Player hand */
.player-hand {
  position: absolute;
  bottom: 20px;
  left: 50%;
  transform: translateX(-50%);
  width: 80%;
  max-width: 800px;
}

.hand-cards {
  display: flex;
  justify-content: center;
  gap: -20px;
  flex-wrap: wrap;
}

.hand-cards .card {
  margin: 0 -10px;
  z-index: 1;
}

.hand-cards .card:hover {
  z-index: 10;
  transform: translateY(-20px);
}

/* Responsive adjustments */
@media (max-width: 768px) {
  .game-board {
    width: 95vw;
    height: 95vh;
  }

  .player-info {
    padding: 8px 12px;
    font-size: 14px;
  }

  .card {
    width: 60px;
    height: 90px;
    font-size: 10px;
  }

  .hand-cards {
    gap: -15px;
  }

  .hand-cards .card {
    margin: 0 -8px;
  }
}
</style>