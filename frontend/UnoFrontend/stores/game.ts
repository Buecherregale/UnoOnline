import { defineStore } from "pinia";
import type { Card } from "~/util/models";

interface GameState {
  playerHand: Card[];
  topCard: Card;
  deckCount: number;
  currentPlayerTimer: number;
  currentPlayerWithTimer: string | null;
  timerInterval: number | null;
}

export const useGameStore = defineStore("game", {
  state: (): GameState => ({
    playerHand: [],
    topCard: {
      color: "green",
      value: 3,
      chosen: null,
    },
    deckCount: 0,
    currentPlayerTimer: 0,
    currentPlayerWithTimer: null,
    timerInterval: null,
  }),
  getters: {
    getPlayerHand: (state) => state.playerHand,
    getTopCard: (state): Card => state.topCard,
    getDeckCount: (state) => state.deckCount,
    getCurrentPlayerTimer: (state) => state.currentPlayerTimer,
    getCurrentPlayerWithTimer: (state) => state.currentPlayerWithTimer,
  },
  actions: {
    setPlayerHand(cards: Card[]) {
      this.playerHand = cards;
    },
    addToPlayerHand(cards: Card[]) {
      this.playerHand.push(...cards);
    },
    removeCardFromHand(index: number) {
      this.playerHand.splice(index, 1);
    },
    setTopCard(card: Card) {
      this.topCard = card;
    },
    setDeckCount(count: number) {
      this.deckCount = count;
    },
    startPlayerTimer(playerId: string) {
      // Clear any existing timer
      this.clearTimer();

      this.currentPlayerWithTimer = playerId;
      this.currentPlayerTimer = 25;

      this.timerInterval = setInterval(() => {
        this.currentPlayerTimer--;
        if (this.currentPlayerTimer <= 0) {
          this.clearTimer();
        }
      }, 1000) as any;
    },
    clearTimer() {
      if (this.timerInterval) {
        clearInterval(this.timerInterval);
        this.timerInterval = null;
      }
      this.currentPlayerTimer = 0;
      this.currentPlayerWithTimer = null;
    },
    resetGameState() {
      this.playerHand = [];
      this.topCard = {
        color: "green",
        value: 3,
        chosen: null,
      };
      this.deckCount = 0;
      this.clearTimer();
    },
  },
});
