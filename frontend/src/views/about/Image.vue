<template>
  <div>{{ isMoreUpdate ? "哦豁" : "" }}</div>
</template>

<script lang="ts" setup>
import { onMounted, computed } from "vue";

const props = defineProps<{ isUpdate: boolean }>();

const isMoreUpdate = computed(() => props.isUpdate);
onMounted(() => {
  console.log("234234234234");
});
</script>
