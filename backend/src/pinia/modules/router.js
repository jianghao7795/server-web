import { asyncRouterHandle } from "@/utils/asyncRouter";

import { asyncMenu } from "@/api/menu";
import { defineStore } from "pinia";
import { ref } from "vue";

const routerListArr = [];
const keepAliveRoutersArr = [];

const formatRouter = (routes, routeMap) => {
  routes &&
    routes.forEach((item) => {
      if ((!item.children || item.children.every((ch) => ch.hidden)) && item.name !== "404" && !item.hidden) {
        routerListArr.push({ label: item.meta.title, value: item.name });
      }
      item.meta.btns = item.btns;
      item.meta.hidden = item.hidden;
      routeMap[item.name] = item;
      if (item.children && item.children.length > 0) {
        formatRouter(item.children, routeMap);
      }
    });
};

const KeepAliveFilter = (routes) => {
  routes &&
    routes.forEach((item) => {
      // 子菜单中有 keep-alive 的，父菜单也必须 keep-alive，否则无效。这里将子菜单中有 keep-alive 的父菜单也加入。
      if ((item.children && item.children.some((ch) => ch.meta.keepAlive)) || item.meta.keepAlive) {
        item.component &&
          item.component().then((val) => {
            keepAliveRoutersArr.push(val.default.name);
          });
      }
      if (item.children && item.children.length > 0) {
        KeepAliveFilter(item.children);
      }
    });
};

export const useRouterStore = defineStore("router", () => {
  const asyncRouters = ref([]);
  const routerList = ref(routerListArr);
  const keepAliveRouters = ref(keepAliveRoutersArr);
  const routeMap = {};
  // 从后台获取动态路由
  const SetAsyncRouter = async () => {
    const baseRouter = [
      {
        path: "/layout",
        name: "layout",
        component: "views/layout/index.vue",
        meta: {
          title: "layout",
        },
        children: [],
      },
    ];
    const asyncRouterRes = await asyncMenu();
    if (asyncRouterRes.code === 0) {
      const asyncRouter = asyncRouterRes.data.menus;
      asyncRouter &&
        asyncRouter.push(
          {
            path: "404",
            name: "404",
            hidden: true,
            meta: {
              title: "404",
              closeTab: true,
            },
            component: "views/error/index.vue",
          },
          {
            path: "reload",
            name: "Reload",
            hidden: true,
            meta: {
              title: "",
              closeTab: true,
            },
            component: "views/error/reload.vue",
          },
        );
      formatRouter(asyncRouter, routeMap);
      baseRouter[0].children = asyncRouter;
      baseRouter.push({
        path: "/:catchAll(.*)",
        redirect: "/layout/404",
      });
      asyncRouterHandle(baseRouter);
      KeepAliveFilter(asyncRouter);
      asyncRouters.value = baseRouter;
      routerList.value = routerListArr;
      keepAliveRouters.value = keepAliveRoutersArr;
      return true;
    } else {
      asyncRouters.value = baseRouter;
      return false;
    }
  };

  return {
    asyncRouters,
    routerList,
    keepAliveRouters,
    SetAsyncRouter,
    routeMap,
  };
});
