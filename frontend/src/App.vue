<script setup lang="ts">
import { RouterView } from "vue-router";
import {
  NNotificationProvider,
  NMessageProvider,
  useMessage,
  useNotification,
} from "naive-ui";

// window.$message = useMessage();
// window.$notification = useNotification();
</script>

<template>
  <div>
    <NNotificationProvider>
      <NMessageProvider>
        <router-view />
      </NMessageProvider>
    </NNotificationProvider>
  </div>
</template>
