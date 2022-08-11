<script setup lang="ts">
import TheWelcome from '@/components/TheWelcome.vue';
import { NInputNumber, NButton, NIcon } from 'naive-ui';
// import { ref } from 'vue';
import { AddFilled } from '@vicons/material';
import { useCounterStore } from '@/stores/counter';
import { useLoading } from '@/hooks';

const counterStore = useCounterStore();
// const loadingRef = ref<Boolean>(false);
const { loading, startLoading, endLoading } = useLoading();
const handlAdd = () => {
  startLoading();
  setTimeout(() => {
    counterStore.increment();
    endLoading();
  }, 1000);
};
</script>

<template>
  <main>
    <TheWelcome />
    <n-input-number placeholder="请输入" v-model:value="counterStore.counter" />
    <p>{{ counterStore.counter }}</p>
    <p>{{ counterStore.doubleCount }}</p>
    <n-button :loading="loading" type="primary" v-on:click="handlAdd">
      <template #icon>
        <n-icon>
          <AddFilled />
        </n-icon>
      </template>
      加
    </n-button>
  </main>
</template>
