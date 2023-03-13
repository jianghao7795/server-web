<script lang="ts" setup>
import { h, defineComponent, ref, onMounted, provide } from "vue";
import { RouterView } from "vue-router";
import {
  NNotificationProvider,
  NMessageProvider,
  NLoadingBarProvider,
  useLoadingBar,
  useMessage,
  useNotification,
  NConfigProvider,
  darkTheme,
} from "naive-ui";
import type { GlobalTheme } from "naive-ui";
import { emitter } from "./utils/common";
// window.$message = useMessage();
// window.$notification = useNotification();

const theme = ref<GlobalTheme | null>(null);

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
      <div>
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
