<template>
  <div class="player-area" :class="[`player-${position}`]">
    <div class="player-info">
      <div class="player-name">{{ player.name }}</div>
      <div
        v-if="showTimer"
        class="player-timer"
        :class="{ 'timer-warning': timerValue <= 10 }"
      >
        {{ timerValue }}s
      </div>
      <div class="player-cards">
        <span class="cards-count">{{ cardCount }} Karten</span>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import type { Player } from "~/util/models";
import { useGameStore } from "~/stores/game";

interface Props {
  player: Player;
  position: "top" | "bottom" | "left" | "right";
  cardCount: number;
}

const props = defineProps<Props>();
const gameStore = useGameStore();

const showTimer = computed(() => {
  return gameStore.getCurrentPlayerWithTimer === props.player.id;
});

const timerValue = computed(() => {
  return gameStore.getCurrentPlayerTimer;
});
</script>

<style scoped>
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
  box-shadow: 0 4px 12px rgba(0, 0, 0, 0.2);
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

.player-timer {
  font-weight: bold;
  font-size: 18px;
  color: #1a3d1a;
  background: #e8f5e8;
  padding: 6px 12px;
  border-radius: 15px;
  margin: 5px 0;
  border: 2px solid #4a7c59;
  transition: all 0.3s ease;
}

.timer-warning {
  background: #ffebee;
  color: #d32f2f;
  border-color: #d32f2f;
  animation: pulse 1s infinite;
}

@keyframes pulse {
  0% {
    transform: scale(1);
  }
  50% {
    transform: scale(1.05);
  }
  100% {
    transform: scale(1);
  }
}

@media (max-width: 768px) {
  .player-info {
    padding: 8px 12px;
    font-size: 14px;
  }
}
</style>
