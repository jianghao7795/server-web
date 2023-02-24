import { defineStore } from "pinia";
import { ref } from "vue";

export const useRefreshStore = defineStore("refresh", () => {
  const isRefresh = ref(true);
  function changIsRefresh(val = true) {
    isRefresh.value = val;
  }

  return { isRefresh, changIsRefresh };
});
