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
                <n-switch
                  v-model:value="darkTheme"
                  v-bind:on-update:value="changeTheme"
                  size="medium"
                  :rail-style="railStyle"
                >
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
            <div class="headerStyleLine">
              <n-space>
                <b @click="() => changePath('/')" v-if="isLogin" style="cursor: pointer">
                  <n-dropdown :options="options" placement="bottom-start" trigger="click" @select="userLogout">
                    {{ userStore.currentUser.user.name }}
                  </n-dropdown>
                </b>
                <b @click="() => changeLogin(true)" v-else style="cursor: pointer">登录</b>
                <span style="cursor: pointer" @click="() => changeActive(true)" v-if="isLogin">更换背景图片</span>
              </n-space>
            </div>
          </template>
          <div style="height: 300px; text-align: center">
            <div v-if="!isMouseOver">
              <n-popover trigger="hover">
                <template #trigger>
                  <Search size="24" theme="outline" @click="() => changeBlur(true)" fill="#333" :strokeWidth="3" />
                </template>
                <span>搜索文章</span>
              </n-popover>
            </div>
            <NInput
              v-else
              :autofocus="true"
              ref="searchInputRef"
              v-model:value="searchInput"
              style="max-width: 30%; margin: 0"
              placeholder="搜索文章"
              type="text"
              @blur="() => changeBlur(false)"
              @keyup.enter="submit"
            />
            <span class="subheading">愈有知，愈无知。</span>
          </div>
        </n-card>
      </n-layout-header>
      <!-- include exclude 包含和不包含 字符串是组件的name -->
      <n-layout-content position="static" class="middle-view">
        <NSpin :show="loadingFlag">
          <RouterView v-slot="{ Component, route }">
            <transition mode="out-in" name="el-fade-in-linear" type="transition" :appear="true">
              <keep-alive v-bind:exclude="['ArticleDetail', 'Article']">
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
    <n-drawer v-model:show="active" placement="bottom" :height="400">
      <n-drawer-content title="更换背景图片">
        <n-carousel :space-between="30" :loop="false" slides-per-view="auto" centered-slides draggable>
          <n-carousel-item style="width: 30%" v-for="item in bgImage" :key="item.ID">
            <n-popconfirm positive-text="确认" negative-text="取消" :on-positive-click="() => changeImages(item)">
              <template #trigger>
                <img
                  :src="item.url.includes('http') ? item.url : `/${item.url}`"
                  :title="item.name"
                  class="carousel-img"
                />
              </template>
              确定更换背景图片
            </n-popconfirm>
          </n-carousel-item>
        </n-carousel>
      </n-drawer-content>
    </n-drawer>
    <n-drawer v-model:show="loginStatus" :width="502" placement="left">
      <n-drawer-content title="登录">
        <n-form
          ref="formRef"
          :model="userInfo"
          :rules="rules"
          label-placement="left"
          label-width="auto"
          require-mark-placement="right-hanging"
          size="large"
        >
          <n-form-item path="name">
            <n-input type="text" v-model:value="userInfo.name" placeholder="账号" />
          </n-form-item>
          <n-form-item path="password">
            <n-input type="password" show-password-on="click" v-model:value="userInfo.password" placeholder="密码" />
          </n-form-item>
          <div>
            <n-button :loading="userStore.loading" type="primary" :block="true" @click="() => login()">登录</n-button>
          </div>
        </n-form>
      </n-drawer-content>
    </n-drawer>
  </div>
</template>

<script lang="ts">
export default {
  name: "Layout",
};
</script>

<script setup lang="ts">
import { KeepAlive, Transition, onMounted, ref, type CSSProperties, watch, inject, type Ref, computed, h } from "vue";
import type { GlobalTheme, FormInst, DropdownOption } from "naive-ui";
import { NIcon } from "naive-ui";
import { RouterView, useRouter, useRoute } from "vue-router";
import { Search, Sun, SunOne, Logout } from "@icon-park/vue-next";
import dayjs from "dayjs";
import { emitter } from "@/utils/common";
import { getImages } from "@/services/image";
import { useUserStore } from "@/stores/user";
import { updateBackgroundImage } from "@/services/user";

const userStore = useUserStore();
const route = useRoute();
const router = useRouter();
const searchInputRef = ref<HTMLInputElement>();
const searchInput = ref<string>("");
const loadingFlag = ref<boolean>(false);
const isMouseOver = ref<boolean>(false);
const viewPage = ref<string>(route.fullPath);
const colorSet = ref<string>(`url(${new URL("/home-bg.png", import.meta.url).href})`);
const bgImage = ref<User.Images[]>([]);
const active = ref<boolean>(false);
//登录页面的status
const loginStatus = ref<boolean>(false);
// 主题状态
const theme = inject<Ref<GlobalTheme | null>>("theme");
const darkTheme = computed(() => !(theme?.value === null));
//form Ref
const formRef = ref<FormInst | null>(null);

// 是否登录
const isLogin = computed(() => !!userStore.currentUser.user.ID);
const userInfo = ref<User.UserInfo>({
  ID: 0,
  name: "",
  introduction: "",
  head_img: "",
  content: "",
  password: "",
});

const rules = {
  name: {
    required: true,
    trigger: ["blur", "input"],
    message: "请输入账号",
  },
  password: {
    required: true,
    trigger: ["blur", "input"],
    message: "请输入密码",
  },
};
//
const options = ref<DropdownOption[]>([
  {
    label: "退出登录",
    key: "logout",
    icon: () => {
      console.log(darkTheme.value);
      return h(NIcon, null, {
        default: () =>
          h(Logout, {
            theme: "outline",
            size: "26",
            fill: darkTheme.value ? "#ddd" : "#333",
            strokeWidth: 3,
          }),
      });
    },
  },
]);

const userLogout = (key: string | number) => {
  if (key === "logout") {
    userStore.$reset();
    localStorage.removeItem("token");
    colorSet.value = `url(${new URL("/home-bg.png", import.meta.url).href})`;
  }
};
//<logout theme="outline" size="26" fill="#ddd" :strokeWidth="3"/>
//<logout theme="outline" size="26" fill="#333" :strokeWidth="3"/>
const changeBlur = (status: boolean) => {
  isMouseOver.value = status;
  if (status) {
    setTimeout(() => {
      searchInputRef.value?.focus();
    });
  }
};

const changeImages = async (data: User.Images) => {
  await updateBackgroundImage({ ID: userStore.currentUser.user.ID, head_img: data.url });
  window.$message.success("更换成功");
  active.value = false;
  colorSet.value = `url(${new URL(data.url.includes("http") ? data.url : `/${data.url}`, import.meta.url).href})`;
};

const changeActive = (status: boolean) => {
  active.value = status;
};

const changeLogin = (status: boolean) => {
  loginStatus.value = status;
};

const login = () => {
  userStore.logins({ name: userInfo.value.name, password: userInfo.value.password }, (imageString: string) => {
    if (!!imageString) {
      colorSet.value = `url(${
        new URL(imageString.includes("http") ? imageString : `/${imageString}`, import.meta.url).href
      })`;
    }
    loginStatus.value = false;
    userInfo.value = {
      ID: 0,
      name: "",
      introduction: "",
      head_img: "",
      content: "",
      password: "",
    };
  });
};

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

  getImages().then((resp) => {
    if (resp) {
      bgImage.value = resp.data;
    }
  });
  userStore.getUser((head_img: string) => {
    colorSet.value = `url(${new URL(head_img.includes("http") ? head_img : `/${head_img}`, import.meta.url).href})`;
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
.carousel-img {
  width: 100%;
  // height: 240px;
  object-fit: cover;
}
.headerStyleLine {
  font-size: 15px;
}

.darkStyle {
  // background: url("/home-bg.png") no-repeat center center;
  background-image: v-bind(colorSet);
  background-position: center;
  height: auto;
  width: 100%;
  background-size: cover;
  background-attachment: scroll;
  z-index: 1;
  margin-bottom: 20px;
  // color: #fff;
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
