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
        <el-card style="margin-top: 20px" class="box-card">
          <template #header>
            <div class="card-header">
              <span>提交记录</span>
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
import pkg from "~/package.json";
import { Commits } from "@/api/github";
import { onMounted, ref } from "vue";

const commits = ref([]);

onMounted(() => {
  Commits(0).then((resp) => {
    commits.value = resp.data.map((i) => ({ name: i.commit.author.name, date: i.commit.author.date, message: i.commit.message }));
  });
});

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
  margin: 20px 0;
}
</style>
