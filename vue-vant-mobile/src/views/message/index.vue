<template>
  <div>
    <lineChart />
    <barChart />
    <pieChart />
    <drag />
  </div>
</template>

<script setup lang="ts">
  import lineChart from './lineChart.vue';
  import barChart from './barChart.vue';
  import pieChart from './pieChart.vue';
  import drag from './drag.vue';
</script>

<style scoped></style>
