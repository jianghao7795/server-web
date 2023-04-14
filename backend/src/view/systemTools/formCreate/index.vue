<template>
  <div style="height: 80vh">
    <iframe width="100%" height="100%" :src="`${basePrefix}/form-generator/`" frameborder="0" />
  </div>
</template>

<script>
export default {
  name: "FormGenerator",
};
</script>

<script setup>
import { ref } from "vue";
const basePrefix = ref(import.meta.env.VITE_BASE_API);
</script>
