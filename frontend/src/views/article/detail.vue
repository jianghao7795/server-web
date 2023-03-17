<template>
  <div class="view-content">
    <div class="view-margin">
      <h1 class="view-center">{{ articleStore.detail?.title }}</h1>
      <h3>简介：{{ articleStore.detail?.title }}</h3>
      <h4>
        <NSpace style="width: 80%">
          标签：
          <n-tag size="small" round v-for="(item, index) in articleStore.detail?.tags" :type="colorIndex(index)">{{
            item.name }}</n-tag>
        </NSpace>
      </h4>
      <!-- <div>作者：{{ detail?.user?.nick_name }}</div> -->
      <div>日期：{{ changeDate(articleStore.detail?.CreatedAt) }}</div>
      <n-divider />
      <MdEditor :style="{ width: '100%' }" :model-value="articleStore.detail?.content" :theme="theme ? 'dark' : 'light'"
        :pageFullscreen="true" :previewOnly="true"></MdEditor>
    </div>
    <n-input placeholder="评论" type="textarea" size="small" :autosize="{
      minRows: 3,
      maxRows: 5
    }" />
    <a-comment>
      <template #actions>
        <span key="comment-nested-reply-to">Reply to</span>
      </template>
      <template #author>
        <a>Han Solo</a>
      </template>
      <template #avatar>
        <a-avatar src="https://joeschmoe.io/api/v1/random" alt="Han Solo" />
      </template>
      <template #content>
        <p>
          We supply a series of design principles, practical patterns and high quality design
          resources (Sketch and Axure).
        </p>
      </template>
      <a-comment>
        <template #actions>
          <span>Reply to</span>
        </template>
        <template #author>
          <a>Han Solo</a>
        </template>
        <template #avatar>
          <a-avatar src="https://joeschmoe.io/api/v1/random" alt="Han Solo" />
        </template>
        <template #content>
          <p>
            We supply a series of design principles, practical patterns and high quality design
            resources (Sketch and Axure).
          </p>
        </template>
        <a-comment>
          <template #actions>
            <span>Reply to</span>
          </template>
          <template #author>
            <a>Han Solo</a>
          </template>
          <template #avatar>
            <a-avatar src="https://joeschmoe.io/api/v1/random" alt="Han Solo" />
          </template>
          <template #content>
            <p>
              We supply a series of design principles, practical patterns and high quality design
              resources (Sketch and Axure).
            </p>
          </template>
        </a-comment>
        <a-comment>
          <template #actions>
            <span>Reply to</span>
          </template>
          <template #author>
            <a>Han Solo</a>
          </template>
          <template #avatar>
            <a-avatar src="https://joeschmoe.io/api/v1/random" alt="Han Solo" />
          </template>
          <template #content>
            <p>
              We supply a series of design principles, practical patterns and high quality design
              resources (Sketch and Axure).
            </p>
          </template>
        </a-comment>
      </a-comment>
    </a-comment>
  </div>
</template>

<script lang="ts">
export default {
  name: "ArticleDetail",
};
</script>

<script lang="ts" setup>
// import { NSpace, NTag, NDivider } from "naive-ui";
import { onMounted, inject } from "vue";
import { useRoute } from "vue-router";
import { colorIndex } from "@/common/article";
import dayjs from "dayjs";
import MdEditor from "md-editor-v3";
import "md-editor-v3/lib/style.css";
import { useArticleStore } from "@/stores/article";

const route = useRoute();
const theme = inject("theme");
const articleStore = useArticleStore();

const changeDate = (timeData?: string): string => {
  return !!timeData ? dayjs(timeData).format("YYYY-MM-DD") : "";
};

onMounted(async () => {
  const params = route.params;
  articleStore.getDetail({ id: params.id as string });
  // console.log(route.params, params);
  //   console.log(resp);
});
</script>
