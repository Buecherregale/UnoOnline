<template>
  <Teleport to="body">
    <div
      v-if="isOpen"
      class="fixed inset-0 flex items-center justify-center bg-black bg-opacity-50 z-50"
      @click.self="$emit('close')"
    >
      <div class="bg-white p-5 rounded shadow-lg w-80 max-w-[90vw]">
        <header v-if="title" class="mb-4">
          <h2 class="text-lg font-semibold">{{ title }}</h2>
        </header>

        <main class="mb-4">
          <slot />
        </main>

        <footer class="flex justify-end gap-3">
          <button
            v-if="showCancel"
            class="px-4 py-2 bg-gray-300 rounded hover:bg-gray-400 transition-colors"
            @click="$emit('close')"
          >
            {{ cancelText }}
          </button>
          <button
            class="px-4 py-2 bg-blue-500 text-white rounded hover:bg-blue-600 transition-colors disabled:opacity-50"
            :disabled="confirmText!.includes('...')"
            @click="$emit('confirm')"
          >
            {{ confirmText }}
          </button>
        </footer>
      </div>
    </div>
  </Teleport>
</template>

<script setup lang="ts">
interface Props {
  isOpen: boolean;
  title?: string;
  showCancel?: boolean;
  cancelText?: string;
  confirmText?: string;
}

interface Emits {
  close: [];
  confirm: [];
}

withDefaults(defineProps<Props>(), {
  title: "",
  showCancel: true,
  cancelText: "Cancel",
  confirmText: "OK",
});

defineEmits<Emits>();
</script>
