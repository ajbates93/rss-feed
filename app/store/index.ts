import { defineStore } from "pinia";
import type { Feed, FeedItem, MetaApiResponse } from "../types";

export const useStore = defineStore("store", {
  state: () => ({
    feeds: [] as Feed[],
    feedItems: [] as FeedItem[],
    feedItemsMeta: {
      page: 1,
      limit: 10,
      items: 0,
      total: 0,
    } as MetaApiResponse,
    page: 1,
    limit: 10,
  }),

  actions: {
    updateFeedItems(feedItems: FeedItem[]) {
      this.feedItems = feedItems;
    },
    updateFeedItemsMeta(feedItemsMeta: MetaApiResponse) {
      this.feedItemsMeta = feedItemsMeta;
    },
  },
});
