import { createRouter, createWebHistory } from "vue-router";
import HomeView from "@/views/home/index.vue";
import Layout from "@/views/layout/index.vue";
import About from "@/views/about/index.vue";
import Tag from "@/views/tag/index.vue";

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
        },
        {
          path: "/about",
          name: "about",
          component: About,
        },
        {
          path: "/tags",
          name: "tag",
          component: Tag,
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

export default router;
