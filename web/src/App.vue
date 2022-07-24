<template>
  <div id="app">
    <router-view />
  </div>
</template>

<script>
export default {
  name: "App",
  data() {
    // console.log(this.$router);
    return {};
  },
};
</script>

<style lang="scss">
// 引入初始化样式
@import "@/style/main.scss";
@import "@/style/base.scss";
@import "@/style/mobile.scss";
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
  -webkit-text-decoration-skip: objects;
}
</style>
