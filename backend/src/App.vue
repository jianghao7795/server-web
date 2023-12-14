<template>
  <div id="app">
    <!-- <router-view /> -->
    <el-config-provider :locale="zhCn">
      <router-view />
    </el-config-provider>
  </div>
</template>

<script>
export default {
  name: "App",
};
</script>

<script setup>
import zhCn from "element-plus/dist/locale/zh-cn.mjs";
</script>

<style lang="scss">
// 引入初始化样式
@import "@/style/main.scss";
@import "@/style/base.scss";
@import "@/style/mobile.scss";
@import "@/style/newLogin.scss";
@import "@/style/element_visiable.scss";
#app {
  background: #eee;
  height: 100vh;
  overflow: hidden;
  font-weight: 400 !important;
}

.el-button {
  font-weight: 400 !important;
}

a {
  color: #1890ff;
  text-decoration: none;
  background-color: #0000;
  outline: none;
  cursor: pointer;
  transition: color 0.3s;
}

// 加载条的 css
#nprogress .bar {
  background-color: #409eff;
  height: 4px;
}
</style>
