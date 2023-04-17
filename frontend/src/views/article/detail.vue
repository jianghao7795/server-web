<template>
  <div class="view-content">
    <div class="view-margin">
      <h1 class="view-center">
        <b>{{ articleStore.detail?.title }}</b>
      </h1>
      <h3>简介：{{ articleStore.detail?.title }}</h3>
      <h4>
        <NSpace style="width: 80%">
          标签：&#129409;
          <n-tag size="small" round v-for="(item, index) in articleStore.detail?.tags" :type="colorIndex(index)">
            {{ item.name }}
          </n-tag>
        </NSpace>
      </h4>
      <!-- <div>作者：{{ detail?.user?.nick_name }}</div> -->
      <NSpace vertical>
        <div>日期：{{ changeDate(articleStore.detail?.CreatedAt) }}</div>
        <div>阅读量：{{ articleStore.detail?.reading_quantity }}</div>
      </NSpace>

      <n-divider />
      <MdEditor
        :style="{ width: '100%' }"
        :model-value="articleStore.detail?.content"
        :theme="theme ? 'dark' : 'light'"
        :pageFullscreen="true"
        :previewOnly="true"
      ></MdEditor>
    </div>
    <div>
      <n-input
        @click="focusInput(true)"
        placeholder="留言评论"
        type="textarea"
        size="small"
        :autosize="true"
        v-model:value="inputRef"
        ref="inputSelectRef"
      ></n-input>
      <div class="comment-line" v-show="isComment">
        <div><Emoji @changeActive="changeActive" @handleFouces="handleFouces" /></div>
        <div>
          <NButton type="primary" @click="submit(0, comment)">发表评论</NButton>
        </div>
      </div>
    </div>
    <n-card :title="`全部评论 (${comment.length})`" :bordered="false">
      <n-empty description="快点评论吧！" v-show="comment.length === 0"></n-empty>
      <a-comment v-for="items in comment" :key="items.ID">
        <template #actions>
          <a key="comment-nested-reply-to" v-if="!isCommentChildren[items.ID]" @click="() => reply(items.ID, true)">
            回复
          </a>
          <n-input-group v-else="isCommentChildren[items.ID]">
            <n-input v-model:value="inputChildren" placeholder="评论" autofocus />
            <n-button type="primary" ghost @click="submit(items.ID, items.children)">回复</n-button>
          </n-input-group>
        </template>
        <template #author>
          <a>{{ items.user_name }}</a>
        </template>
        <template #avatar>
          <a-avatar
            v-if="items.user_id === 0"
            src="https://zos.alipayobjects.com/rmsportal/ODTLcjxAfvqbxHnVXCYX.png"
            alt="Han Solo"
          />
          <a-avatar v-else :src="items.user.header ? `${Base_URL}/${items.user.header}` : ''" alt="Han Solo" />
        </template>
        <template #content>
          <div>
            <div v-html="items.content"></div>
          </div>
        </template>
        <a-comment v-for="item in items.children" :key="item.ID">
          <template #author>
            <a>{{ item.user_name }}</a>
          </template>
          <template #avatar>
            <a-avatar
              v-if="item.user_id === 0"
              src="https://zos.alipayobjects.com/rmsportal/ODTLcjxAfvqbxHnVXCYX.png"
              alt="Han Solo"
            />
            <a-avatar v-else :src="item.user.header ? `${Base_URL}/${item.user.header}` : ''" alt="Han Solo" />
          </template>
          <template #content>
            <div>
              <div v-html="item.content"></div>
            </div>
          </template>
        </a-comment>
      </a-comment>
    </n-card>
  </div>
</template>

<script lang="ts">
export default {
  name: "ArticleDetail",
};
</script>

<script lang="ts" setup>
import { onMounted, inject, ref, type Ref } from "vue";
import { useRoute } from "vue-router";
import { colorIndex } from "@/common/article";
import dayjs from "dayjs";
import MdEditor from "md-editor-v3";
import "md-editor-v3/lib/style.css";
import { useArticleStore } from "@/stores/article";
import { getArticleComment, createdComment } from "@/services/comment";
import { useMessage, type GlobalTheme } from "naive-ui";
import { useUserStore } from "@/stores/user";
import Emoji from "@/components/emoji/index.vue";

const Base_URL = ref<string>(import.meta.env.VITE_BASE_API);

const inputSelectRef = ref();
const message = useMessage();
const userStroe = useUserStore();
const route = useRoute();
const theme = inject<Ref<GlobalTheme | null>>("theme");
const articleStore = useArticleStore();
const isComment = ref<boolean>(false);
const isCommentChildren = ref<{ [id: number]: boolean }>({});

const comment = ref<Comment.comment[]>([]);
const inputRef = ref<string>("");
const inputChildren = ref<string>("");
const changeLogin = inject<(status: boolean) => void>("changeLogin");

const focusInput = (isFouse: boolean) => {
  isComment.value = isFouse;
};

const changeActive = (n: EmojiType.emoji) => {
  inputSelectRef.value.focus();
  inputRef.value = inputRef.value + n.char;
};

const handleFouces = () => {
  inputSelectRef.value.focus();
};

const reply = (id: number, status: boolean) => {
  isCommentChildren.value = { [id]: status };
};
const submit = async (parentId: number, children: Comment.comment[]) => {
  if (userStroe.currentUser.user.ID === 0) {
    message.warning("请登录");
    changeLogin?.(true);
    return;
  }
  if (parentId === 0 && inputRef.value === "") {
    message.warning("请输入评论");
    return;
  }
  if (parentId !== 0 && inputChildren.value === "") {
    message.warning("请输入评论");
    return;
  }
  // console.log(inputRef.value);
  // setTimeout(() => {
  //   inputRef.value = '';
  //   focusInput(false);
  // }, 3000);
  const resp = await createdComment({
    content: inputRef.value || inputChildren.value,
    articleId: Number(route.params.id).valueOf(),
    parentId,
    user_id: userStroe.currentUser.user.ID || 0,
    user_name: userStroe.currentUser.user.name || "测试",
  } as Comment.comment);
  if (resp?.code === 0) {
    message.success("评论成功");
    setTimeout(() => {
      let child = undefined;
      if (parentId === 0) {
        child = [];
      }
      isComment.value = false;
      isCommentChildren.value[parentId] = false;
      children.unshift({
        user: userStroe.currentUser.user,
        content: inputRef.value || inputChildren.value,
        articleId: Number(route.params.id).valueOf(),
        parentId,
        user_id: userStroe.currentUser.user.ID || 0,
        user_name: userStroe.currentUser.user.name || "测试",
        ID: resp.data.id,
        children: child,
      } as Comment.comment);
      inputRef.value = "";
      inputChildren.value = "";
    }, 1000);
  }
};

const changeDate = (timeData?: string): string => {
  return !!timeData ? dayjs(timeData).format("YYYY-MM-DD") : "";
};

onMounted(async () => {
  const params = route.params;
  articleStore.getDetail({ id: params.id as string });
  const resp = await getArticleComment({ articleId: params.id as string });
  comment.value = resp.data;
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
