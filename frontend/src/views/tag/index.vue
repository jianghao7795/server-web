<template>
  <div class="view-content">
    <n-space>
      <n-button strong secondary round v-for="(item, index) in tagList" :key="item.ID" :type="colorIndex(index)">{{ item.name }}</n-button>
    </n-space>
  </div>
</template>

<script lang="ts" setup>
import { NSpace, NButton } from "naive-ui";
import { getTagList } from "@/services/tag";
import { onMounted, ref } from "vue";
import type { API } from "@/type/article";
import { colorIndex } from "@/common/article";

const tagList = ref<API.Tag[]>([]);
onMounted(async () => {
  const resp = await getTagList();
  tagList.value = resp?.data?.list as API.Tag[];
});
</script>
