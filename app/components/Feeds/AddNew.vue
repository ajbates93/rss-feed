<template>
  <div class="p-5 mb-5 border rounded-xl bg-gray-50">
    <div class="text-2xl mb-5">Add New Feed</div>
    <form @submit.prevent="onSubmit">
      <div class="flex flex-col mb-5">
        <label for="name" class="text-lg inline-block mb-1">Name</label>
        <input
          type="text"
          id="name"
          v-model="name"
          class="border p-2 rounded"
        />
      </div>
      <div class="flex flex-col">
        <label for="url" class="text-lg inline-block mb-1">URL</label>
        <input type="text" id="url" v-model="url" class="border p-2 rounded" />
      </div>
      <button type="submit" class="mt-5 bg-orange-500 text-white p-2 rounded">
        Add Feed
      </button>
    </form>
  </div>
</template>

<script lang="ts" setup>
import { ref } from "vue";
import type { AddNewFeedApiResponse } from "../../types";
import { useStore } from "@/store";

const store = useStore();

const url = ref("");
const name = ref("");

const onSubmit = async () => {
  const { data } = await useFetch<AddNewFeedApiResponse>(
    "http://localhost:4000/feeds",
    {
      method: "POST",
      body: JSON.stringify({ Title: name.value, URL: url.value }),
      headers: {
        "Content-Type": "application/json",
      },
    },
  );

  if (data.value?.success) {
    // Push to store
    store.feeds.push(data.value?.data);
  }
};
</script>
