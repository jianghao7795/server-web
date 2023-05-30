<script setup lang="ts">
const props = defineProps();
console.log(props);
// import WelcomeItem from "./components/WelcomeItem.vue";
import { Button } from "vant";
</script>

<template>
  <div><router-view /></div>
</template>

<style scoped></style>
