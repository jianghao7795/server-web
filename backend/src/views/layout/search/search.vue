<template>
  <div class="search-component">
    <!-- <transition name="el-fade-in-linear">
      <div v-show="show" class="transition-box" style="display: inline-block; margin-top: 0px">
        <el-select ref="searchInput" v-model="value" filterable placeholder="请选择" @blur="hiddenSearch" @change="changeRouter">
          <el-option v-for="item in routerStore.routerList" :key="item.value" :label="item.label" :value="item.value" />
        </el-select>
      </div>
    </transition> -->
    <div class="user-box">
      <div class="Icon Icon-refresh" :class="[reload ? 'reloading' : '']" @click="handleReload"></div>
    </div>
    <div class="user-box">
      <div class="Icon Icon-search" @click="showSearch(true)"></div>
    </div>
    <div class="user-box">
      <Screenfull class="search-icon" :style="{ cursor: 'pointer' }" />
    </div>
    <!-- <div v-if="btnShow" class="user-box">
      <div class="Icon Icon-customer-service" @click="toService"></div>
    </div> -->
    <SearchModal :show="show" v-bind:changeShowSearchStatus="showSearch" />
  </div>
</template>

<script>
export default {
  name: "BtnBox",
};
</script>

<script setup>
import Screenfull from "@/views/layout/screenfull/index.vue";
import { emitter } from "@/utils/bus.js";
import { ref } from "vue";
// import { useRouterStore } from "@/pinia/modules/router";
import { ElLoading } from "element-plus";
import SearchModal from "@/views/layout/search/components/SearchModal.vue";

// const router = useRouter();

// const routerStore = useRouterStore();

const value = ref("");
// const changeRouter = () => {
//   router.push({ name: value.value });
//   value.value = "";
// };

const show = ref(false);
// const btnShow = ref(true);
// const hiddenSearch = () => {
//   show.value = false;
//   setTimeout(() => {
//     btnShow.value = true;
//   }, 500);
// };

// const searchInput = ref(null);
const showSearch = (statusValue = true) => {
  // console.log(statusValue);
  show.value = statusValue;
  // btnShow.value = false;
  // show.value = true;
  // await nextTick();
  // searchInput.value.focus();
};

const reload = ref(false);
const handleReload = () => {
  reload.value = true;
  const loadingInstance = ElLoading.service({ fullscreen: false, target: "#refreshView", text: "加载中" });
  emitter.emit("reload");
  setTimeout(() => {
    reload.value = false;
    loadingInstance.close();
  }, 500);
};
// const refresh = () => {
//   setTimeout(() => {
//     refreshStore.changIsRefresh(false);
//     setTimeout(() => {
//       refreshStore.changIsRefresh(true);
//     });
//     setTimeout(() => {
//       loadingInstance.close();
//       NProgress.done();
//     }, 100);
//   }, 1000);
// };
// const toService = () => {
//   window.open("https://support.qq.com/product/371961");
// };
</script>
<style scoped lang="scss">
.reload {
  font-size: 18px;
}

.reloading {
  animation: turn 0.5s linear infinite;
}
@keyframes turn {
  0% {
    -webkit-transform: rotate(0deg);
  }
  25% {
    -webkit-transform: rotate(90deg);
  }
  50% {
    -webkit-transform: rotate(180deg);
  }
  75% {
    -webkit-transform: rotate(270deg);
  }
  100% {
    -webkit-transform: rotate(360deg);
  }
}
</style>
