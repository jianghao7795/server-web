<script lang="ts" setup>
import { h, defineComponent, ref, onMounted, provide, computed } from "vue";
import { RouterView } from "vue-router";
import {
  useLoadingBar,
  useMessage,
  useNotification,
  darkTheme,
} from "naive-ui";
import type { GlobalTheme } from "naive-ui";
import { emitter } from "./utils/common";

const theme = ref<GlobalTheme | null>(null);
const color = computed(() => theme.value ? '#000' : '#fff');

onMounted(() => {
  emitter.on("darkMode", () => {
    theme.value = darkTheme;
  });

  emitter.on("lightMode", () => {
    theme.value = null;
  });

  provide("theme", theme);
});

function registerNaiveTools() {
  window.$loadingBar = useLoadingBar();
  window.$message = useMessage();
  window.$notification = useNotification();
}

const NaiveProviderContent = defineComponent({
  name: "NaiveProviderContent",
  setup() {
    registerNaiveTools();
  },
  render() {
    return h("div");
  },
});
</script>

<template>
  <NLoadingBarProvider :loading-bar-style="{ loading: { height: '4px', background: '#1e80ff' } }">
    <NConfigProvider :theme="theme">
      <div class="view-dark">
        <NNotificationProvider>
          <NMessageProvider>
            <router-view />
            <NaiveProviderContent />
          </NMessageProvider>
        </NNotificationProvider>
      </div>
    </NConfigProvider>
  </NLoadingBarProvider>
</template>

<style lang="less">
.view-dark {
  h1 h2 h3 h4 h5 {
    color: v-bind(color);
  }
}
</style>