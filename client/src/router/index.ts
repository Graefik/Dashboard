import { createRouter, createWebHistory } from "vue-router";

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes: [
    {
      path: "/",
      name: "home",
      component: () => import("@/features/dashboard/Dashboad.view.vue"),
    },
    {
      path: "/login",
      name: "login",
      component: () => import("@/features/login/Login.view.vue"),
    },
  ],
});

export default router;
