<template>
  <div class="game-center">
    <div class="deck-area">
      <div class="draw-pile" @click="$emit('drawCard')">
        <div class="card card-back">
          <span>Deck</span>
          <div class="card-count">{{ deckCount }}</div>
        </div>
      </div>
      <div class="discard-pile">
        <div
          class="card card-front"
          :style="{ backgroundColor: mapColor(topCard?.color!) }"
        >
          <span v-if="topCard">{{ mapValue(topCard.value) }}</span>
          <span v-else>Aktuelle Karte</span>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import type { Card } from "~/util/models";
import { mapColor, mapValue } from "~/util/cardParser";

interface Props {
  topCard?: Card;
  deckCount: number;
}

interface Emits {
  drawCard: [];
}

defineProps<Props>();
defineEmits<Emits>();
</script>

<style scoped>
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
  box-shadow: 0 4px 8px rgba(0, 0, 0, 0.2);
  position: relative;
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

.card-count {
  position: absolute;
  bottom: 4px;
  right: 4px;
  font-size: 10px;
  background: rgba(0, 0, 0, 0.7);
  color: white;
  padding: 2px 4px;
  border-radius: 4px;
}

@media (max-width: 768px) {
  .card {
    width: 60px;
    height: 90px;
    font-size: 10px;
  }
}
</style>
