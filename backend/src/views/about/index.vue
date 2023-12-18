<template>
  <div>
    <el-row :gutter="10">
      <el-col :span="24" v-auth="btnAuth.about">
        <el-card>
          <template #header>
            <div class="card-header">
              <span>项目信息</span>
              <!-- changeVNode() -->
            </div>
          </template>
          <el-descriptions :column="4" v-bind:border="true" class="margin-top" direction="vertical" size="default">
            <el-descriptions-item label="项目名称">
              <el-tag type="success">{{ pkg.name }}</el-tag>
            </el-descriptions-item>
            <el-descriptions-item label="版本">
              <el-tag type="success">{{ pkg.version }}</el-tag>
            </el-descriptions-item>
            <el-descriptions-item label="vue版本">
              <el-tag type="success">{{ version }}</el-tag>
            </el-descriptions-item>
            <el-descriptions-item label="Github地址">
              <a href="https://github.com/JiangHaoCode/server-web" style="color: #606266" target="_blank">Github地址</a>
            </el-descriptions-item>
          </el-descriptions>
        </el-card>
      </el-col>
      <!-- <el-col :span="24">
        <ElCard class="box-card">
          <template #header>
            <div class="card-header">
              <span>测试</span>
              changeVNode() 
            </div>
          </template>
          <H :sum="88">
            <template #active>
              <div>123123</div>
            </template>
            <template v-slot:default="slotProps">
              <div>我是Slot 嘿嘿秘密 {{ slotProps.text }} {{ slotProps.count }}</div>
            </template>
            <template #[templateComponent]><div>dfasdfasdf</div></template>
          </H> 
        </ElCard>
      </el-col> -->
      <el-col :span="24">
        <Draggable />
      </el-col>
      <el-col :span="24" v-auth="btnAuth.about">
        <el-card class="box-card">
          <template #header>
            <div class="card-header">
              <span style="margin-right: 40px">提交记录 {{ t("menus.home") }}</span>
              <div :class="isShake ? 'swing' : ''" class="sncok" @click="changeShake()">
                <span class="iconfont icon-aixin1"></span>
              </div>
              <div :class="isShake ? 'animate__shakeX' : ''" class="animate__animated sncok" @click="changeShake()">
                <like fill="#f00" size="16" strokeLinejoin="miter" theme="filled" />
                <bank-card-one fill="#333" size="16" strokeLinejoin="miter" theme="outline" />
              </div>
              <!-- prop	description	type	default	note
theme	Theme of the icons.	'outline' | 'filled' | 'two-tone' | 'multi-color'	'outline'
size	The width/height of the icon	number | string	'1em'
spin	Rotate icon with animation	boolean	false
fill	Colors of theme	string | string[]	'currentColor'
strokeLinecap	the stroke-linecap prop of svg element	'butt' | 'round' | 'square'	'round'
strokeLinejoin	the stroke-linejoin prop of svg element	'miter' | 'round' | 'bevel'	'round'
strokeWidth	the stroke-width prop of svg element	number	4
-->
            </div>
          </template>
          <el-timeline>
            <el-timeline-item v-for="(activity, index) in commits" :key="index" :timestamp="activity.date" placement="top">
              <el-card>
                <h4>{{ activity.name }}</h4>
                <p>{{ activity.message }}</p>
              </el-card>
            </el-timeline-item>
          </el-timeline>
          <el-icon v-if="isLoading" class="is-loading">
            <Loading />
          </el-icon>
          <el-button v-show="!isShow" link size="default" type="primary" @click="commitHistory">Loading More</el-button>
        </el-card>
      </el-col>
    </el-row>
  </div>
</template>

<script>
export default {
  name: "About",
};
</script>

<script setup>
import Draggable from "./draggable.vue";
import pkg from "~/package.json";
import { getGithubCommitList } from "@/api/github";
import { onMounted, ref } from "vue";
import { Like, BankCardOne } from "@icon-park/vue-next";
import { useI18n } from "vue-i18n";
import { useBtnAuth } from "@/utils/btnAuth";
import { version } from "vue";
import { useUserStore } from "@/pinia/modules/user";

const btnAuth = useBtnAuth();

const { t } = useI18n();

const isShake = ref(false);

const commits = ref([]);

const isLoading = ref(false);
const isShow = ref(false);
const page = ref(1);

const userStore = useUserStore();

// const templateComponent = "active";

const changeShake = () => {
  isShake.value = true;
  setTimeout(() => {
    isShake.value = false;
  }, 1100);
};

onMounted(() => {
  if (btnAuth.about?.includes(userStore.userInfo.authorityId)) {
    getGithubCommitList({ page: page.value }).then((resp) => {
      commits.value = resp.data.list.map((i) => ({
        name: i.author,
        date: i.commit_time,
        message: i.message,
      }));
    });
  }

  // Members().then((resp) => {
  //   console.log(resp);
  // });
});

const commitHistory = () => {
  isLoading.value = true;
  page.value = page.value + 1;
  getGithubCommitList({ page: page.value }).then((resp) => {
    isLoading.value = false;
    if (resp.data.list?.length < 10) {
      isShow.value = true;
    }
    commits.value = [
      ...commits.value,
      ...resp.data.list.map((i) => ({
        name: i.author,
        date: i.commit_time,
        message: i.message,
      })),
    ];
  });
};

// import ButtonSlot from "./slot";
// import { reactive, ref } from "vue";
// import draggableVue from "./draggable.vue";
// import {
//   useMouse,
//   usePreferredDark,
//   useLocalStorage,
//   debounceFilter,
//   useDebounceFn,
//   throttleFilter,
//   pausableFilter,
//   useDeviceMotion,
// } from "@vueuse/core";
// import Button from "./Tabs/Button.vue";

// const values = ref("");

// const motionControl = pausableFilter();
// const motion = useDeviceMotion({ eventFilter: motionControl.eventFilter });
// motionControl.pause();
// motionControl.resume();

// const storage = useLocalStorage(
//   "keys",
//   { foo: "bar" },
//   { eventFilter: throttleFilter(100) }
// );
// console.log(storage);
// const { x, y } = useMouse({ eventFilter: debounceFilter(100) });

// // const { x, y } = useMouse();
// const isDark = usePreferredDark();
// const store = useLocalStorage("appColor", {
//   name: "Apple",
//   color: "red",
// });

// const data = reactive(useMouse());

// const timeDelay = useMouse({ eventFilter: debounceFilter(110) });
// const search = () => {
//   console.log("search");
//   inputValue.value = "mimi";
// };

// const searchValue = useDebounceFn(search, 500, { maxWait: 1000 });

// const lists = [
//   { name: "aaa", age: 12 },
//   { name: "bbb", age: 78 },
//   { name: "ccc", age: 45 },
// ];

// const inputValue = ref("");

// const changeInput = (e) => {
//   console.log(e);

//   inputValue.value = e;
// };
</script>

<style scoped>
.load-more {
  margin-left: 120px;
}

.avatar-img {
  float: left;
  height: 40px;
  width: 40px;
  border-radius: 50%;
  -webkit-border-radius: 50%;
  -moz-border-radius: 50%;
  margin-top: 15px;
}

.org-img {
  height: 150px;
  width: 150px;
}

.author-name {
  float: left;
  line-height: 65px !important;
  margin-left: 10px;
  color: darkblue;
  line-height: 100px;
  font-family: "Lucida Sans", "Lucida Sans Regular", "Lucida Grande", "Lucida Sans Unicode", Geneva, Verdana, sans-serif;
}

.dom-center {
  margin-left: 50%;
  transform: translateX(-50%);
}

h4 {
  margin: 1.3em 0;
}

.sncok {
  display: inline-block;
  width: 32px;
  height: 32px;
  text-align: center;
  transform-origin: 50% 0;
}
</style>
