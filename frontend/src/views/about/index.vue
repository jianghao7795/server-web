<template>
  <div class="view-content">
    <h1 class="view-center">关于我</h1>
    <div>
      <Image />
    </div>
    <p>{{ aboutMeMessage?.aboutMe }}</p>
  </div>
</template>

<script lang="ts" name="About" setup>
import Image from "./Image.vue";
import { aboutMe } from "@/services/about";
import { onMounted, ref } from "vue";

const aboutMeMessage = ref<API.AboutMe>();

onMounted(async () => {
  const resp = await aboutMe();
  // console.log(resp);
  if (resp?.code === 0) {
    aboutMeMessage.value = resp.data;
  }
});
</script>
