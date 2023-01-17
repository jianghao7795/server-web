<template>
  <div class="view-content">
    <div class="view-margin">
      <h1 class="view-center">{{ detail?.title }}</h1>
      <h3>简介：{{ detail?.title }}</h3>
      <h4>
        <NSpace style="width: 80%">
          标签：
          <n-tag size="small" round v-for="(item, index) in detail?.tags" :type="colorIndex(index)">{{ item.name }}</n-tag>
        </NSpace>
      </h4>
      <MdEditor style="width: 1000px" :model-value="detail?.content" :pageFullscreen="true" :previewOnly="true"></MdEditor>
    </div>
  </div>
</template>

<script lang="ts" setup>
import { NSpace, NTag } from "naive-ui";
import { ref, onMounted } from "vue";
import { getArticleDetail } from "@/services/article";
import { useRoute } from "vue-router";
import type { API } from "@/type/article";
import { colorIndex } from "@/common/article";
import MdEditor from "md-editor-v3";
import "md-editor-v3/lib/style.css";

const detail = ref<API.Article>();
const route = useRoute();
// const router = useRouter();

onMounted(async () => {
  const params = route.params;
  console.log(route.params, params);
  const resp = await getArticleDetail(params.id as string);
  detail.value = resp.data?.article as API.Article;
  //   console.log(resp);
});
</script>
