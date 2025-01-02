<script setup lang="ts">
  import type { Room } from "~/util/models";

  const route = useRoute()
  let id = route.params.id
  let room = useState<Room>('room')
  const players = ref(room?.value?.players)

  onMounted(async () => {
    if (!room.value) {
      const data  = await $fetch(`/api/room/${id}`,
      );
      if(data) {
        room.value = data;
      }
    }
    players.value = room.value?.players || [];
  });
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
  </div>
</template>

<style scoped>

</style>