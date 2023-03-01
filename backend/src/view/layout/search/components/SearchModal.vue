<template>
  <div>
    <el-dialog :model-value="show" :before-close="() => (show = false)" :draggable="true" :width="500">
      <el-form>
        <el-form-item>
          <el-input
            :prefix-icon="Search"
            ref="searchInputRef"
            v-model="serarchValue"
            placeholder="请输入关键词搜索"
            clearable
            @input="changeSearchRouter"
          ></el-input>
        </el-form-item>
      </el-form>
      <template #header>
        <div>搜索</div>
      </template>
      <el-empty description="暂无搜索结果" v-if="searchRouter.length === 0" />
      <el-scrollbar v-else>
        <template v-for="(i, index) in searchRouter" :key="i.label">
          <div
            class="border-every padding-pad"
            :class="!!mouseoverdown[index] ? 'padding-pad-over' : ''"
            @mouseover="() => (mouseoverdown = { [index]: true })"
            @click="linkRouter(i.value)"
            @enter="handleEnter"
          >
            <!-- @mouseleave="() => (mouseoverdown[index] = false)" -->
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
      <template #footer>
        <div>
          <el-space wrap>
            <div class="after-search">
              <enter-key-one theme="outline" size="24" fill="#333" />
              确认
            </div>
            <span class="after-search">
              <arrow-circle-up theme="outline" size="24" fill="#333" />
              <arrow-circle-down theme="outline" size="24" fill="#333" />
              切换
            </span>
          </el-space>
        </div>
      </template>
    </el-dialog>
  </div>
</template>

<script>
export default {
  name: "SearchModal",
};
</script>
<script setup>
import { ref, computed, watch } from "vue";
import { Connection, EnterKeyOne, ArrowCircleUp, ArrowCircleDown } from "@icon-park/vue-next";
import { useRouterStore } from "@/pinia/modules/router";
import { useDebounceFn, onKeyStroke } from "@vueuse/core";
import { useRouter } from "vue-router";
import { Search } from "@element-plus/icons-vue";
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

const searchInputRef = ref();
const serarchValue = ref("");
const mouseoverdown = ref({});

const show = computed({
  get() {
    return props.show;
  },
  set(val) {
    props.changeShowSearchStatus(val);
    if (!val) {
      serarchValue.value = "";
    }
  },
});

watch(show, (val) => {
  if (val) {
    /** 自动聚焦 */
    // await nextTick(() => searchInputRef.value.focus()); //不行 无法渲染成功
    setTimeout(() => searchInputRef.value.focus()); // 需要延迟才能正常渲染
  }
});

// onMounted(() => {
//   searchInputRef.value?.focus();
// });

// const changeShow = () => {
//   //   console.log(emit);
//   props.changeShowSearchStatus(false);
//   serarchValue.value = "";
// };

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
    if (searchRouterInfo.length !== 0) {
      mouseoverdown.value = { 0: true };
    }
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

const handleEnter = (e) => {
  e.preventDefault();
  if (props.show) {
    const name = searchRouter.value[Object.keys(mouseoverdown.value)[0]].value;
    router.push({ name });
    props.changeShowSearchStatus(false);
    searchRouter.value = [];
    serarchValue.value = "";
  }
};

/* key up */
const handleUp = (e) => {
  e.preventDefault();
  // console.log("up", e);
  if (searchRouter.value.length <= 1) {
    return;
  }
  const lengthSearchRouter = searchRouter.value.length;
  let mouseoverdownPoints = { ...mouseoverdown.value };
  searchRouter.value.forEach((_, index) => {
    if (mouseoverdown.value[index]) {
      if (index === 0) {
        mouseoverdownPoints = { [lengthSearchRouter - 1]: true };
      } else {
        mouseoverdownPoints = { [index - 1]: true };
      }
    }
  });

  mouseoverdown.value = mouseoverdownPoints;
};

/* key down */
const handleDown = (e) => {
  e.preventDefault();
  // console.log("down", e);
  if (searchRouter.value.length <= 1) {
    return;
  }
  const lengthSearchRouter = searchRouter.value.length;
  let mouseoverdownPoints = { ...mouseoverdown.value };
  searchRouter.value.forEach((_, index) => {
    if (mouseoverdown.value[index]) {
      if (index === lengthSearchRouter - 1) {
        mouseoverdownPoints = { 0: true };
      } else {
        mouseoverdownPoints = { [index + 1]: true };
      }
    }
  });
  mouseoverdown.value = mouseoverdownPoints;
};

// onKeyStroke('Escape', handleClose);
onKeyStroke("Enter", handleEnter);
onKeyStroke("ArrowUp", handleUp);
onKeyStroke("ArrowDown", handleDown);
</script>

<style lang="scss" scoped>
.header .el-dialog__header {
  border-bottom: 0;
  .el-dialog__header {
    padding: 0;
    border-bottom: 0;
  }
  .el-dialog .el-dialog__header {
    padding: 2px 20px 12px 20px;
    border-bottom: 0;
  }
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
