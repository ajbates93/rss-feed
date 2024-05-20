<template>
  <div class="feeds max-w-screen-lg mx-auto">
    <div class="title text-2xl text-left p-5">Current Feeds</div>
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
            <UButton>
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

const feedItems = ref<FeedItem[]>([]);

const fetchFeedItems = async () => {
  const { data } = await useFetch<GetFeedItemsApiResponse>(
    "http://localhost:4000/feed-items",
    {
      lazy: true,
    },
  );
  if (!data.value || !data.value?.success) {
    return;
  }
  if (data.value) {
    feedItems.value = data.value.data;
  }
};

fetchFeedItems();
</script>
