<script setup lang="ts">
import TheWelcome from '@/components/TheWelcome.vue';
import { NInput, NButton, NIcon } from 'naive-ui';
import { AddFilled } from '@vicons/material';
import { useCounterStore } from '@/stores/counter';

const counterStore = useCounterStore();
</script>

<template>
  <main>
    <TheWelcome />
    <n-input placeholder="请输入" />
    <p>{{ counterStore.counter }}</p>
    <p>{{ counterStore.doubleCount }}</p>
    <n-button v-on:click="counterStore.increment">
      <template #icon>
        <AddFilled />
      </template>
      加
    </n-button>
  </main>
</template>
