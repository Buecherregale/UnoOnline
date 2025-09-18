<template>
  <div class="player-area" :class="[`player-${position}`]">
    <div class="player-info">
      <div class="player-name">{{ player.name }}</div>
      <div class="player-cards">
        <span class="cards-count">{{ cardCount }} Karten</span>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import type { Player } from "~/util/models";

interface Props {
  player: Player;
  position: "top" | "bottom" | "left" | "right";
  cardCount: number;
}

defineProps<Props>();
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

@media (max-width: 768px) {
  .player-info {
    padding: 8px 12px;
    font-size: 14px;
  }
}
</style>
