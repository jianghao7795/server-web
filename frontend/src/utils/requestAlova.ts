import { createAlova } from "alova";
import GlobalFetch from "alova/GlobalFetch";
import VueHook from "alova/vue";
const alovaInstance = createAlova({
  // 服务器根目录
  baseURL: "/api",

  // VueHook用于创建ref状态，包括请求状态loading、响应数据data、请求错误对象error等（后续详细介绍）
  statesHook: VueHook,

  // 请求适配器，我们推荐并提供了fetch请求适配器
  requestAdapter: GlobalFetch(),
});
