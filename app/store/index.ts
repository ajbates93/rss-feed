import { defineStore } from "pinia";
import { ref } from "vue";
import type { Feed } from "../types";

export const useStore = defineStore("", () => {
  const feeds = ref<Feed[]>([]);

  return {
    feeds,
  };
});
