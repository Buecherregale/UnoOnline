<script setup lang="ts">
const showPopup = ref(false);
const selectedPlayers = ref(2);

function confirmedHost() {
  alert(`You selected ${selectedPlayers.value} players.`);
  showPopup.value = false;
}
</script>

<template>
  <div
      class="flex items-center justify-center min-h-screen bg-zinc-100"
      role="region"
      aria-label="Game Selection"
      @click="showPopup = true"
  >
    <div class="flex gap-5">
      <button
          class="flex flex-col items-center justify-center p-10 bg-white border border-gray-300 rounded shadow-md hover:bg-gray-100"
          aria-label="Host a Game"
      >
        <img
            src="https://cdn.builder.io/api/v1/image/assets%2F05a77bbafa28470b8ee45012047001f8%2Fbabeb6cc0365430b80c152aebd5d4865"
            class="w-16 h-16 mb-4"
            alt="Host A Game Icon"
        />
        <span class="text-lg font-semibold">Host Game</span>
      </button>

      <button
          class="flex flex-col items-center justify-center p-10 bg-white border border-gray-300 rounded shadow-md hover:bg-gray-100"
          aria-label="Join a Game"
      >
        <img
            src="https://cdn.builder.io/api/v1/image/assets%2F05a77bbafa28470b8ee45012047001f8%2F4babab5f4b0f48049c82c51268547b6a"
            class="w-16 h-16 mb-4"
            alt="Join A Game Icon"
        />
        <span class="text-lg font-semibold">Join Game</span>
      </button>
    </div>

    <div
        v-if="showPopup"
        class="fixed inset-0 flex items-center justify-center bg-black bg-opacity-50"
    >
      <div class="bg-white p-5 rounded shadow-lg w-80">
        <h2 class="text-lg font-semibold mb-4">Select Number of Players</h2>
        <select
            v-model="selectedPlayers"
            class="w-full p-2 border border-gray-300 rounded mb-4"
        >
          <option v-for="n in 7" :key="n" :value="n + 1">{{ n + 1 }}</option>
        </select>
        <div class="flex justify-end gap-3">
          <button
              class="px-4 py-2 bg-gray-300 rounded hover:bg-gray-400"
              @click.stop="showPopup = false"
          >
            Cancel
          </button>
          <button
              class="px-4 py-2 bg-blue-500 text-white rounded hover:bg-blue-600"
              @click="confirmedHost"
          >
            OK
          </button>
        </div>
      </div>
    </div>

  </div>
</template>

<style scoped>

</style>