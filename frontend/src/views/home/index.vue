<template>
  <div class="article-list">
    <n-list hoverable clickable>
      <n-list-item>
        <n-thing title="相见恨晚" content-style="margin-top: 10px;">
          <template #description>
            <n-space size="small" style="margin-top: 4px">
              <n-tag :bordered="false" type="info" size="small"> 暑夜 </n-tag>
              <n-tag :bordered="false" type="info" size="small"> 晚春 </n-tag>
            </n-space>
          </template>
          奋勇呀然后休息呀<br />
          完成你伟大的人生
        </n-thing>
      </n-list-item>
      <n-list-item>
        <n-thing content-style="margin-top: 10px;">
          <template #header>
            <div><h1>123123123</h1></div>
          </template>
          <template #description>
            <n-space size="small" style="margin-top: 4px">
              <n-tag :bordered="false" type="info" size="small"> 环形公路 </n-tag>
              <n-tag :bordered="false" type="info" size="small"> 潜水艇司机 </n-tag>
            </n-space>
          </template>
          最新的打印机<br />
          复制着彩色傀儡<br />
          早上好我的罐头先生<br />
          让他带你去被工厂敲击
        </n-thing>
      </n-list-item>
    </n-list>
  </div>
</template>

<script setup lang="ts">
import { NList, NThing, NListItem, NSpace, NTag } from "naive-ui";
import { ref, onMounted } from "vue";
import { getArticleList } from "@/services/article";

const data = ref([]);

onMounted(async () => {
  const response = await getArticleList();
});
</script>

<style scoped lang="less">
.article-list {
  margin: auto 25%;
}
</style>
