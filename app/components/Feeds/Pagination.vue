<template>
  <div
    class="relative z-20 pagination text-lg flex items-center justify-end gap-2"
  >
    <div class="flex items-center gap-2 mr-5">
      <span class="text-sm">Results to show:</span>
      <USelect v-model="store.limit" :options="[10, 20, 50]" class="rounded" />
    </div>
    <UButton
      @click="store.page--"
      :disabled="store.page === 1"
      class="!bg-white text-orange-500 border p-1.5 rounded"
    >
      Previous
    </UButton>
    <UButton
      @click="store.page++"
      :disabled="store.feedItemsMeta.items < store.feedItemsMeta.limit"
      class="!bg-white text-orange-500 border p-1.5 rounded"
    >
      Next
    </UButton>
  </div>
</template>

<script lang="ts" setup>
import { useStore } from "@/store";

const store = useStore();
</script>
