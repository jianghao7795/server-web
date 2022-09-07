import { createApp } from 'vue';
// import 'element-plus/dist/index.css';
import './style/element_visiable.scss';
// import ElementPlus from 'element-plus';
// import zhCn from 'element-plus/es/locale/lang/zh-cn';
import VueViewer from 'v-viewer';
// 引入gin-vue-admin前端初始化相关内容
import './core/gin-vue-admin';
// 引入封装的router
import router from '@/router/index';
import '@/permission';
import run from '@/core/gin-vue-admin.js';
import auth from '@/directive/auth';
import { store } from '@/pinia';
import App from './App.vue';
const app = createApp(App);
app.config.productionTip = false;

app.use(run).use(store).use(auth).use(router).use(VueViewer).mount('#app');

export default app;
