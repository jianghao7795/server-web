<template>
  <div>
    <n-card
      class="darkStyle"
      :bordered="false"
      title="瓦都尅'Blog"
      header-style="headerStyle"
    >
      <template #header-extra>
        <div class="headerStyleLine">123123</div>
      </template>
      <router-view />
    </n-card>
  </div>
</template>

<script setup lang="ts">
import { RouterView } from "vue-router";
import { NCard } from "naive-ui";
</script>

<style scoped lang="less">
.n-card > .n-card-header .n-card-header__main {
  color: #fff;
}
.headerStyle {
  padding: 15px 15px;
  height: 60px;
  line-height: 60px;
}

.headerStyleLine {
  color: #fff;
}

.darkStyle {
  background-image: url("/home-bg.png");
  height: auto;
}
</style>
