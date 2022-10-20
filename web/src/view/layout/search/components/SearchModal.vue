<template>
  <div>
    <el-dialog :model-value="props.show" :before-close="changeShow" :draggable="true" :destroy-on-close="true" :width="500">
      <el-input ref="searchInput" v-model="serarchValue" placeholder="请输入" @input="changeSearchRouter"></el-input>
      <el-empty description="暂无搜索结果" v-show="searchRouter.length === 0" />
      <el-scrollbar v-if="searchRouter.length !== 0">
        <template v-for="(i, index) in searchRouter" :key="i.label">
          <div
            class="border-every padding-pad"
            :class="mouseoverdown[index] ? 'padding-pad-over' : ''"
            @mouseover="() => (mouseoverdown[index] = true)"
            @mouseleave="() => (mouseoverdown[index] = false)"
            @click="linkRouter(i.value)"
          >
            <div class="after-search">
              <div class="margin-icon">
                <el-space wrap :size="5">
                  <el-icon v-if="routeMap[i.value]?.meta?.icon"><component :is="routeMap[i.value]?.meta?.icon" /></el-icon>
                  <span>{{ i.label }}</span>
                </el-space>
              </div>
              <div><connection theme="outline" size="12" /></div>
            </div>
          </div>
        </template>
      </el-scrollbar>
    </el-dialog>
  </div>
</template>

<script>
export default {
  name: "SearchModal",
};
</script>
<script setup>
import { ref } from "vue";
import { Connection } from "@icon-park/vue-next";
import { useRouterStore } from "@/pinia/modules/router";
import { useDebounceFn } from "@vueuse/core";
import { useRouter } from "vue-router";
const router = useRouter();

const props = defineProps({
  show: {
    type: Boolean,
    default: false,
  },
  changeShowSearchStatus: {
    type: Function,
  },
});

// const emit = defineEmits(["changeShowSearchStatus"]);

const searchInput = ref(null);
const serarchValue = ref("");
const mouseoverdown = ref({});

const changeShow = () => {
  //   console.log(emit);
  props.changeShowSearchStatus(false);
  serarchValue.value = "";
};

const routerStore = useRouterStore();
const routeMap = routerStore.routeMap;
const searchRouter = ref([]);
const changeInputSearch = (e) => {
  if (e) {
    const searchRouterInfo = routerStore.routerList.filter((i) => {
      //   console.log(i.value);
      //   console.log(routeMap[i.value].meta.icon);
      if (i.label.includes(e)) {
        return true;
      }
      return false;
    });

    searchRouter.value = searchRouterInfo;
  } else {
    searchRouter.value = [];
  }
};

const changeSearchRouter = useDebounceFn(changeInputSearch, 1000);
// console.log(routerStore.asyncRouters);
// console.log(routerStore.routeMap["dashboard"].meta.icon);
// console.log(props.show);

const linkRouter = (name = "") => {
  if (name !== "") {
    router.push({ name });
    props.changeShowSearchStatus(false);
    searchRouter.value = [];
    serarchValue.value = "";
  }
};
</script>

<style>
.el-dialog .el-dialog__header {
  padding: 0;
  border-bottom: 0;
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
  justify-content: center;
  height: 50px;
  background: var(--el-color-primary-light-9);
  margin: 10px;
  color: var(--el-color-primary);
}
.infinite-list .infinite-list-item + .list-item {
  margin-top: 10px;
}

.after-search {
  display: flex;
  justify-content: space-between;
  align-items: center;
  height: 30px;
}

.padding-pad {
  background-color: #eee;
  color: #111;
}
.padding-pad-over {
  background-color: #177cb0;
  color: #fff;
  cursor: pointer;
}

.border-every {
  border: 1px solid #eee;
  margin: 5px 0;
  padding: 5px 10px;
}

.margin-icon {
  line-height: 32px;
  height: 32px;
  display: flex;
  justify-content: space-between;
  align-items: center;
}
</style>
