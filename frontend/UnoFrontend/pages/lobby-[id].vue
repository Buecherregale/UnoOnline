<script setup lang="ts">
import type { Room, Player } from "~/util/models";
import {
  loadRoomFromCookie,
  getHostStatusFromCookie,
  saveRoomToCookie, saveHostStatusToCookie,
} from "~/util/roomCookie";
import WebSocketHelper, {
  type WebSocketEventHandlers,
} from "~/util/webSocketHelper";
import { loadPlayerFromCookie } from "~/util/playerCookie";

definePageMeta({
  middleware: ["check-join"],
});

// Route and player data with explicit types
const route = useRoute();
const id: string = route.params.id as string;
const player: Player | null = loadPlayerFromCookie();

// State management with explicit types
const room = useState<Room | null>("room", (): Room | null => {
  return loadRoomFromCookie();
});

const isHost = useState<boolean>("isHost", (): boolean => {
  return getHostStatusFromCookie();
});

const players = ref<Player[]>(room?.value?.players || []);

// WebSocket instance - use ref to maintain single instance
const wsHelper = ref<WebSocketHelper | null>(null);

// Save rooms to cookie whenever it changes
watch(
  room,
  (newRoom: Room | null): void => {
    if (newRoom) {
      saveRoomToCookie(newRoom);
      players.value = newRoom.players || [];
    }
  },
  { deep: true }
);

/**
 * Initializes component and establishes WebSocket connection
 */
onMounted(async (): Promise<void> => {
  if (!room.value) {
    const data: Room = await $fetch<Room>(`/api/rooms/${id}`);
    if (data) {
      room.value = data;
    }
  }
  players.value = room.value?.players || [];

  // Only create WebSocket connection if it doesn't exist and player is available
  if (!wsHelper.value && player) {
    wsHelper.value = new WebSocketHelper(player.id, id);

    // Set up event handlers
    const eventHandlers: WebSocketEventHandlers = {
      onPlayerJoined: (newPlayer: Player): void => {
        console.log("new Player:", newPlayer);
        players.value.push(newPlayer);
      },
      onPlayerLeft: (oldPlayer: Player, newOwner: Player): void => {
        console.log("Player left:", oldPlayer);
        // Remove player from list
        let index = players.value.findIndex((p) => p.id === oldPlayer.id);
        console.log("index:", index);
        if (index !== -1) {
          players.value.splice(index, 1);
        }
        // Update room owner
        room.value!.owner = newOwner;
        // Update status if current player is the new owner
        if(newOwner.id === player.id) {
          isHost.value = true;
          saveHostStatusToCookie(true)
        }
      },
      onError: (error: Error): void => {
        console.error("WebSocket error:", error);
        // Handle WebSocket errors appropriately
      },
    };

    wsHelper.value.setEventHandlers(eventHandlers);
  }
});

/**
 * Cleans up WebSocket connection when component is unmounted
 */
onBeforeUnmount((): void => {
  if (wsHelper.value) {
    wsHelper.value.disconnect();
    wsHelper.value = null;
  }
});

/**
 * Handles leaving the room
 */
async function leaveRoom(): Promise<void> {
  // Disconnect WebSocket before leaving
  if (wsHelper.value) {
    wsHelper.value.disconnect();
    wsHelper.value = null;
  }
  await navigateTo(`/hostOrJoin`);
}

/**
 * Handles starting the game (host only)
 */
async function startRoom(): Promise<void> {
  await navigateTo(`/game/${id}`);
}
</script>

<template>
  <div class="flex flex-col items-center">
    <p class="text-lg font-bold mb-4">Room ID: {{ id }}</p>
    <div class="grid grid-cols-2 gap-4">
      <div
        v-for="player in players"
        :key="player.id"
        class="bg-gray-100 text-center p-4 rounded-xl border-2 border-gray-300 shadow-md"
      >
        {{ player.name }}
      </div>
    </div>
    <button
      @click="leaveRoom"
      class="mt-8 px-6 py-3 bg-red-500 text-white font-bold rounded-lg shadow-lg hover:bg-red-600 focus:outline-none"
    >
      Leave
    </button>
    <button
      v-if="isHost"
      @click="startRoom"
      class="mt-4 px-6 py-3 bg-green-500 text-white font-bold rounded-lg shadow-lg hover:bg-green-600 focus:outline-none"
    >
      Start Game
    </button>
  </div>
</template>

<style scoped></style>
