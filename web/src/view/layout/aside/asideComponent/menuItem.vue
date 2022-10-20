<template>
  <el-menu-item :index="props.routerInfo.name">
    <el-icon v-if="props.routerInfo.meta.icon">
      <component :is="props.routerInfo.meta.icon" />
    </el-icon>
    <template #title>
      <span class="gva-menu-item-title">{{ props.routerInfo.meta.title }}</span>
    </template>
  </el-menu-item>
</template>

<script>
export default {
  name: "MenuItem",
};
</script>

<script setup>
const props = defineProps({
  routerInfo: {
    default: function () {
      return null;
    },
    type: Object,
  },
});
// console.log(typeof props.routerInfo.meta.icon);
</script>

<style lang="scss" scoped>
// .gva-menu-item-title {
//   min-width: 160px;
// }
</style>
