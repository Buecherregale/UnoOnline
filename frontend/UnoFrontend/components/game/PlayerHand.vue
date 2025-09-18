<template>
  <div class="player-hand">
    <div class="hand-cards">
      <div
        v-for="(card, index) in cards"
        :key="`${card.color}-${card.value}-${index}`"
        class="card card-front"
        :style="{ backgroundColor: card.color }"
        @click="$emit('playCard', card, index)"
      >
        {{ card.value }}
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
interface Card {
  color: string;
  value: string;
}

interface Props {
  cards: Card[];
}

interface Emits {
  playCard: [card: Card, index: number];
}

defineProps<Props>();
defineEmits<Emits>();
</script>

<style scoped>
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
  margin: 0 -10px;
  z-index: 1;
}

.card:hover {
  z-index: 10;
  transform: translateY(-20px);
}

.card-front {
  background: white;
  color: #333;
  border: 2px solid #ddd;
  font-size: 12px;
}

@media (max-width: 768px) {
  .card {
    width: 60px;
    height: 90px;
    font-size: 10px;
    margin: 0 -8px;
  }

  .hand-cards {
    gap: -15px;
  }
}
</style>
