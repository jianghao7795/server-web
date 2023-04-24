<template>
  <div>
    <slot name="active"></slot>
    <div><slot /></div>
    <span>{{ props.sum }}</span>
  </div>
</template>

<script setup>
const props = defineProps({
  sum: {
    type: Number,
    default: 1,
  },
});
</script>

<style scoped></style>
