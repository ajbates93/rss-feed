<template>
  <div class="feeds">
    <div class="title text-2xl text-left p-5">Active Feeds</div>
    <div class="feeds-list">
      <div
        v-for="feed in store.feeds"
        :key="feed.ID"
        class="feed-item border-b p-5"
      >
        <div class="feed-title">{{ feed.Title }}</div>
        <div class="feed-url">{{ feed.URL }}</div>
      </div>
    </div>
  </div>
</template>

<script lang="ts" setup>
import { useStore } from "@/store";

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

fetchFeeds();
</script>
