<template>
  <div class="feeds max-w-screen-lg mx-auto">
    <div
      class="title text-2xl text-left py-3 my-5 text-orange-400 flex justify-between items-center"
    >
      <div class="relative z-20">
        Displaying {{ store.feedItemsMeta.items }} results out of
        {{ store.feedItemsMeta.total }}
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
      <div v-if="!store.feedItems?.length" class="no-feeds p-5">
        No feeds found :-(
      </div>
    </div>
    <div class="flex items-center justify-end">
      <FeedsPagination />
    </div>
  </div>
</template>

<script lang="ts" setup>
import type { GetFeedItemsApiResponse } from "@/types";
import { useStore } from "@/store";

const store = useStore();

const fetchFeedItems = async () => {
  const { data } = await useFetch<GetFeedItemsApiResponse>(
    `http://localhost:4000/feed-items?page=${store.page}&limit=${store.limit}`,
    {
      lazy: true,
    },
  );
  if (!data.value || !data.value?.success) {
    return;
  }
  if (data.value && data.value?.success) {
    store.updateFeedItems(data.value.data);
    store.updateFeedItemsMeta(data.value.meta);
  }
};

watch([store.page, store.limit], () => {
  fetchFeedItems();
});

fetchFeedItems();
</script>
