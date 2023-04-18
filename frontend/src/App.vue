<script lang="ts" setup>
import { computed, defineComponent, h, onMounted, provide, ref } from "vue";
import { RouterView } from "vue-router";
import type { GlobalTheme } from "naive-ui";
import { darkTheme, useLoadingBar, useMessage, useNotification } from "naive-ui";
import { emitter } from "./utils/common";

const theme = ref<GlobalTheme | null>(null);
const color = computed(() => (theme.value === null ? "#000" : "#fff"));
const colorComment = computed(() => (theme.value === null ? "#999" : "#aaa"));

provide("theme", theme);
emitter.on("darkMode", () => {
  theme.value = darkTheme;
});

emitter.on("lightMode", () => {
  theme.value = null;
});

onMounted(() => {
  const isDarkTheme = window.matchMedia("(prefers-color-scheme: dark)"); // 是深色
  if (isDarkTheme.matches) {
    theme.value = darkTheme;
  } else {
    theme.value = null;
  }
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
      <div class="view-dark view-comment">
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
  h1,
  h2,
  h3,
  h4,
  h5,
  h6 {
    color: v-bind(color);
  }
}

.view-comment {
  .ant-comment-content-author-name > * {
    color: v-bind(colorComment);
  }
}
</style>
