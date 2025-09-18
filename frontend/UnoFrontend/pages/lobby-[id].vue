<script setup lang="ts">
import type { Room } from "~/util/models";
import {getIDFromCookie, getPlayerFromCookie} from "~/util/getIDFromCookie";
import {
  getRoomFromCookie,
  getHostStatusFromCookie,
  saveRoomToCookie,
} from "~/util/cookieHelpers";
import WebSocketHelper from "~/util/webSocketHelper";

definePageMeta({
  middleware: ["check-join"],
});

const route = useRoute();
let id = route.params.id;
const player = getPlayerFromCookie();

const room = useState<Room | null>("room", () => {
  return getRoomFromCookie();
});

// Load isHost from cookie
const isHost = useState<boolean>("isHost", () => {
  return getHostStatusFromCookie();
});

const players = ref(room?.value?.players);

// WebSocket instance - use ref to maintain single instance
const wsHelper = ref<WebSocketHelper | null>(null);

// Save rooms to cookie whenever it changes
watch(
  room,
  (newRoom) => {
    if (newRoom) {
      saveRoomToCookie(newRoom);
      players.value = newRoom.players || [];
    }
  },
  { deep: true }
);

onMounted(async () => {
  if (!room.value) {
    const data = await $fetch<Room>(`/api/rooms/${id}`);
    if (data) {
      room.value = data;
    }
  }
  players.value = room.value?.players || [];

  // Only create WebSocket connection if it doesn't exist
  if (!wsHelper.value) {
    wsHelper.value = new WebSocketHelper(player!.id, id as string);

    // Set callback to update room when someone joins
    wsHelper.value.setRoomUpdateCallback((updatedRoom: Room) => {
      console.log('Room updated:', updatedRoom);
      room.value = updatedRoom;
      players.value = updatedRoom.players;
    });


    wsHelper.value.playerJoined(player!);

  }
});

// Clean up WebSocket connection when component is unmounted
onBeforeUnmount(() => {
  if (wsHelper.value) {
    wsHelper.value.disconnect();
    wsHelper.value = null;
  }
});

async function leaveRoom() {
  // Disconnect WebSocket before leaving
  if (wsHelper.value) {
    wsHelper.value.disconnect();
    wsHelper.value = null;
  }
  navigateTo(`/hostOrJoin`);
}

async function startRoom() {
  navigateTo(`/game/${id}`);
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
