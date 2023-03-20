<template>
  <div class="view-content">
    <div class="view-margin">
      <h1 class="view-center"><b>{{ articleStore.detail?.title }}</b></h1>
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
    <div @click="() => focusInput(true)">
      <n-input placeholder="评论" type="textarea" size="small" :autosize="true" v-model:value="inputRef">
      </n-input>
      <div class="comment-line" v-show="isComment">
        <div></div>
        <div>
          <NButton type="primary" @click="submit">发表评论</NButton>
        </div>
      </div>
    </div>
    <a-comment v-for="items in comment" :key="items.ID">
      <template #actions>
        <div key="comment-nested-reply-to">回复</div>
        <n-input />
      </template>
      <template #author>
        <a>Han Solo</a>
      </template>
      <template #avatar>
        <a-avatar src="https://zos.alipayobjects.com/rmsportal/ODTLcjxAfvqbxHnVXCYX.png" alt="Han Solo" />
      </template>
      <template #content>
        <p>
          {{ items.content }}
        </p>
      </template>
      <a-comment v-for="item in items.children" :key="item.ID">
        <template #author>
          <a>Han Solo</a>
        </template>
        <template #avatar>
          <a-avatar src="https://zos.alipayobjects.com/rmsportal/ODTLcjxAfvqbxHnVXCYX.png" alt="Han Solo" />
        </template>
        <template #content>
          <p>
            {{ item.content }}
          </p>
        </template>

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
import { onMounted, inject, ref } from "vue";
import { useRoute } from "vue-router";
import { colorIndex } from "@/common/article";
import dayjs from "dayjs";
import MdEditor from "md-editor-v3";
import "md-editor-v3/lib/style.css";
import { useArticleStore } from "@/stores/article";
import { getArticleComment, createdComment } from "@/services/comment";
import { useMessage } from 'naive-ui';

const message = useMessage();

const route = useRoute();
const theme = inject("theme");
const articleStore = useArticleStore();
const isComment = ref<boolean>(false);

const comment = ref<Comment.comment[]>([])
const inputRef = ref<string>('')

const focusInput = (isFouse: boolean) => {
  isComment.value = isFouse;
}

const submit = async () => {
  if (inputRef.value === '') {
    message.warning("请输入评论")
    return;
  }
  // console.log(inputRef.value);
  // setTimeout(() => {
  //   inputRef.value = '';
  //   focusInput(false);
  // }, 3000);
  const resp = await createdComment({ content: inputRef.value, articleId: Number(route.params.id).valueOf(), parentId: 0 })
}

const changeDate = (timeData?: string): string => {
  return !!timeData ? dayjs(timeData).format("YYYY-MM-DD") : "";
};

onMounted(async () => {
  const params = route.params;
  articleStore.getDetail({ id: params.id as string });
  const resp = await getArticleComment({ articleId: params.id as string })
  comment.value = resp.data
  // console.log(route.params, params);
  //   console.log(resp);
});
</script>

<style scoped lang="less">
.comment-line {
  margin-top: 5px;
  display: flex;
  justify-content: space-between;
  padding-bottom: 5px;
}
</style>