<template>
  <!-- <el-button type="primary">
    <slot #[item]="data"></slot>
  </el-button> -->
  <ul>
    <li v-for="item in props.lists" :key="item.id || item.name">
      <slot name="two" :item="item"></slot>
    </li>
  </ul>
  <div><slot name="active"></slot></div>
</template>

<script setup>
const props = defineProps({
  lists: {
    type: Array,
    default: [],
  },
});
</script>

<style scoped></style>
