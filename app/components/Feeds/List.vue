<template>
  <div class="feeds max-w-screen-lg mx-auto">
    <div
      class="title text-2xl text-left py-3 my-5 text-orange-400 flex justify-between items-center"
    >
      <div v-if="loading" class="loading flex items-center">
        <div class="text-gray-700">Loading feeds...</div>
        <UIcon name="i-heroicons-arrow-path" class="animate-spin ml-3" />
      </div>
      <div v-if="!loading" class="relative z-20 text-gray-700">
        Displaying page
        <span class="text-orange-400">{{ store.feedItemsMeta.page }}</span>
        of
        <span class="text-orange-400">{{
          store.feedItemsMeta.totalPages
        }}</span>
        page<span v-if="store.feedItemsMeta.totalPages > 1">s</span>
      </div>
      <FeedsPagination />
    </div>
    <div class="feeds-list">
      <div
        v-for="feed in store.feedItems"
        :key="feed.id"
        class="feed-item border rounded-xl mb-10 overflow-hidden flex items-center justify-between bg-white shadow-sm"
      >
        <div class="flex items-stretch">
          <div class="feed-image w-80">
            <picture class="feed-image">
              <img :src="feed.imageURL" alt="Feed Image" />
            </picture>
          </div>
          <div class="flex items-start justify-start p-10 flex-col">
            <div class="feed-author text-orange-500 italic mb-1 text-sm">
              {{ feed.author }}
            </div>
            <div class="feed-title text-xl mb-3">
              <ULink :to="feed.url">
                {{ feed.title }}
              </ULink>
            </div>

            <div class="feed-date mb-5">
              {{ feed.publishedAt }}
            </div>
            <UButton class="bg-orange-500 mt-auto">
              <ULink :to="feed.url" class="feed-url">Read More</ULink>
            </UButton>
          </div>
        </div>
      </div>
      <div v-if="!store.feedItems?.length && !loading" class="no-feeds p-5">
        No feeds found :-(
      </div>
    </div>
    <div v-if="store.feedItems?.length" class="flex items-center justify-end">
      <FeedsPagination />
    </div>
  </div>
</template>

<script lang="ts" setup>
import type { GetFeedItemsApiResponse } from "@/types";
import { useStore } from "@/store";
import { storeToRefs } from "pinia";

const store = useStore();
const { page, limit } = storeToRefs(store);
const loading = ref(false);
const errorMessage = ref<string>("");

const fetchFeedItems = async () => {
  try {
    loading.value = true;
    const response = await $fetch<GetFeedItemsApiResponse>(
      `http://localhost:4000/feed-items?page=${store.page}&limit=${store.limit}`,
    );
    if (!response || !response?.success) {
      return;
    }
    if (response && response.success) {
      store.updateFeedItems(response.data);
      store.updateFeedItemsMeta(response.meta);
    }
  } catch (e: Error) {
    errorMessage.value = e.message;
  } finally {
    loading.value = false;
  }
};

watch([page, limit], () => {
  fetchFeedItems();
});

fetchFeedItems();
</script>
