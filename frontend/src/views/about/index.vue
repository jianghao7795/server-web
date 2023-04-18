<template>
  <div class="view-content">
    <h4>介绍: {{ userStore.currentUser.user.introduction }}</h4>
    <p v-if="userStore.currentUser.user.ID !== 0">
      描述：{{ userStore.currentUser.user.content }}
    </p>
    <Image :is-update="true"></Image>
  </div>
</template>

<script lang="ts">
export default {
  name: "About",
};
</script>

<script lang="ts" setup>
import Image from "./Image.vue";
import { useUserStore } from "@/stores/user";
const userStore = useUserStore();
</script>
