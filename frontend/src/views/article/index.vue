<template>
  <div class="article-list">
    <n-list clickable hoverable v-bind:style="{ marginButton: 10 }">
      <n-list-item v-for="item in article.list" :key="item.ID" @click="changeUrl(item.ID)">
        <n-thing content-style="margin-top: 10px;">
          <template #header>
            <div>
              <h1>{{ item.title }}</h1>
            </div>
          </template>
          <template #description>
            <div>
              简述:
              <span>{{ item.desc }}</span>
              <div>阅读量: {{ item.reading_quantity }}</div>
              <div>发布于：{{ calculationTime(item.CreatedAt) }}</div>
            </div>
          </template>
          <!-- <md-editor v-model="item.content" preview-only /> -->
          <template #footer>
            <n-space size="small">
              <n-tag v-for="i in item.tags" :key="i.ID" :bordered="false" :type="colorIndex(i.ID)" size="small">
                {{ i.name }}
              </n-tag>
            </n-space>
          </template>
        </n-thing>
      </n-list-item>
    </n-list>
    <div class="article-list-more">
      <n-button block secondary strong v-show="article.showMore" @click="changePage">{{ isMore }}</n-button>
    </div>
  </div>
</template>

<script lang="ts">
export default {
  name: "Article",
};
</script>

<script lang="ts" setup>
// import { NList, NThing, NListItem, NSpace, NTag, NButton } from "naive-ui";
import { onMounted, computed, h } from "vue";
import { useRouter } from "vue-router";
import { colorIndex } from "@/common/article";
import { useArticleStore } from "@/stores/article";
import { calculationTime } from "@/utils/date";

const article = useArticleStore();
const router = useRouter();

const isMore = computed(() => {
  if (article.loading) {
    return "加载中...";
  }
  return "查看更多";
});

const changeUrl = (id: number) => {
  router.push(`/articles/${id}`);
};

const changePage = () => {
  article.getList({ pageSize: 10 });
};

onMounted(async () => {
  try {
    await article.getList({ pageSize: 10 });
  } catch (e) {
    console.log(e);
  }
});
</script>

<style lang="less" scoped>
.article-list {
  margin: auto 25%;
}

.article-list-more {
  width: 100%;
  text-align: center;
  border: 1px;
}

.page {
  margin: 50px;
  display: flex;
  justify-content: flex-end;
}
</style>
