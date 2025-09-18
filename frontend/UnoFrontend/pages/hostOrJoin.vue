<script setup lang="ts">
import type { Room, CreateRoomRequest, JoinRoomRequest } from "~/util/models";
import { saveGameStateToCookies } from "~/util/roomCookie";
import { loadPlayerFromCookie } from "~/util/playerCookie";

// Reactive state with explicit types
const showPopupHost = ref<boolean>(false);
const showPopupJoin = ref<boolean>(false);
const selectedPlayerCount = ref<number>(2);
const enteredLobbyID = ref<string>("");

/**
 * Handles host game confirmation with proper error handling
 */
async function confirmedHost(): Promise<void> {
  showPopupHost.value = false;
  const player = loadPlayerFromCookie();

  if (!player?.id) {
    throw createError({
      statusCode: 400,
      message: "Player not found in session",
    });
  }

  try {
    const requestBody: CreateRoomRequest = {
      id: player.id,
      maxPlayers: selectedPlayerCount.value,
    };

    const responseRoom: Room = await $fetch<Room>("/api/rooms", {
      method: "POST",
      body: requestBody,
    });

    // Use cookie helper instead of sessionStorage
    saveGameStateToCookies(responseRoom, true);

    useState<Room>("rooms", () => responseRoom);
    useState<boolean>("isHost", () => true);
    const roomID: string = responseRoom.id;

    await navigateTo(`/lobby-${roomID}`);
  } catch (error) {
    console.error("Error communicating with internal API:", error);
    throw createError({
      statusCode: 500,
      message: "Failed to communicate with internal API",
    });
  }
}

/**
 * Handles join game confirmation with proper error handling
 */
async function confirmedJoin(): Promise<void> {
  showPopupJoin.value = false;
  const player = loadPlayerFromCookie();

  if (!player?.id) {
    throw createError({
      statusCode: 400,
      message: "Player not found in session",
    });
  }

  console.log(enteredLobbyID.value);

  try {
    const requestBody: JoinRoomRequest = {
      id: player.id,
    };

    const responseRoom: Room = await $fetch<Room>(
      `api/rooms/${enteredLobbyID.value}/players`,
      {
        method: "POST",
        body: requestBody,
      }
    );
    console.log(responseRoom);

    saveGameStateToCookies(responseRoom, false);

    useState<Room>("rooms", () => responseRoom);
    useState<boolean>("isHost", () => false);
    const roomID: string = responseRoom.id;

    await navigateTo(`/lobby-${roomID}`);
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
  <div
    class="flex items-center justify-center min-h-screen bg-zinc-100"
    role="region"
    aria-label="Game Selection"
  >
    <div class="flex gap-5">
      <button
        class="flex flex-col items-center justify-center p-10 bg-white border border-gray-300 rounded shadow-md hover:bg-gray-100"
        aria-label="Host a Game"
        @click="showPopupHost = true"
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
        @click="showPopupJoin = true"
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
      v-if="showPopupHost"
      class="fixed inset-0 flex items-center justify-center bg-black bg-opacity-50"
    >
      <div class="bg-white p-5 rounded shadow-lg w-80">
        <h2 class="text-lg font-semibold mb-4">Select Number of Players</h2>
        <select
          v-model="selectedPlayerCount"
          class="w-full p-2 border border-gray-300 rounded mb-4"
        >
          <option v-for="n in 7" :key="n" :value="n + 1">{{ n + 1 }}</option>
        </select>
        <div class="flex justify-end gap-3">
          <button
            class="px-4 py-2 bg-gray-300 rounded hover:bg-gray-400"
            @click.stop="showPopupHost = false"
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

    <div
      v-if="showPopupJoin"
      class="fixed inset-0 flex items-center justify-center bg-black bg-opacity-50"
    >
      <div class="bg-white p-5 rounded shadow-lg w-80">
        <h2 class="text-lg font-semibold mb-4">Enter Lobby ID</h2>
        <input
          v-model="enteredLobbyID"
          class="w-full p-2 border border-gray-300 rounded mb-4"
          placeholder="Enter Lobby ID"
        />
        <div class="flex justify-end gap-3">
          <button
            class="px-4 py-2 bg-gray-300 rounded hover:bg-gray-400"
            @click.stop="showPopupJoin = false"
          >
            Cancel
          </button>
          <button
            class="px-4 py-2 bg-blue-500 text-white rounded hover:bg-blue-600"
            @click="confirmedJoin"
          >
            OK
          </button>
        </div>
      </div>
    </div>
  </div>
</template>

<style scoped></style>
