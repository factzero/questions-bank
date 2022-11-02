import { createRouter, createWebHistory } from "vue-router";

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes: [
    {
      path: "/",
      redirect: "/login",
    },
    {
      path: "/login",
      name: "login",
      component: () => import("@/views/login/login.vue"),
    },
    {
      path: "/layout",
      name: "layout",
      component: () => import("@/views/layout/layout.vue"),
    },
  ],
});

export default router;
