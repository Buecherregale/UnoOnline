<script setup>
import {postPlayerUUID} from "~/util/postPlayerUUID.ts";

let name = ref("")

async function handleSubmit() {
  if (name.value.trim()) {
    alert(`Hello, ${name.value}!`); // You can replace this with any action
    const playerUUID = await postPlayerUUID(name.value, "localhost");
    const playerUUIDCookie = useCookie('playerUUID');
    playerUUIDCookie.value = playerUUID;
  } else {
    alert('Please enter a valid name.');
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
