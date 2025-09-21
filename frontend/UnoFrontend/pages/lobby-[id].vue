<script setup lang="ts">
import type { Room, Player, StartGameRequest } from "~/util/models";
import {
  loadRoomFromCookie,
  getHostStatusFromCookie,
  saveRoomToCookie,
  saveHostStatusToCookie,
} from "~/util/roomCookie";
import type { WebSocketEventHandlers } from "~/util/webSocketHelper";
import { loadPlayerFromCookie } from "~/util/playerCookie";
import { handleApiError, validatePlayerSession } from "~/util/errorUtils";
import { useClipboard } from "@vueuse/core";
import { useWebSocket } from "~/composables/useWebSocket";

definePageMeta({
  middleware: ["check-join"],
});

// Route and player data with explicit types
const route = useRoute();
const id: string = route.params.id as string;
const player: Player | null = loadPlayerFromCookie();

const source = ref<string>(route.params.id as string);
const { copy, copied, isSupported } = useClipboard({ source });

// State management with explicit types
const room = useState<Room | null>("room", (): Room | null => {
  return loadRoomFromCookie();
});

const isHost = useState<boolean>("isHost", (): boolean => {
  return getHostStatusFromCookie();
});

const players = ref<Player[]>(room?.value?.players || []);

// Use the global WebSocket service
const { connect, addEventHandlers, removeEventHandlers, forceDisconnect } =
  useWebSocket();

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

  // Connect to WebSocket using global service
  if (player) {
    connect(player.id, id);

    // Set up event handlers specific to lobby
    const lobbyEventHandlers: WebSocketEventHandlers = {
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
        if (newOwner.id === player.id) {
          isHost.value = true;
          saveHostStatusToCookie(true);
        }
      },
      onRoomStarted: (): void => {
        console.log("Game started");
        // Navigate to game page
        navigateTo(`/game/${id}`);
      },
      onError: (error: Error): void => {
        console.error("WebSocket error:", error);
        // Handle WebSocket errors appropriately
      },
    };

    addEventHandlers(lobbyEventHandlers);
  }
});

/**
 * Clean up lobby-specific event handlers when component is unmounted
 * Note: We don't disconnect the WebSocket to maintain connection across page transitions
 */
onBeforeUnmount((): void => {
  // Remove only lobby-specific event handlers
  removeEventHandlers([
    "onPlayerJoined",
    "onPlayerLeft",
    "onRoomStarted",
    "onError",
  ]);
});

/**
 * Handles leaving the room
 */
async function leaveRoom(): Promise<void> {
  // Force disconnect WebSocket when actually leaving the room
  forceDisconnect();
  await navigateTo(`/hostOrJoin`);
}

/**
 * Handles starting the game (host only)
 */
async function startRoom(): Promise<void> {
  const player = loadPlayerFromCookie();
  validatePlayerSession(player);

  try {
    const requestBody: StartGameRequest = {
      id: player.id,
    };

    //start the game via api call
    await $fetch<Room>(`/api/rooms/${id}`, {
      method: "POST",
      body: requestBody,
    });
  } catch (error) {
    handleApiError(error);
  }
}
</script>

<template>
  <div class="flex flex-col items-center">
    <div class="flex items-center justify-between w-full max-w-md mb-4">
      <p class="text-lg font-bold">Room ID: {{ id }}</p>
      <div v-if="isSupported">
        <button
          @click="copy(source)"
          class="ml-4 px-3 py-1 bg-blue-500 text-white rounded hover:bg-blue-600"
        >
          <span v-if="!copied">Copy</span>
          <span v-else>Copied!</span>
        </button>
      </div>
      <p v-else-if="!isSupported" class="text-sm text-gray-500 ml-4">
        Browser unterstützt Clipboard API nicht
      </p>
    </div>
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
