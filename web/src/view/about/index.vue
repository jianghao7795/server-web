<template>
  <div>
    <el-row :gutter="10">
      <el-col :span="24">
        <el-card>
          <template #header>
            <div class="card-header">
              <span>项目信息</span>
            </div>
          </template>
          <el-descriptions direction="vertical" :column="3" size="default" border class="margin-top">
            <el-descriptions-item label="项目名称">
              <el-tag type="success">{{ pkg.name }}</el-tag>
            </el-descriptions-item>
            <el-descriptions-item label="版本">
              <el-tag type="success">{{ pkg.version }}</el-tag>
            </el-descriptions-item>
            <el-descriptions-item label="Github地址">
              <el-button link href="https://github.com/JiangHaoCode/server-web" target="_blank">Github地址</el-button>
            </el-descriptions-item>
          </el-descriptions>
        </el-card>
      </el-col>
      <el-col :span="24">
        <draggableVue />
      </el-col>
      <el-col :span="24">
        <el-card style="margin-top: 20px" class="box-card">
          <template #header>
            <div class="card-header">
              <span style="margin-right: 40px">提交记录</span>
              <div @click="changeShake()" :class="isShake ? 'swing' : ''" class="sncok">
                <span class="iconfont icon-aixin1"></span>
              </div>
              <div class="sncok animate__shakeX">
                <span class="iconfont icon-aixin1"></span>
              </div>
              <h1 class="animate__animated" :class="isShake ? 'animate__shakeX' : ''">An animated element</h1>
            </div>
          </template>
          <el-timeline>
            <el-timeline-item placement="top" v-for="(activity, index) in commits" :key="index" :timestamp="activity.date">
              <el-card>
                <h4>{{ activity.name }}</h4>
                <p>{{ activity.message }}</p>
              </el-card>
            </el-timeline-item>
          </el-timeline>
          <el-icon class="is-loading" v-if="isLoading"><Loading></Loading></el-icon>
          <el-button link type="primary" v-show="!isShow" size="default" @click="commitHistory">Loading More</el-button>
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
import draggableVue from "./draggable.vue";
import pkg from "~/package.json";
import { Commits } from "@/api/github";
import { onMounted, ref } from "vue";

const isShake = ref(false);

const commits = ref([]);

const isLoading = ref(false);
const isShow = ref(false);
const page = ref(1);

const changeShake = () => {
  isShake.value = true;
  setTimeout(() => {
    isShake.value = false;
  }, 1500);
};

onMounted(() => {
  Commits(page.value).then((resp) => {
    commits.value = resp.data.map((i) => ({ name: i.commit.author.name, date: i.commit.author.date, message: i.commit.message }));
  });
});

const commitHistory = () => {
  isLoading.value = true;
  page.value = page.value + 1;
  Commits(page.value).then((resp) => {
    isLoading.value = false;
    isShow.value = true;
    commits.value = [
      ...commits.value,
      ...resp.data.map((i) => ({ name: i.commit.author.name, date: i.commit.author.date, message: i.commit.message })),
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
.swing {
  animation: 1s ease-in-out 0s normal running upAnimation;
}

@keyframes upAnimation {
  /* 0% {
    transform: rotate(0deg);
    transition-timing-function: cubic-bezier(0.215, 0.61, 0.355, 1);
  }

  10% {
    transform: rotate(-12deg);
    transition-timing-function: cubic-bezier(0.215, 0.61, 0.355, 1);
  }

  20% {
    transform: rotate(12deg);
    transition-timing-function: cubic-bezier(0.215, 0.61, 0.355, 1);
  }

  28% {
    transform: rotate(-10deg);
    transition-timing-function: cubic-bezier(0.215, 0.61, 0.355, 1);
  }

  36% {
    transform: rotate(10deg);
    transition-timing-function: cubic-bezier(0.755, 0.5, 0.855, 0.06);
  }

  42% {
    transform: rotate(-8deg);
    transition-timing-function: cubic-bezier(0.755, 0.5, 0.855, 0.06);
  }

  48% {
    transform: rotate(8deg);
    transition-timing-function: cubic-bezier(0.755, 0.5, 0.855, 0.06);
  }

  52% {
    transform: rotate(-4deg);
    transition-timing-function: cubic-bezier(0.755, 0.5, 0.855, 0.06);
  }

  56% {
    transform: rotate(4deg);
    transition-timing-function: cubic-bezier(0.755, 0.5, 0.855, 0.06);
  }

  60% {
    transform: rotate(0deg);
    transition-timing-function: cubic-bezier(0.755, 0.5, 0.855, 0.06);
  }

  100% {
    transform: rotate(0deg);
    transition-timing-function: cubic-bezier(0.215, 0.61, 0.355, 1);
  } */
  0% {
    transform: translateX(0px);
  }
  33.3% {
    transform: translateX(10px);
  }
  66.7% {
    transform: translateX(-10px);
  }
  100% {
    transform: translateX(0px);
  }
}
</style>
