import { createRouter, createWebHistory } from "vue-router";

const routes = [
  {
    path: "/",
    redirect: "/login",
  },
  {
    path: "/init",
    name: "Init",
    component: () => import("@/view/init/index.vue"),
  },
  {
    path: "/login",
    name: "Login",
    component: () => import("@/view/login/index.vue"),
  },
];

const router = createRouter({
  history: createWebHistory(),
  routes,
  linkActiveClass: "active",
});

export default router;
