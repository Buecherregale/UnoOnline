<script setup lang="ts">
import type { Room } from "~/util/models";
import { getIDFromCookie } from "~/util/getIDFromCookie";

definePageMeta({
  middleware: ["check-join"],
});

const route = useRoute();
let id = route.params.id;
let room = useState<Room>("room");
const players = ref(room?.value?.players);

onMounted(async () => {
  if (!room.value) {
    const data = await $fetch(`/api/room/${id}`);
    if (data) {
      room.value = data;
    }
  }
  players.value = room.value?.players || [];
});

async function leaveRoom() {
  const id = getIDFromCookie();
  const roomID: string = useState<Room>("room").value.id;
  try {
    await $fetch(`/api/room/${roomID}/players`, {
      method: "DELETE",
      body: {
        id: id,
      },
    });
    navigateTo(`/hostOrJoin`);
  } catch (error) {
    console.error("Error communicating with internal API:", error);
    throw createError({
      statusCode: 500,
      message: "Failed to communicate with internal API",
    });
  }
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
  </div>
</template>

<style scoped></style>
