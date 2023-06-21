<template>
  <div class="about">
    <h1>This is an about page</h1>
    <p><Button type="primary" size="small">+page</Button></p>
  </div>
</template>

<script lang="ts" setup>
import { Button } from "vant";
</script>

<style>
@media (min-width: 1024px) {
  .about {
    min-height: 100vh;
    display: flex;
    align-items: center;
  }
}
</style>
