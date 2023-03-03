<template>
  <div>
    <n-layout position="absolute">
      <n-layout-header position="static">
        <n-card :bordered="false" header-style="headerStyle" class="darkStyle">
          <template #header-extra>
            <div class="headerStyleLine">
              <NSpace>
                <n-tabs
                  type="bar"
                  animated
                  :value="viewPage"
                  size="small"
                  :bar-width="28"
                  justify-content="space-evenly"
                  :tab-style="{ margin: '0 5px', fontWeight: 'bold' }"
                  :on-update:value="(e: string) => changePath(e)"
                >
                  <n-tab-pane name="/" tab="首页"></n-tab-pane>
                  <n-tab-pane name="/articles" tab="文章"></n-tab-pane>
                  <n-tab-pane name="/tags" tab="标签"></n-tab-pane>
                  <n-tab-pane name="/about" tab="关于"></n-tab-pane>
                </n-tabs>
                <n-switch v-bind:on-update:value="changeTheme" size="medium" :rail-style="railStyle">
                  <template #checked-icon>
                    <NIcon :component="Sun" />
                  </template>
                  <template #unchecked-icon>
                    <NIcon :component="SunOne" />
                  </template>
                </n-switch>
              </NSpace>
            </div>
          </template>
          <template #header>
            <div class="headerStyleLine"><b @click="changePath('/')">吴昊</b></div>
          </template>
          <div style="height: 300px; text-align: center">
            <span v-if="!isMouseOver">
              <span class="small-h1">
                <b>吴昊</b>
              </span>
              <Search
                @click="
                  () => {
                    isMouseOver = true;
                  }
                "
              />
            </span>
            <NInput
              v-else
              :autofocus="true"
              ref="searchInputRef"
              v-model:value="searchInput"
              style="max-width: 30%; margin: '15px 0'"
              placeholder="搜索文章"
              type="text"
              @keyup.enter="submit"
            />
            <hr class="small" />
            <span class="subheading">愈有知，愈无知。</span>
          </div>
        </n-card>
      </n-layout-header>
      <n-layout-content position="static" class="middle-view">
        <NSpin :show="loadingFlag">
          <RouterView v-slot="{ Component, route }">
            <transition mode="out-in" name="el-fade-in-linear" type="transition" :appear="true">
              <keep-alive>
                <component :is="Component" :key="route.name" />
              </keep-alive>
            </transition>
          </RouterView>
        </NSpin>
      </n-layout-content>
      <n-layout-footer position="static">
        <!-- style="width: 100%; height: 100px; position: absolute; bottom: 0px; left: 0px" -->
        <footer class="footerStyle">
          <span>Copyright © {{ dayjs().format("YYYY") }}</span>
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
import { KeepAlive, Transition, onMounted, ref, type CSSProperties, watch } from "vue";
import { RouterView, useRouter, useRoute } from "vue-router";
import { NCard, NSpace, NLayout, NLayoutContent, NLayoutFooter, NLayoutHeader, NInput, NSpin, NIcon, NSwitch, NTabs, NTabPane } from "naive-ui";
import { Search, Sun, SunOne } from "@icon-park/vue-next";
import dayjs from "dayjs";
import { emitter } from "@/utils/common";

const route = useRoute();
const router = useRouter();
const searchInput = ref<string>("");
const searchInputRef = ref<HTMLInputElement>();

const loadingFlag = ref<boolean>(false);

const isMouseOver = ref<boolean>(false);
const viewPage = ref<string>(route.fullPath);

const railStyle = ({ checked }: { checked: boolean }) => {
  const style: CSSProperties = {};
  if (checked) {
    style.background = "#222";
  } else {
    style.background = "#eee";
  }
  return style;
};

const changeTheme = (e: boolean) => {
  if (e) {
    emitter.emit("darkMode");
  } else {
    emitter.emit("lightMode");
  }
};

watch(
  () => route.fullPath,
  (value) => {
    viewPage.value = value;
  },
);

onMounted(() => {
  emitter.on("showLoading", () => {
    loadingFlag.value = true;
  });
  emitter.on("closeLoading", () => {
    loadingFlag.value = false;
  });
});

const changePath = (url: string) => {
  router.push(url);
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
  font-size: 15px;
}

.darkStyle {
  background: url("/home-bg.png") no-repeat center center;
  height: auto;
  width: 100%;
  background-size: cover;
  background-attachment: scroll;
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
}
.small {
  font-size: 85%;
}

.small-h1 {
  font-size: 21px;
  margin-right: 5px;
}
hr {
  border: 0;
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
  min-height: calc(100% - 443px);
}
</style>
