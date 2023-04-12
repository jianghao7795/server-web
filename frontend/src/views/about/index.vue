<template>
  <div class="view-content">
    <h4>介绍: {{ userStore.currentUser.user.introduction }}</h4>
    <p v-if="userStore.currentUser.user.ID !== 0">描述：{{ userStore.currentUser.user.content }}</p>
  </div>
</template>

<script lang="ts" name="About" setup>
import { useUserStore } from "@/stores/user";
const userStore = useUserStore();
</script>
