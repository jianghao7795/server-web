<template>
  <div class="page">
    <div class="card-box">
      <div class="card top-card">
        <div class="top-card-left">
          <div class="top-card-left-title">{{ period }}，{{ userStore.userInfo.nickName }}，请开始{{ timeNow }}的工作吧</div>
          <div class="top-card-left-dot">当前时间：{{ formatted }}</div>
          <div class="top-card-left-dot">{{ weatherInfo }}</div>
          <div class="top-card-left-rows">
            <el-row v-auth="888">
              <el-col :span="6" :xs="24" :sm="6">
                <div class="flex-center">
                  <el-icon class="dasboard-icon">
                    <sort />
                  </el-icon>
                  今日流量 {{ flowmeter }} M
                </div>
              </el-col>
              <el-col :span="6" :xs="24" :sm="6" v-auth="[888]">
                <div class="flex-center">
                  <el-icon class="dasboard-icon">
                    <avatar />
                  </el-icon>
                  总用户数 ({{ userNumber }})
                </div>
              </el-col>
              <el-col :span="6" :xs="24" :sm="6">
                <div class="flex-center">
                  <el-icon class="dasboard-icon">
                    <comment />
                  </el-icon>
                  好评率 (99%)
                </div>
              </el-col>
              <el-col :span="6" :xs="24" :sm="6">
                <div class="flex-center">
                  <el-icon class="dasboard-icon">
                    <View />
                  </el-icon>
                  阅读量 ({{ readingQuantity }})
                </div>
              </el-col>
            </el-row>
          </div>
        </div>
        <img src="@/assets/dashboard.png" class="top-card-right" alt />
      </div>
    </div>
    <div class="card-box" v-auth="[888, 998]">
      <el-card class="card quick-entrance">
        <template #header>
          <div class="card-header">
            <span>快捷入口</span>
          </div>
        </template>
        <el-row :gutter="20">
          <el-col v-for="(card, key) in toolCards" :key="key" :span="4" :xs="8" class="quick-entrance-items" @click="toTarget(card.name)">
            <div class="quick-entrance-item">
              <div class="quick-entrance-item-icon" :style="{ backgroundColor: card.bg }">
                <el-icon>
                  <component :is="card.icon" :style="{ color: card.color }" />
                </el-icon>
              </div>
              <p>{{ card.label }}</p>
            </div>
          </el-col>
        </el-row>
      </el-card>
      <!-- <div class="quick-entrance-title"></div> -->
    </div>
    <div class="card-box">
      <div class="card">
        <div class="card-header">
          <span>数据统计</span>
        </div>
        <div class="echart-box">
          <el-row :gutter="20">
            <el-col :xs="24" :sm="12">
              <!-- <swiper
                :modules="modules"
                :speed="1200"
                :loop="true"
                :slides-per-view="3"
                :space-between="10"
                navigation
                :pagination="{ clickable: true }"
                :scrollbar="{ draggable: true }"
                :autoplay="{ delay: 3000, disableOnInteraction: false }"
                effect="fade"
              >
                <swiper-slide><img src="https://static.runoob.com/images/demo/demo1.jpg" alt="" /></swiper-slide>
                <swiper-slide><img src="https://static.runoob.com/images/demo/demo2.jpg" alt="" /></swiper-slide>
                <swiper-slide><img src="https://static.runoob.com/images/demo/demo3.jpg" alt="" /></swiper-slide>
              </swiper> -->
              <div>
                <h3 class="py-24px text-24px font-bold">{{ swiperUploadFile.label }}</h3>
                <swiper v-bind="swiperUploadFile.options">
                  <swiper-slide v-for="i in uploadedFile" :key="i.id">
                    <div class="centerText">
                      <el-image class="centerFloat" :src="path + i.url" fit="contain"></el-image>
                    </div>
                  </swiper-slide>
                </swiper>
              </div>
              <!-- <div v-for="item in swiperExample" :key="item.id">
                <h3 class="py-24px text-24px font-bold">{{ item.label }}</h3>
                <swiper v-bind="item.options">
                  <swiper-slide v-for="i in 5" :key="i">
                    <div class="centerFloat">Slide{{ i }}</div>
                  </swiper-slide>
                </swiper>
              </div> -->
            </el-col>
            <el-col :xs="24" :sm="12">
              <dashboard-table />
            </el-col>
          </el-row>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup>
// import echartsLine from "@/view/dashboard/dashboardCharts/echartsLine.vue";
import dashboardTable from "@/views/dashboard/dashboardTable/dashboardTable.vue";
import { ref, onMounted } from "vue";
import { useRouter } from "vue-router";
import { useUserStore } from "@/pinia/modules/user";
import { useNow, useDateFormat } from "@vueuse/core";
import SwiperCore, { Navigation, Pagination, Autoplay } from "swiper";
import { Swiper, SwiperSlide } from "swiper/vue";
import { useWeatherInfo } from "./weather";
import { View, Comment, Sort, Avatar } from "@element-plus/icons-vue";
import "swiper/css";
import "swiper/css/pagination";
import "swiper/css/navigation";
import { getFlowmeter, userCount } from "@/api/user";
import { getReading } from "@/api/article";
import { getFileList } from "@/api/fileUploadAndDownload";

const path = import.meta.env.VITE_BASE_API + "/";

SwiperCore.use([Navigation, Pagination, Autoplay]);
// const swiperRef = useSwiper();

const weatherInfo = useWeatherInfo();
const userStore = useUserStore();
const formatted = useDateFormat(useNow(), "YYYY-MM-DD HH:mm:ss");
const period = ref("");
const userNumber = ref(0);
const flowmeter = ref(0);
const readingQuantity = ref(0);
const timeNow = ref("");

onMounted(() => {
  const nowHours = useNow().value.getHours();
  if (nowHours <= 6) {
    period.value = "凌晨好";
    timeNow.value = "一天";
  } else if (nowHours <= 12) {
    period.value = "上午好";
    timeNow.value = "一天";
  } else if (nowHours <= 18) {
    period.value = "下午好";
    timeNow.value = "下午";
  } else {
    period.value = "晚上好";
    timeNow.value = "晚上";
  }
  userCount().then((resp) => {
    if (resp?.code === 0) {
      userNumber.value = resp.data.count || 0;
    }
  });
  getFlowmeter().then((resp) => {
    flowmeter.value = ((resp.data.transmitBytes || 0) / 1024 / 1024).toFixed(0);
  });
  getReading().then((resp) => {
    readingQuantity.value = resp.data.reading_quantity || 0;
  });
  getFileList({ pageSize: 10, proportion: 1.5 }).then((resp) => {
    if (resp?.code === 0) {
      uploadedFile.value = resp.data.list;
    }
  });
});

const swiperUploadFile = ref({
  id: 8,
  label: "最新图片",
  options: {
    // navigation: {
    //   nextEl: ".swiper-button-next",
    //   prevEl: ".swiper-button-prev",
    // },
    pagination: {
      clickable: true,
    },
    loop: true,
    autoplay: {
      delay: 3000,
      stopOnLastSlide: false,
      disableOnInteraction: false, // 设置为false和自动播放将在用户交互后不会被禁用（Swipes），互动后每次都会重新启动它
      pauseOnMouseEnter: true, // 当启用时，将在鼠标上通过Swiper容器进入鼠标时暂停。如果还启用了disableOninteraction，它将停止自动播放而不是暂停
    },
  },
});

const uploadedFile = ref([]);

// const swiperExample = ref([
//   {
//     id: 0,
//     label: "Default",
//     options: {
//       autoplay: {
//         // 自动播放
//         delay: 3000,
//         stopOnLastSlide: false,
//         disableOnInteraction: false, // 设置为false和自动播放将在用户交互后不会被禁用（Swipes），互动后每次都会重新启动它
//         pauseOnMouseEnter: true, // 当启用时，将在鼠标上通过Swiper容器进入鼠标时暂停。如果还启用了disableOninteraction，它将停止自动播放而不是暂停
//       },
//     },
//   },
//   {
//     id: 1,
//     label: "Navigation",
//     autoplay: true,
//     options: {
//       loop: true,
//       navigation: true,
//       autoplay: {
//         delay: 3000,
//         stopOnLastSlide: false,
//         disableOnInteraction: true,
//       },
//     },
//   },
//   {
//     id: 2,
//     label: "Pagination",
//     options: {
//       loop: true,
//       pagination: {
//         clickable: true,
//       },
//     },
//   },
//   {
//     id: 3,
//     label: "Pagination dynamic",
//     options: {
//       loop: true,
//       pagination: { dynamicBullets: true },
//     },
//   },
//   {
//     id: 4,
//     label: "Pagination progress",
//     options: {
//       navigation: true,
//       loop: true,
//       pagination: {
//         type: "progressbar",
//       },
//     },
//   },
//   {
//     id: 5,
//     label: "Pagination fraction",
//     options: {
//       navigation: true,
//       loop: true,
//       pagination: {
//         type: "fraction",
//       },
//     },
//   },
//   {
//     id: 6,
//     label: "Slides per view",
//     options: {
//       pagination: {
//         clickable: true,
//       },
//       slidesPerView: 3,
//       spaceBetween: 30,
//       loop: true,
//     },
//   },
//   {
//     id: 7,
//     label: "Infinite loop",
//     options: {
//       navigation: true,
//       pagination: {
//         clickable: true,
//       },
//       loop: true,
//       autoplay: {
//         delay: 3000,
//         stopOnLastSlide: false,
//         disableOnInteraction: false, // 设置为false和自动播放将在用户交互后不会被禁用（Swipes），互动后每次都会重新启动它
//         pauseOnMouseEnter: true, // 当启用时，将在鼠标上通过Swiper容器进入鼠标时暂停。如果还启用了disableOninteraction，它将停止自动播放而不是暂停
//       },
//     },
//   },
// ]);

const toolCards = ref([
  {
    label: "用户管理",
    icon: "monitor",
    name: "user",
    color: "#ff9c6e",
    bg: "rgba(255, 156, 110,.3)",
  },
  {
    label: "角色管理",
    icon: "setting",
    name: "authority",
    color: "#69c0ff",
    bg: "rgba(105, 192, 255,.3)",
  },
  {
    label: "菜单管理",
    icon: "menu",
    name: "menu",
    color: "#b37feb",
    bg: "rgba(179, 127, 235,.3)",
  },
  {
    label: "代码生成器",
    icon: "cpu",
    name: "autoCode",
    color: "#ffd666",
    bg: "rgba(255, 214, 102,.3)",
  },
  {
    label: "表单生成器",
    icon: "document-checked",
    name: "formCreate",
    color: "#ff85c0",
    bg: "rgba(255, 133, 192,.3)",
  },
  {
    label: "关于我们",
    icon: "user",
    name: "about",
    color: "#5cdbd3",
    bg: "rgba(92, 219, 211,.3)",
  },
]);

const router = useRouter();

const toTarget = (name) => {
  // console.log(router);
  router.push({ name });
};
</script>
<script>
export default {
  name: "Dashboard",
};
</script>

<style lang="scss" scoped>
@mixin flex-center {
  display: flex;
  align-items: center;
}

.centerFloat {
  height: 300px;
  border: #eee solid 1px;
}

.centerText {
  text-align: center;
}

.page {
  background: #f0f2f5;
  padding: 0;

  .card-box {
    padding: 12px 16px;

    & + .card-box {
      padding-top: 0px;
    }
  }

  .card {
    box-sizing: border-box;
    background-color: #fff;
    border-radius: 2px;
    height: auto;
    padding: 26px 30px;
    overflow: hidden;
    box-shadow: 0 0 7px 1px rgba(0, 0, 0, 0.03);
  }

  .top-card {
    height: 260px;
    @include flex-center;
    justify-content: space-between;
    color: #777;

    &-left {
      height: 100%;
      display: flex;
      flex-direction: column;

      &-title {
        font-size: 22px;
        color: #343844;
      }

      &-dot {
        font-size: 14px;
        color: #6b7687;
        margin-top: 24px;
      }

      &-rows {
        // margin-top: 15px;
        margin-top: 18px;
        color: #6b7687;
        width: 600px;
        align-items: center;
      }

      &-item {
        + .top-card-left-item {
          margin-top: 24px;
        }

        margin-top: 14px;
      }
    }

    &-right {
      height: 600px;
      width: 600px;
      margin-top: 28px;
    }
  }

  ::v-deep(.el-card__header) {
    padding: 0;
    border-bottom: none;
  }

  .card-header {
    padding-bottom: 20px;
    border-bottom: 1px solid #e8e8e8;
  }

  .quick-entrance-title {
    height: 30px;
    font-size: 22px;
    color: #333;
    width: 100%;
    border-bottom: 1px solid #eee;
  }

  .quick-entrance-items {
    @include flex-center;
    justify-content: center;
    text-align: center;
    color: #333;

    .quick-entrance-item {
      padding: 16px 28px;
      margin-top: -16px;
      margin-bottom: -16px;
      border-radius: 4px;
      transition: all 0.2s;

      &:hover {
        box-shadow: 0px 0px 7px 0px rgba(217, 217, 217, 0.55);
      }

      cursor: pointer;
      height: auto;
      text-align: center;

      // align-items: center;
      &-icon {
        width: 50px;
        height: 50px !important;
        border-radius: 8px;
        @include flex-center;
        justify-content: center;
        margin: 0 auto;

        i {
          font-size: 24px;
        }
      }

      p {
        margin-top: 10px;
      }
    }
  }

  .echart-box {
    padding: 14px;
  }
}

.dasboard-icon {
  font-size: 20px;
  color: rgb(85, 160, 248);
  width: 30px;
  height: 30px;
  margin-right: 10px;
  @include flex-center;
}

.flex-center {
  @include flex-center;
}

//小屏幕不显示右侧，将登陆框居中
@media (max-width: 750px) {
  .card {
    padding: 20px 10px !important;

    .top-card {
      height: auto;

      &-left {
        &-title {
          font-size: 20px !important;
        }

        &-rows {
          margin-top: 15px;
          align-items: center;
        }
      }

      &-right {
        display: none;
      }
    }

    .middle-card {
      &-item {
        line-height: 20px;
      }
    }

    .dasboard-icon {
      font-size: 18px;
    }
  }
}
</style>
