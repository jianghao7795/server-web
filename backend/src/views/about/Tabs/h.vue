<template>
  <div>
    <slot name="active"></slot>
    <div><slot :text="greetingMessage" :count="1"></slot></div>
    <span>{{ props.sum }}</span>
  </div>
</template>

<script setup lang="ts">
import { ref } from "vue";
const props = defineProps({
  sum: {
    type: Number,
    default: 1,
  },
});

const greetingMessage = ref<string>("greetingMessage");
</script>

<style scoped></style>
