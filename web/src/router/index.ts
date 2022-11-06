import { createRouter, createWebHistory } from "vue-router";
import type { RouteRecordRaw } from "vue-router";

// ¾²Ì¬Â·ÓÉ
const routes = [
  {
    path: "/",
    redirect: "/login",
  },
  {
    path: "/login",
    name: "Login",
    component: () => import("@/view/login/login.vue"),
  },
  // {
  //   path: "/layout",
  //   name: "Layout",
  //   component: () => import("@/view/layout/layout.vue"),
  //   children: [
  //     {
  //       path: "dashboard",
  //       component: () => import("@/view/dashboard/dashboard.vue"),
  //     },
  //   ],
  // },
];

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes,
});

export default router;
