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
                <div style="margin-top: 2px">
                  <n-switch
                    v-model:value="darkTheme"
                    v-bind:on-update:value="changeTheme"
                    size="medium"
                    :rail-style="railStyle"
                  >
                    <template #checked-icon>
                      <NIcon style="line-height: 0.7rem">
                        <moon
                          theme="filled"
                          size="26"
                          fill="#333"
                          :strokeWidth="3"
                        />
                      </NIcon>
                    </template>
                    <template #unchecked-icon>
                      <NIcon style="line-height: 0.7rem">
                        <sun-one
                          theme="outline"
                          size="26"
                          fill="#333"
                          :strokeWidth="3"
                        />
                      </NIcon>
                    </template>
                  </n-switch>
                </div>
              </NSpace>
            </div>
          </template>
          <template #header>
            <span class="headerStyleLine">
              <b
                @click="() => changePath('/')"
                v-if="isLogin"
                style="cursor: pointer"
              >
                <n-dropdown
                  :options="options"
                  placement="bottom-end"
                  trigger="click"
                  :show-arrow="true"
                  @select="userLogout"
                >
                  <n-avatar round size="small" :src="headImage"></n-avatar>
                </n-dropdown>
              </b>
              <span v-else>
                <b @click="() => changeLogin(true)" style="cursor: pointer"
                  >登录</b
                >
                <NDivider vertical />
                <b
                  @click="() => changeRegisterStatus(true)"
                  style="cursor: pointer"
                  >注册</b
                >
              </span>
            </span>
          </template>
          <div style="height: 300px; text-align: center">
            <div v-if="!isMouseOver">
              <n-popover trigger="hover">
                <template #trigger>
                  <Search
                    size="24"
                    theme="outline"
                    @click="() => changeBlur(true)"
                    :strokeWidth="3"
                  />
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
            <span class="subheading">{{ currentRouter }}</span>
          </div>
        </n-card>
      </n-layout-header>
      <!-- include exclude 包含和不包含 字符串是组件的name -->
      <n-layout-content position="static" class="middle-view">
        <NSpin :show="loadingFlag">
          <RouterView v-slot="{ Component, route }">
            <transition
              mode="out-in"
              name="fade-in-linear"
              type="transition"
              :appear="true"
            >
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
      <n-back-top :right="50" />
    </n-layout>
    <n-drawer v-model:show="active" placement="bottom" :height="400">
      <n-drawer-content title="更换背景图片">
        <n-carousel
          :space-between="30"
          :loop="false"
          slides-per-view="auto"
          draggable
        >
          <n-carousel-item
            style="width: 30%"
            v-for="item in bgImage"
            :key="item.ID"
          >
            <n-popconfirm
              positive-text="确认"
              negative-text="取消"
              :on-positive-click="() => changeImages(item)"
            >
              <template #trigger>
                <img
                  :src="
                    item.url.includes('http')
                      ? item.url
                      : `${Base_URL}/${item.url}`
                  "
                  :title="item.name"
                  class="carousel-img"
                />
              </template>
              确定更换背景图片？
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
          @keyup.enter.native="login"
        >
          <n-form-item path="name">
            <n-input
              type="text"
              v-model:value="userInfo.name"
              placeholder="账号"
            />
          </n-form-item>
          <n-form-item path="password">
            <n-input
              type="password"
              show-password-on="click"
              v-model:value="userInfo.password"
              placeholder="密码"
            />
          </n-form-item>
          <div>
            <n-button
              :loading="userStore.loading"
              type="primary"
              :block="true"
              @click="() => login()"
              >登录</n-button
            >
          </div>
        </n-form>
      </n-drawer-content>
    </n-drawer>
    <register
      :register-status="registerStatus"
      @change-status="changeRegisterStatus"
    />
    <ResetPassord
      :active="revisePassword"
      @change-status="changeResetPasswordStatus"
      @resetStore="resetStore"
    />
    <Person
      :active="personalInformationStatus"
      @changeStatus="changePersonalInformationStatus"
    />
  </div>
</template>

<script lang="ts">
export default {
  name: "Layout",
};
</script>

<script setup lang="ts">
import {
  KeepAlive,
  Transition,
  onMounted,
  ref,
  watch,
  inject,
  provide,
  computed,
  h,
} from "vue";
import type { CSSProperties, Ref } from "vue";
import type { GlobalTheme, FormInst } from "naive-ui";
import { NIcon } from "naive-ui";
import { RouterView, useRouter, useRoute } from "vue-router";
import {
  Search,
  Logout,
  Change,
  Moon,
  SunOne,
  SettingTwo,
  Lock,
} from "@icon-park/vue-next";
import dayjs from "dayjs";
import { emitter } from "@/utils/common";
import { getImages } from "@/services/image";
import { useUserStore } from "@/stores/user";
import { updateBackgroundImage } from "@/services/user";
import Register from "./components/register.vue";
import Person from "./components/person.vue";
import ResetPassord from "./components/reset_password.vue";

const Base_URL = import.meta.env.VITE_BASE_API;

const headImage = computed(
  () => `${Base_URL}/${userStore.currentUser.user.headerImg}`
);

const userStore = useUserStore();
const route = useRoute();
const router = useRouter();
const searchInputRef = ref<HTMLInputElement>();
const searchInput = ref<string>("");
const loadingFlag = ref<boolean>(false);
const isMouseOver = ref<boolean>(false);
const viewPage = ref<string>(route.fullPath);
const colorSet = ref<string>(
  `url(${new URL("/home-bg.png", import.meta.url).href})`
);
const bgImage = ref<User.Images[]>([]);
const active = ref<boolean>(false);
const currentRouter = ref<string>("");
// 注册页面的status
const registerStatus = ref<boolean>(false);
//登录页面的status
const loginStatus = ref<boolean>(false);

const personalInformationStatus = ref<boolean>(false);
// 主题状态
const theme = inject<Ref<GlobalTheme | null>>("theme");
const darkTheme = computed(() => !(theme?.value === null));
//form Ref
const formRef = ref<FormInst | null>(null);
// 修改密码status
const revisePassword = ref<boolean>(false);

// 是否登录
const isLogin = computed(() => !!userStore.currentUser.user.ID);
const userInfo = ref<{ name: string; password: string }>({
  name: "",
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
const options = [
  {
    label: "个人信息",
    key: "setting",
    icon: () => {
      return h(NIcon, null, {
        default: () =>
          h(SettingTwo, {
            theme: "outline",
            size: 26,
            strokeWidth: 3,
          }),
      });
    },
  },
  //<setting-two theme="outline" size="26" fill="#333" :strokeWidth="3"/>
  {
    label: "修改密码",
    key: "lock",
    icon: () => {
      return h(NIcon, null, {
        default: () => {
          return h(Lock, {
            theme: "outline",
            size: 26,
            strokeWidth: 3,
          });
        },
      });
    },
  },
  // <lock theme="outline" size="26" fill="#333" :strokeWidth="3"/>
  {
    label: "更改背景图",
    key: "change",
    icon: () => {
      return h(NIcon, null, {
        default: () =>
          h(Change, {
            theme: "outline",
            size: "26",
            strokeWidth: 3,
          }),
      });
    },
  },
  {
    label: "退出登录",
    key: "logout",
    icon: () => {
      return h(NIcon, null, {
        default: () =>
          h(Logout, {
            theme: "outline",
            size: "26",
            strokeWidth: 3,
          }),
      });
    },
  },
];

const resetStore = () => {
  userStore.$reset();
  localStorage.removeItem("token");
  colorSet.value = `url(${new URL("/home-bg.png", import.meta.url).href})`;
  loginStatus.value = true;
};

const changePersonalInformationStatus = (status: boolean) => {
  personalInformationStatus.value = status;
};

const userLogout = (key: string | number) => {
  if (key === "logout") {
    resetStore();
  }
  // debugger;
  if (key === "change") {
    changeActive(true);
  }

  if (key === "lock") {
    changeResetPasswordStatus(true);
  }

  if (key === "setting") {
    changePersonalInformationStatus(true);
  }
};

const changeRegisterStatus = (status: boolean): void => {
  registerStatus.value = status;
};

const changeResetPasswordStatus = (status: boolean): void => {
  revisePassword.value = status;
};

const changeBlur = (status: boolean) => {
  isMouseOver.value = status;
  if (status) {
    setTimeout(() => {
      searchInputRef.value?.focus();
    });
  }
};

const changeImages = async (data: User.Images) => {
  await updateBackgroundImage({
    ID: userStore.currentUser.user.ID,
    head_img: data.url,
  });
  window.$message.success("更换成功");
  active.value = false;
  colorSet.value = `url(${
    new URL(
      data.url.includes("http") ? data.url : `${Base_URL}/${data.url}`,
      import.meta.url
    ).href
  })`;
};

const changeActive = (status: boolean) => {
  active.value = status;
};

const changeLogin = (status: boolean): void => {
  loginStatus.value = status;
};
provide("changeLogin", changeLogin); // 传递方法给下级

const login = () => {
  userStore.logins(
    { username: userInfo.value.name, password: userInfo.value.password },
    (imageString: string) => {
      if (!!imageString) {
        colorSet.value = `url(${
          new URL(
            imageString.includes("http")
              ? imageString
              : `${Base_URL}/${imageString}`,
            import.meta.url
          ).href
        })`;
      }
      getImages().then((resp) => {
        if (resp) {
          bgImage.value = resp.data;
        }
      });
      loginStatus.value = false;
      userInfo.value = {
        name: "",
        password: "",
      };
    }
  );
};

const railStyle = ({ checked }: { checked: boolean }) => {
  const style: CSSProperties = { lineHeight: "1rem" };
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
    const pathArray = value.split("/");
    viewPage.value = `/${pathArray[1]}`;
    switch (pathArray[1]) {
      case undefined:
        currentRouter.value = "首页";
        break;
      case "articles":
        currentRouter.value = "文章";
        break;
      case "about":
        currentRouter.value = "关于我";
        break;
      case "tags":
        currentRouter.value = "标签";
        break;
      default:
        currentRouter.value = "首页";
    }
  }
);

onMounted(() => {
  emitter.on("showLoading", () => {
    loadingFlag.value = true;
  });
  emitter.on("closeLoading", () => {
    loadingFlag.value = false;
  });

  const pathArray = route.fullPath.split("/");
  viewPage.value = `/${pathArray[1]}`;
  switch (pathArray[1]) {
    case undefined:
      currentRouter.value = "首页";
      break;
    case "articles":
      currentRouter.value = "文章";
      break;
    case "about":
      currentRouter.value = "关于我";
      break;
    case "tags":
      currentRouter.value = "标签";
      break;
    default:
      currentRouter.value = "首页";
  }
  const token = localStorage.getItem("token");
  if (token) {
    userStore.getUser((head_img: string) => {
      getImages().then((resp) => {
        if (resp?.code === 0) {
          bgImage.value = resp.data;
        }
      });
      if (head_img !== "") {
        colorSet.value = `url(${
          new URL(
            head_img.includes("http") ? head_img : `${Base_URL}/${head_img}`,
            import.meta.url
          ).href
        })`;
      }
    });
  }
});

const changePath = (url: string) => {
  router.push(url);
};

const submit = () => {
  if (searchInput.value === "") {
    window.$message.warning("请输入");
    return;
  }
  router.push(`/articles/search/${searchInput.value}`);
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
  min-height: calc(100% - 460px);
}
</style>
