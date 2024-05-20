<template>
  <div class="feeds">
    <div class="title text-2xl text-left p-5">Current Feeds</div>
    <div class="feeds-list">
      <div
        v-for="feed in feedItems"
        :key="feed.id"
        class="feed-item border-b p-5 flex items-center justify-between"
      >
        <div class="flex items-start">
          <picture class="feed-image mr-5">
            <img :src="feed.imageURL" alt="Feed Image" />
          </picture>
          <div>
            <div class="feed-title text-xl mb-1">{{ feed.title }}</div>
            <div class="feed-author mb-5">
              {{ feed.author }} - Published: {{ feed.publishedAt }}
            </div>
            <UButton>
              <ULink :href="feed.url" class="feed-url">Read More</ULink>
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
  if (!data.value || !data.value?.success) return;

  if (data.value) feedItems.value = data.value.data;
};

fetchFeedItems();
</script>
