<template>
  <div class="article-list">
    <n-list clickable hoverable>
      <n-list-item
        v-for="item in article.list"
        :key="item.ID"
        @click="changeUrl(item.ID)"
      >
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
            <n-space size="small">
              <n-tag
                v-for="i in item.tags"
                :key="i.ID"
                :bordered="false"
                :type="colorIndex(i.ID)"
                size="small"
              >
                {{ i.name }}
              </n-tag>
            </n-space>
          </template>
        </n-thing>
      </n-list-item>
    </n-list>
    <div class="pageNext">
      <n-space justify="space-between">
        <n-button
          v-show="page !== 1"
          icon-placement="left"
          @click="() => changePage(false)"
        >
          <template #icon>
            <right fill="#333" size="24" theme="outline" />
          </template>
          上一页
        </n-button>
        <n-button v-show="articleLength === 10" icon-placement="right">
          下一页
          <template #icon>
            <right fill="#333" size="24" theme="outline" />
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

<script lang="ts" setup>
// import { NList, NThing, NListItem, NSpace, NTag, NButton } from "naive-ui";
import { ref, onMounted, computed } from "vue";
import { Right } from "@icon-park/vue-next";
import { useRouter } from "vue-router";
import { colorIndex } from "@/common/article";
import { useArticleStore } from "@/stores/article";

const article = useArticleStore();
const router = useRouter();
const data = ref<API.Article[]>([]);
const page = ref<number>(1);
const articleLength = computed(() => data.value.length);

const changeUrl = (id: number) => {
  router.push(`/articles/${id}`);
};

const changePage = (isPage: boolean) => {
  if (isPage) {
    page.value = page.value + 1;
  } else {
    page.value = page.value - 1;
  }
  article.getList({ page: page.value, pageSize: 10 });
};

onMounted(async () => {
  try {
    await article.getList({ page: page.value, pageSize: 10 });
  } catch (e) {
    console.log(e);
  }
});
</script>

<style lang="less" scoped>
.article-list {
  margin: auto 25%;
}

.pageNext {
  margin: 15px 0 45px 0;
}
</style>
