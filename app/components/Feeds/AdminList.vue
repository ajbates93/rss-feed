<template>
  <div class="feeds">
    <div class="title text-2xl text-left p-5">Active Feeds</div>
    <div class="feeds-list">
      <div
        v-for="feed in store.feeds"
        :key="feed.ID"
        class="feed-item border-b p-5 flex items-center justify-between"
      >
        <div>
          <div class="feed-title">{{ feed.Title }}</div>
          <div class="feed-url">{{ feed.URL }}</div>
        </div>
        <div>
          <button @click="removeFeed(feed.ID)">
            <UIcon name="i-heroicons-trash" />
          </button>
        </div>
      </div>
      <div v-if="!store.feeds.length" class="no-feeds p-5">
        No feeds found :-(
      </div>
    </div>
  </div>
</template>

<script lang="ts" setup>
import { useStore } from "@/store";
import type { RemoveFeedApiResponse } from "@/types";

const store = useStore();

type Feed = {
  ID: number;
  Title: string;
  URL: string;
};

// Fetch feeds from the server
const fetchFeeds = async () => {
  const { data } = await useFetch<Feed[]>("http://localhost:4000/feeds");
  if (!data.value) return;
  store.feeds = data.value;
};

const removeFeed = async (id: number) => {
  const { data } = await useFetch<RemoveFeedApiResponse>(
    `http://localhost:4000/feeds/${id}`,
    {
      method: "DELETE",
    },
  );
  if (data.value?.success) {
    store.feeds = store.feeds.filter((feed: Feed) => feed.ID !== id);
  }
};

fetchFeeds();
</script>
