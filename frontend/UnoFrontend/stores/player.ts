import { defineStore } from "pinia";
import type { Player } from "~/util/models";

export const usePlayerStore = defineStore("player", {
  state: () => ({
    player: null as Player | null,
    isLoggedIn: false,
  }),

  getters: {
    getPlayer: (state): Player | null => state.player,
    getPlayerName: (state): string => state.player?.name || "",
    getPlayerId: (state): string => state.player?.id || "",
    hasPlayer: (state): boolean => state.isLoggedIn && state.player !== null,
  },

  actions: {
    setPlayer(player: Player) {
      this.player = player;
      this.isLoggedIn = true;
    },

    clearPlayer() {
      this.player = null;
      this.isLoggedIn = false;
    },

    updatePlayerName(name: string) {
      if (this.player) {
        this.player.name = name;
      }
    },
  },
});
