<template>
  <div>
    <n-layout position="absolute">
      <n-layout-header position="static">
        <n-card :bordered="false" header-style="headerStyle" class="darkStyle">
          <template #header-extra>
            <div class="headerStyleLine">
              <NSpace>
                <a @click="changePath('/')">
                  <span :class="underLineLable('/')"><b>首页</b></span>
                </a>
                <a @click="changePath('/articles')">
                  <span :class="underLineLable('/articles')"><b>文章</b></span>
                </a>
                <a @click="changePath('/tags')">
                  <span :class="underLineLable('/tags')"><b>标签</b></span>
                </a>
                <a @click="changePath('/about')">
                  <span :class="underLineLable('/about')"><b>关于</b></span>
                </a>
              </NSpace>
            </div>
          </template>
          <template #header>
            <div class="headerStyleLine"><b @click="changePath('/')">吴昊</b></div>
          </template>
          <div style="height: 300px; text-align: center">
            <span class="small-h1" v-if="!isMouseOver || route.path.includes('/search/article')" @mouseover="mouseOverTitle(true)"><b>吴昊</b></span>
            <NInput
              ref="searchInputRef"
              v-model:value="searchInput"
              @blur="mouseOverTitle(false)"
              style="max-width: 30%, margin: '15px 0'"
              placeholder="搜索文章"
              type="text"
              @keyup.enter="submit"
              v-else
            />
            <hr class="small" />
            <span class="subheading">愈有知，愈无知。</span>
          </div>
        </n-card>
      </n-layout-header>
      <n-layout-content position="static" class="middle-view">
        <RouterView />
      </n-layout-content>
      <n-layout-footer position="static">
        <!-- style="width: 100%; height: 100px; position: absolute; bottom: 0px; left: 0px" -->
        <footer class="footerStyle">
          <span>Copyright ©{{ dayjs().format("YYYY") }}</span>
        </footer>
      </n-layout-footer>
    </n-layout>
  </div>
</template>

<script lang="ts">
export default {
  name: "Layout",
};
</script>

<script setup lang="ts">
import { RouterView, useRouter, useRoute } from "vue-router";
import { NCard, NSpace, NLayout, NLayoutContent, NLayoutFooter, NLayoutHeader, NInput } from "naive-ui";
import dayjs from "dayjs";
import { ref, nextTick } from "vue";
const route = useRoute();
const router = useRouter();
const searchInput = ref("");
const searchInputRef = ref<HTMLInputElement>();

const isMouseOver = ref(false);
const underLineLable = (url: string): string => {
  if (route.path === url) {
    return "underLine";
  } else {
    return "";
  }
};
const changePath = (url: string) => {
  router.push(url);
};

const mouseOverTitle = async (isStatus: boolean) => {
  isMouseOver.value = isStatus;
  if (!isStatus) {
    await nextTick();
    searchInputRef.value?.focus();
  }
};

const submit = () => {
  if (searchInput.value === "") {
    window.$message.warning("请输入");
    return;
  }
  router.push(`/search/article/${searchInput.value}`);
  isMouseOver.value = false;
  searchInput.value = "";
};
</script>

<style scoped lang="less">
.headerStyleLine {
  color: #fff;
  font-size: 15px;
}

.darkStyle {
  background: url("/home-bg.png") no-repeat center center;
  height: auto;
  width: 100%;
  background-size: cover;
  background-attachment: scroll;
  color: #fff;
  z-index: 1;
  margin-bottom: 20px;
}

.footerStyle {
  text-align: center;
  height: 40px;
  flex: 0 0 auto;
  line-height: 40px;
}

hr.small {
  max-width: 150px;
  margin: 15px auto;
  border-width: 4px;
  border-color: white;
}
.small {
  font-size: 85%;
}

.small-h1 {
  font-size: 21px;
}
hr {
  border: 0;
  border-top: 1px solid #eee;
}

.subheading {
  font-size: 24px;
  line-height: 1.1;
  display: block;
  font-weight: 300;
  margin: 10px 0 0;
}
a::before {
  cursor: pointer;
}

.middle-view {
  min-height: calc(100% - 470px);
}
</style>
