import { defineStore } from "pinia";
import type { Room } from "~/util/models";

export const useRoomStore = defineStore("room", {
  state: () => ({
    room: null as Room | null,
    isHost: false,
  }),

  getters: {
    getRoom: (state): Room | null => state.room,
    getRoomId: (state): string => state.room?.id || "",
    getRoomOwner: (state) => state.room?.owner || null,
    getRoomPlayers: (state) => state.room?.players || [],
    hasRoom: (state): boolean => state.room !== null,
    getPlayerCount: (state): number => state.room?.players.length || 0,
    getIsHost: (state): boolean => state.isHost,
  },

  actions: {
    setRoom(room: Room, isHost: boolean) {
      this.room = room;
      this.isHost = isHost;
    },

    clearRoom() {
      this.room = null;
      this.isHost = false;
    },

    updateRoom(room: Room) {
      if (this.room) {
        this.room = room;
      }
    },

    setHost(isHost: boolean) {
      this.isHost = isHost;
    },
  },
});
