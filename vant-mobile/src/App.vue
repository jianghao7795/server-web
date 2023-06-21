<script setup lang="ts">
import { ConfigProvider } from "vant";
</script>

<template>
  <div>
    <ConfigProvider theme="dark"><router-view /></ConfigProvider>
  </div>
</template>

<style scoped></style>
