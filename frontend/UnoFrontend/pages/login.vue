<script setup lang="ts">
import type { Player, CreatePlayerRequest } from "~/util/models";
import { savePlayerToCookie } from "~/util/playerCookie";
import { handleApiError } from "~/util/errorUtils";

// Reactive state with explicit types
const name = ref<string>("");

/**
 * Fetches player data from API with proper error handling
 */
const playerFetches = async (name: string): Promise<Player> => {
  try {
    const requestBody: CreatePlayerRequest = { name };
    return await $fetch<Player>("/api/players", {
      method: "POST",
      body: requestBody,
    });
  } catch (error) {
    handleApiError(error);
  }
};

/**
 * Handles form submission with validation
 */
async function handleSubmit(): Promise<void> {
  if (name.value.trim()) {
    const player: Player = await playerFetches(name.value);
    savePlayerToCookie(player);
    await navigateTo("/hostOrJoin");
  } else {
    alert("Please enter a valid name.");
  }
}
</script>

<template>
  <div class="name-input-component">
    <p>Please enter a name:</p>
    <input
      type="text"
      v-model="name"
      @keyup.enter="handleSubmit"
      placeholder="Type your name here"
    />
    <button @click="handleSubmit">Submit</button>
  </div>
</template>

<style scoped>
.name-input-component {
  font-family: Arial, sans-serif;
  padding: 10px;
}
input {
  margin-right: 10px;
  padding: 5px;
  font-size: 16px;
}
button {
  padding: 5px 10px;
  font-size: 16px;
  cursor: pointer;
}
</style>
