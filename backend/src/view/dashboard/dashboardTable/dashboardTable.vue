<template>
  <div class="commit-table">
    <div class="commit-table-title">
      <p @click="loadCommits">
        更新日志
        <el-icon v-loading="loadingPlue"><loading theme="outline" size="24" fill="#333" /></el-icon>
      </p>
    </div>
    <div class="log">
      <!-- <el-scrollbar @scroll="scrollChange" height="400px">
        <div v-for="(item, key) in dataTimeline" :key="item.message" class="log-item">
          <div class="flex-1 flex key-box">
            <span class="key" :class="key < 3 && 'top'">{{ key + 1 }}</span>
          </div>
          <div class="flex-5 flex message" :title="item.message">{{ item.message }}</div>
          <div class="flex-3 flex form" :title="item.from">{{ item.from }}</div>
        </div>
      </el-scrollbar> -->
      <div>
        <ul v-infinite-scroll="scrollChange" class="infinite-list" :infinite-scroll-delay="500" :infinite-scroll-immediate="false" style="overflow: auto">
          <li v-for="(item, key) in dataTimeline" :key="item.commit_time" class="infinite-list-item">{{ key + 1 }} - {{ item.message }} - {{ item.author }}</li>
        </ul>
      </div>
    </div>
  </div>
</template>

<script>
export default {
  name: "DashboardTable",
};
</script>
<script setup>
// import { Commits } from "@/api/github";
import { formatTimeToStr } from "@/utils/date.js";
import { getGithubCommitList, createCommit, Commits } from "@/api/github";
import { ref, onMounted } from "vue";
import { Loading } from "@icon-park/vue-next";
import { ElMessage } from "element-plus";

const dataTimeline = ref([]);
const page = ref(1);
// const loading = ref(false);
const loadingPlue = ref(false);

const scrollChange = () => {
  page.value = page.value + 1;
  handleCreateGithub();
};
const handleCreateGithub = () => {
  // loading.value = true;
  getGithubCommitList({ page: page.value, pageSize: 10 }).then((resp) => {
    // console.log(resp);
    dataTimeline.value = [...dataTimeline.value, ...resp.data.list];
  });
};

const loadCommits = () => {
  // loadingPlue.value = true;
  Commits({ page: page.value }).then(({ data }) => {
    const line = data.map((element) => {
      return {
        commit_time: formatTimeToStr(element.commit.author.date, "yyyy-MM-dd hh:mm:ss"),
        author: element.commit.author.name,
        message: element.commit.message,
      };
    });
    createCommit(line).then((resp) => {
      if (resp.code === 0) {
        ElMessage({
          type: "success",
          message: "更新成功!",
        });
      }
    });
  });
};

onMounted(() => {
  handleCreateGithub();
});
</script>

<style lang="scss" scoped>
.loadingCenter {
  text-align: center;
  margin-bottom: 15px;
}
.infinite-list {
  height: 300px;
  padding: 0;
  margin: 0;
  list-style: none;
}
.infinite-list .infinite-list-item {
  display: flex;
  align-items: center;
  justify-content: flex-start;
  height: 40px;
  background: #fff;
  margin: 2px;
  color: #111;
}
.infinite-list .infinite-list-item + .list-item {
  margin-top: 10px;
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
