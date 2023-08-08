<template>
  <div>
    <van-action-sheet v-model:show="show" title="隐私政策" teleport="body">
      <div>
        <p>内容</p>
        <p>内容</p>
        <p>内容</p>
        <p>内容</p>
        <p>内容</p>
        <p>内容</p>
        <p>内容</p>
        <p>内容</p>
        <p>内容</p>
        <p>内容</p>
        <p>内容</p>
        <p>内容</p>
        <p>内容</p>
        <p>内容</p>
        <p>内容</p>
        <p>内容</p>
        <p>内容</p>
      </div>
    </van-action-sheet>
  </div>
</template>

<script lang="ts" setup>
  import { computed } from 'vue';
  const props = defineProps<{
    show: boolean;
  }>();

  const emit = defineEmits<{
    (e: 'handleView', value: boolean): void;
  }>();
  const show = computed({
    get: () => {
      return props.show;
    },
    set: (value: boolean) => {
      emit('handleView', value);
    },
  });
</script>

<style lang="scss" scoped></style>
