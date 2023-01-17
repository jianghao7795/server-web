import { createRouter, createWebHistory } from "vue-router";
import HomeView from "@/views/home/index.vue";
import Layout from "@/views/layout/index.vue";
import About from "@/views/about/index.vue";
import Tag from "@/views/tag/index.vue";
import Article from "@/views/article/index.vue";
import ArticleDetail from "@/views/article/detail.vue";

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes: [
    {
      path: "/",
      name: "home",
      component: Layout,
      children: [
        {
          path: "/",
          name: "index",
          component: HomeView,
          meta: {
            title: "首页",
          },
        },
        {
          path: "/about",
          name: "about",
          component: About,
          meta: {
            title: "关于",
          },
        },
        {
          path: "/tags",
          name: "tag",
          component: Tag,
          meta: {
            title: "标签",
          },
        },
        {
          path: "/articles",
          name: "article",
          component: Article,
          meta: {
            title: "文章",
          },
        },
        {
          path: "/articles/:id",
          name: "id",
          component: ArticleDetail,
          meta: {
            title: "文章详情",
          },
        },
      ],
    },
    // {
    //   path: '/about',
    //   name: 'about',
    //   // route level code-splitting
    //   // this generates a separate chunk (About.[hash].js) for this route
    //   // which is lazy-loaded when the route is visited.
    //   component: () => import('../views/AboutView.vue'),
    // },
  ],
});

router.beforeEach((to, from, next) => {
  window.document.title = to.meta.title as string;
  window.$loadingBar.start();
  next();
});

router.afterEach(() => {
  window.$loadingBar.finish();
});

router.onError(() => {
  window.$loadingBar.error();
});

export default router;
