<template>
  <div class="feeds max-w-screen-lg mx-auto">
    <div
      class="title text-2xl text-left py-3 my-5 text-orange-400 flex justify-between items-center"
    >
      <div class="relative z-20">Latest Feeds</div>
      <div class="relative z-20 pagination text-lg">
        <UButton
          @click="page--"
          :disabled="page === 1"
          class="!bg-white text-orange-500 border p-2 rounded mr-2"
        >
          Previous
        </UButton>
        <UButton
          @click="page++"
          class="!bg-white text-orange-500 border p-2 rounded"
        >
          Next
        </UButton>
        <USelect v-model="limit" :options="[10, 20, 50]" />
      </div>
    </div>
    <div class="feeds-list">
      <div
        v-for="feed in feedItems"
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
      <div v-if="!feedItems?.length" class="no-feeds p-5">
        No feeds found :-(
      </div>
    </div>
  </div>
</template>

<script lang="ts" setup>
import type { FeedItem, GetFeedItemsApiResponse } from "@/types";
import { useStore } from "@/store";

const { feedItems } = useStore();
const page = ref<number>(1);
const limit = ref<number>(10);

const fetchFeedItems = async () => {
  const { data } = await useFetch<GetFeedItemsApiResponse>(
    `http://localhost:4000/feed-items?page=${page.value}&limit=${limit.value}`,
    {
      lazy: true,
    },
  );
  if (!data.value || !data.value?.success) {
    return;
  }
  if (data.value && data.value?.success) {
    feedItems = data.value.data;
  }
};

watch([page, limit], () => {
  fetchFeedItems();
});

fetchFeedItems();
</script>
