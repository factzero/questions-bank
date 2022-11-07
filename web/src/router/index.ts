import { createRouter, createWebHistory } from "vue-router";
import type { RouteRecordRaw } from "vue-router";

// 静态路由
const routes = [
  {
    path: "/login",
    name: "Login",
    component: () => import("@/view/login/login.vue"),
  },
  {
    path: "/",
    name: "Layout",
    component: () => import("@/view/layout/layout.vue"),
    redirect: "/layout/dashboard",
    children: [
      {
        path: "dashboard",
        component: () => import("@/view/dashboard/dashboard.vue"),
      },
    ],
  },
];

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes,
});

export default router;
