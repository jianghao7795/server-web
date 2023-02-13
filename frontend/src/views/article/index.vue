<template>
  <div class="article-list">
    <n-list hoverable clickable>
      <n-list-item v-for="item in data" :key="item.ID" @click="changeUrl(item.ID)">
        <n-thing content-style="margin-top: 10px;">
          <template #header>
            <div>
              <h1>{{ item.title }}</h1>
            </div>
          </template>
          <template #description>
            <div>
              简述:
              <b>{{ item.desc }}</b>
            </div>
          </template>
          <!-- <md-editor v-model="item.content" preview-only /> -->
          <template #footer>
            <n-space size="small" style="margin-top: 4px">
              <n-tag :bordered="false" size="small" v-for="i in item.tags" :key="i.ID" :type="colorIndex(i.ID)">
                {{ i.name }}
              </n-tag>
            </n-space>
          </template>
        </n-thing>
      </n-list-item>
    </n-list>
    <div class="pageNext">
      <n-space justify="space-between">
        <n-button v-show="page !== 1" icon-placement="left">
          <template #icon>
            <right theme="outline" size="24" fill="#333" />
          </template>
          上一页
        </n-button>
        <n-button icon-placement="right" v-show="articleLength === 10">
          下一页
          <template #icon>
            <right theme="outline" size="24" fill="#333" />
          </template>
        </n-button>
      </n-space>
    </div>
  </div>
</template>

<script lang="ts">
export default {
  name: "Article",
};
</script>

<script setup lang="ts">
import { NList, NThing, NListItem, NSpace, NTag, NButton } from "naive-ui";
import { ref, onMounted, computed } from "vue";
import { getArticleList } from "@/services/article";
import type { API } from "@/type/article";
import { Right } from "@icon-park/vue-next";
import { useRouter } from "vue-router";
import { colorIndex } from "@/common/article";

const router = useRouter();
const data = ref<API.Article[]>([]);
const page = ref<number>(1);
const articleLength = computed(() => data.value.length);

const changeUrl = (id: number) => {
  router.push(`/articles/${id}`);
};

onMounted(async () => {
  try {
    const response = await getArticleList({ page: 1 });
    if (response?.code === 0) {
      data.value = (response.data?.list as API.Article[]) || [];
    }
  } catch (e) {
    console.log(e);
  }
});
</script>

<style scoped lang="less">
.article-list {
  margin: auto 25%;
}
.pageNext {
  margin: 15px 0 45px 0;
}
</style>
