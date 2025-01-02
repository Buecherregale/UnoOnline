<script setup lang="ts">
let name = ref("")

const playerFetches = async (name: string) => {
  try {
    return await $fetch('/api/playerID', {
          method: 'POST',
          body: {
            name: name,
          },
        }
    );
  } catch (error) {
    console.error('Error communicating with internal API:', error);
    throw createError({
      statusCode: 500,
      message: 'Failed to communicate with internal API',
    });
  }
}

async function handleSubmit() {
  if (name.value.trim()) {
    const playerUUID = await playerFetches(name.value);
    const playerUUIDCookie = useCookie('playerUUID');
    playerUUIDCookie.value = playerUUID;
    navigateTo("/hostOrJoin")
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
