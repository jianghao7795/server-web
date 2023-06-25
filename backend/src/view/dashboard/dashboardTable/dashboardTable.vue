<template>
  <div class="commit-table">
    <div class="commit-table-title">
      <el-button link @click="loadCommits" type="success" :loading="loadingPlue" :loading-icon="LoadingFour">更新日志</el-button>
      <el-button link @click="() => scrollChange(1)" type="success" :loading="loading" :loading-icon="LoadingFour">更新commit</el-button>
    </div>
    <div class="log">
      <div>
        <ul v-infinite-scroll="scrollChange" class="infinite-list log" :infinite-scroll-delay="500" :infinite-scroll-immediate="false" style="overflow: auto">
          <li v-for="(item, key) in dataTimeline" :key="item.commit_time" class="infinite-list-item log-item">
            <!-- {{ key + 1 }} - {{ item.message }} - {{ item.author }} -->

            <div class="flex-1 flex key-box">
              <span class="key" :class="key < 3 && 'top'">{{ key + 1 }}</span>
            </div>
            <div class="flex-7 flex message" :title="item.message">{{ item.message }}</div>
            <div class="flex-3 flex form" :title="item.from">{{ item.commit_time }}</div>
          </li>
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
// import { formatTimeToStr } from "@/utils/date.js";
import { getGithubCommitList, createCommit, Commits } from "@/api/github";
import { ref, onMounted } from "vue";
import { ElMessage, ElNotification } from "element-plus";
import { LoadingFour } from "@icon-park/vue-next";

const dataTimeline = ref([]);
const page = ref(1);
const loading = ref(false);
const loadingPlue = ref(false);

const scrollChange = (pageValue) => {
  if (pageValue) {
    page.value = pageValue;
    loading.value = true;
  } else {
    page.value = page.value + 1;
  }

  handleCreateGithub(!!pageValue);
};
const handleCreateGithub = (status = false) => {
  getGithubCommitList({ page: page.value, pageSize: 10 })
    .then((resp) => {
      if (resp?.code === 0) {
        if (status) {
          dataTimeline.value = resp.data.list || [];
        } else {
          dataTimeline.value = [...dataTimeline.value, ...resp.data.list];
        }
      }
    })
    .finally(() => {
      loading.value = false;
    });
};

const loadCommits = () => {
  loadingPlue.value = true;
  createCommit()
    .then((resp) => {
      if (resp.code === 0) {
        handleCreateGithub(true);
        ElMessage({
          type: "success",
          message: "更新 " + resp.data.total + " commit",
        });
      } else {
        ElMessage({
          type: "error",
          message: resp.msg,
        });
      }
    })
    .finally(() => {
      loadingPlue.value = false;
    })
    .catch((e) => {
      ElNotification({
        title: "请求错误：" + e.response.status,
        message: e.response.data.message || "",
        duration: 0,
      });
    })
    .finally(() => {
      loadingPlue.value = false;
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
  height: 20px;
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
  .commit-table-title {
    display: flex;
    justify-content: space-between;
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
        width: 25px;
        height: 25px;
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
        margin-left: 5px;
      }
      .flex {
        line-height: 14px;
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
      .flex-6 {
        flex: 6;
      }
      .flex-7 {
        flex: 13;
      }
    }
  }
}
</style>
