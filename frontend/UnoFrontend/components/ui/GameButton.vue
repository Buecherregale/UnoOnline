<template>
  <button
    class="flex flex-col items-center justify-center p-10 bg-white border border-gray-300 rounded shadow-md hover:bg-gray-100 transition-colors"
    :aria-label="ariaLabel"
    @click="$emit('click')"
  >
    <img :src="iconSrc" class="w-16 h-16 mb-4" :alt="iconAlt" />
    <span class="text-lg font-semibold">{{ label }}</span>
  </button>
</template>

<script setup lang="ts">
interface Props {
  label: string;
  iconSrc: string;
  iconAlt: string;
  ariaLabel: string;
}

interface Emits {
  click: [];
}

defineProps<Props>();
defineEmits<Emits>();
</script>
