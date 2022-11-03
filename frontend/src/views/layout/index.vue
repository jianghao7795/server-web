<template>
  <div>
    <n-layout position="absolute">
      <n-layout-header position="static">
        <n-card :bordered="false" header-style="headerStyle" class="darkStyle">
          <template #header-extra>
            <div class="headerStyleLine">
              <NSpace>
                <b>首页</b>
                <b>标签</b>
                <b>公众号</b>
                <b>关于</b>
              </NSpace>
            </div>
          </template>
          <template #header>
            <div class="headerStyleLine"><b>瓦都尅</b></div>
          </template>
          <div style="height: 300px; text-align: center">
            <h1>瓦都尅</h1>
            <hr class="small" />
            <span class="subheading">愈有知，愈无知。</span>
          </div>
        </n-card>
      </n-layout-header>
      <n-layout-content position="static">
        <RouterView />
      </n-layout-content>
      <n-layout-footer position="absolute">
        <!-- style="width: 100%; height: 100px; position: absolute; bottom: 0px; left: 0px" -->
        <footer class="footerStyle">
          <span>Copyright ©2022</span>
        </footer>
      </n-layout-footer>
    </n-layout>
  </div>
</template>

<script setup lang="ts">
import { RouterView } from "vue-router";
import {
  NCard,
  NSpace,
  NLayout,
  NLayoutContent,
  NLayoutFooter,
  NLayoutHeader,
} from "naive-ui";
</script>

<style scoped lang="less">
@import "@/static/less/content.less";
.headerStyle {
  padding: 15px 15px;
  height: 60px;
  line-height: 60px;
}

.headerStyleLine {
  color: #fff;
  font-size: 15px;
}

.layoutHight {
  height: calc(100vh - 2000px);
}

.darkStyle {
  background: url("/home-bg.png") no-repeat center center;
  height: auto;
  width: 100%;
  background-size: cover;
  background-attachment: scroll;
  color: #fff;
  z-index: 1;
}

.footerStyle {
  text-align: center;
  height: 40px;
  flex: 0 0 auto;
}

hr.small {
  max-width: 150px;
  margin: 15px auto;
  border-width: 4px;
  border-color: white;
}
.small {
  font-size: 85%;
}
hr {
  border: 0;
  border-top: 1px solid #eee;
}

.subheading {
  font-size: 24px;
  line-height: 1.1;
  display: block;
  font-weight: 300;
  margin: 10px 0 0;
}
// .admin-layout__footer {
//   left: 0;
//   bottom: 0;
//   flex-shrink: 0;
//   box-sizing: border-box;
//   width: 100%;
//   transition-property: padding-left;
//   flex: 0 0 auto;
// }
</style>
