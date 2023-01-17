<script setup lang="ts">
import { h, defineComponent } from "vue";
import { RouterView } from "vue-router";
import { NNotificationProvider, NMessageProvider, NLoadingBarProvider, useLoadingBar, useMessage, useNotification } from "naive-ui";
// window.$message = useMessage();
// window.$notification = useNotification();

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
    <div>
      <NNotificationProvider>
        <NMessageProvider>
          <router-view />
          <NaiveProviderContent />
        </NMessageProvider>
      </NNotificationProvider>
    </div>
  </NLoadingBarProvider>
</template>
