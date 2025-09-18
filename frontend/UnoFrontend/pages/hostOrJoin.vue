<script setup lang="ts">
import { useRoomActions } from "~/composables/useRoomActions";

// Reactive state with explicit types
const showPopupHost = ref<boolean>(false);
const showPopupJoin = ref<boolean>(false);
const selectedPlayerCount = ref<number>(2);
const enteredLobbyID = ref<string>("");
const isLoading = ref<boolean>(false);
const errorMessage = ref<string>("");

const { createRoom, joinRoom, enterLobby } = useRoomActions();

/**
 * Handles host game confirmation with proper error handling
 */
async function confirmedHost(): Promise<void> {
  isLoading.value = true;
  await new Promise(r => setTimeout(r, 2000));
  errorMessage.value = "";

  try {
    const room = await createRoom(selectedPlayerCount.value);
    await enterLobby(room, true);
    showPopupHost.value = false;
  } catch (error) {
    errorMessage.value =
      error instanceof Error ? error.message : "Ein unerwarteter Fehler ist aufgetreten";
  } finally {
    isLoading.value = false;
  }
}

/**
 * Handles join game confirmation with proper error handling
 */
async function confirmedJoin(): Promise<void> {
  isLoading.value = true;
  errorMessage.value = "";

  try {
    const room = await joinRoom(enteredLobbyID.value);
    await enterLobby(room, false);
    showPopupJoin.value = false;
  } catch (error) {
    errorMessage.value =
      error instanceof Error ? error.message : "Ein unerwarteter Fehler ist aufgetreten";
  } finally {
    isLoading.value = false;
  }
}

function closeModal(): void {
  showPopupHost.value = false;
  showPopupJoin.value = false;
  errorMessage.value = "";
}
</script>

<template>
  <div
    class="flex items-center justify-center min-h-screen bg-zinc-100"
    role="region"
    aria-label="Game Selection"
  >
    <div class="flex gap-5">
      <UiGameButton
        label="Host Game"
        icon-src="https://cdn.builder.io/api/v1/image/assets%2F05a77bbafa28470b8ee45012047001f8%2Fbabeb6cc0365430b80c152aebd5d4865"
        icon-alt="Host A Game Icon"
        ariaLabel="Host a Game"
        @click="showPopupHost = true"
      />

      <UiGameButton
        label="Join Game"
        icon-src="https://cdn.builder.io/api/v1/image/assets%2F05a77bbafa28470b8ee45012047001f8%2F4babab5f4b0f48049c82c51268547b6a"
        icon-alt="Join A Game Icon"
        ariaLabel="Join a Game"
        @click="showPopupJoin = true"
      />
    </div>

    <!-- Host Game Modal -->
    <UiModal
      :is-open="showPopupHost"
      title="Select Number of Players"
      :confirm-text="isLoading ? 'Erstelle...' : 'OK'"
      @close="closeModal"
      @confirm="confirmedHost"
    >
      <div v-if="errorMessage" class="mb-4 p-3 bg-red-100 border border-red-400 text-red-700 rounded">
        {{ errorMessage }}
      </div>

      <select
        v-model="selectedPlayerCount"
        class="w-full p-2 border border-gray-300 rounded mb-4"
        :disabled="isLoading"
      >
        <option v-for="n in 7" :key="n" :value="n + 1">{{ n + 1 }}</option>
      </select>
    </UiModal>

    <!-- Join Game Modal -->
    <UiModal
      :is-open="showPopupJoin"
      title="Enter Lobby ID"
      :confirm-text="isLoading ? 'Beitrete...' : 'OK'"
      @close="closeModal"
      @confirm="confirmedJoin"
    >
      <div v-if="errorMessage" class="mb-4 p-3 bg-red-100 border border-red-400 text-red-700 rounded">
        {{ errorMessage }}
      </div>

      <input
        v-model="enteredLobbyID"
        class="w-full p-2 border border-gray-300 rounded mb-4"
        placeholder="Enter Lobby ID"
        :disabled="isLoading"
      />
    </UiModal>

    <!-- Simple Loading Overlay -->
    <div
      v-if="isLoading"
      class="fixed inset-0 flex items-center justify-center bg-black bg-opacity-50 z-50"
    >
      <div class="bg-white p-8 rounded-lg shadow-lg flex items-center gap-4">
        <div class="animate-spin rounded-full h-6 w-6 border-b-2 border-blue-500"></div>
        <span>Loading...</span>
      </div>
    </div>
  </div>
</template>

<style scoped></style>
