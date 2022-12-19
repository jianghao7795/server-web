<template>
  <div class="commit-table">
    <div class="commit-table-title">更新日志</div>
    <div class="log">
      <el-scrollbar height="400px" @scroll="scrollChange">
        <div v-for="(item, key) in dataTimeline" :key="item.message" class="log-item">
          <div class="flex-1 flex key-box">
            <span class="key" :class="key < 3 && 'top'">{{ key + 1 }}</span>
          </div>
          <div class="flex-5 flex message" :title="item.message">{{ item.message }}</div>
          <div class="flex-3 flex form" :title="item.from">{{ item.from }}</div>
        </div>
      </el-scrollbar>
    </div>
  </div>
</template>

<script>
export default {
  name: "DashboardTable",
};
</script>
<script setup>
import { Commits } from "@/api/github";
import { formatTimeToStr } from "@/utils/date.js";
import { ref, onMounted } from "vue";

const dataTimeline = ref([]);
const page = ref(1);

const scrollChange = (event) => {
  console.log(event);
  if (event.scrollTop === 1620) {
    page.value = page.value + 1;
    loadCommits();
  }
};

const loadCommits = () => {
  Commits(page.value).then(({ data }) => {
    // console.log(data);
    const line = [];
    data.forEach((element) => {
      if (element.commit.message) {
        line.push({
          from: formatTimeToStr(element.commit.author.date, "yyyy-MM-dd hh:mm:ss"),
          title: element.commit.author.name,
          showDayAndMonth: true,
          message: element.commit.message,
        });
      }
    });
    dataTimeline.value = line;
    // console.log(line);
  });
};

onMounted(() => {
  loadCommits();
});
</script>

<style lang="scss" scoped>
.loadingCenter {
  text-align: center;
  margin-bottom: 15px;
}
.commit-table {
  background-color: #fff;
  height: 400px;
  &-title {
    font-weight: 600;
    margin-bottom: 12px;
  }
  .log {
    &-item {
      display: flex;
      justify-content: space-between;
      margin-top: 14px;
      .key-box {
        justify-content: center;
      }
      .key {
        &.top {
          background: #314659;
          color: #ffffff;
        }
        display: inline-flex;
        justify-content: center;
        align-items: center;
        width: 20px;
        height: 20px;
        border-radius: 50%;
        background: #f0f2f5;
        text-align: center;
        color: rgba($color: #000000, $alpha: 0.65);
      }
      .message {
        color: rgba(0, 0, 0, 0.65);
      }
      .form {
        color: rgba(0, 0, 0, 0.65);
        margin-left: 12px;
      }
      .flex {
        line-height: 20px;
        overflow: hidden;
        text-overflow: ellipsis;
        white-space: nowrap;
      }
      .flex-1 {
        flex: 1;
      }
      .flex-2 {
        flex: 2;
      }
      .flex-3 {
        flex: 3;
      }
      .flex-4 {
        flex: 4;
      }
      .flex-5 {
        flex: 5;
      }
    }
  }
}
</style>
