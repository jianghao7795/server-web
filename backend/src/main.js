/*
 * @Author: jianghao
 * @Date: 2022-09-05 07:56:08
 * @LastEditors: jianghao
 * @LastEditTime: 2022-10-18 10:57:23
 */
import { createApp } from "vue";
import "@/style/element_visiable.scss";
// import ElementPlus from "element-plus";
// import zhCn from "element-plus/es/locale/lang/zh-cn";
// import "./core/vue-admin";
// 引入封装的router
import router from "@/router/index";
import "@/permission";
import run from "@/core/vue-admin";
import auth from "@/directive/auth";
import { store } from "@/pinia";
import App from "./App.vue";
import "vue-cropper/dist/index.css";
// import "virtual:windi.css"; // 会搞乱的
// import "animate.css"; //导入animate.css包 https://animate.style/
import i18n from "@/locals/index"; //语言翻译 i18b

const app = createApp(App);
app.config.productionTip = false;

app.use(run).use(store).use(auth).use(router).use(i18n).mount("#app");

export default app;
