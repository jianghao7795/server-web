<template>
  <div>{{ props.imgUrl }}</div>
</template>

<script lang="ts" setup>
import { onMounted, computed } from "vue";

const props = defineProps<{ imgUrl: string }>();

onMounted(() => {
  console.log("234234234234");
});
</script>
