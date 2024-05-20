import { defineStore } from "pinia";
import { ref } from "vue";
import type { Feed, FeedItem } from "../types";

export const useStore = defineStore("", () => {
  state: () => ({
    feeds: [] as Feed[],
    feedItems: [] as FeedItem[],
  });

  actions: {
    updateFeedItems(feedItems: FeedItem[]) {
      this.feedItems = feedItems;
    }
  }
});
