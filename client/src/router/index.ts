import { createRouter, createWebHistory } from "vue-router";
import { useAuth } from "@/shared/auth/useAuth";

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes: [
    {
      path: "/",
      name: "home",
      component: () => import("@/features/dashboard/Dashboard.view.vue"),
      meta: { requiresAuth: true },
    },
    {
      path: "/login",
      name: "login",
      component: () => import("@/features/login/Login.view.vue"),
      meta: { guestOnly: true },
    },
  ],
});

router.beforeEach((to) => {
  const { isAuthenticated } = useAuth();
  if (to.meta.requiresAuth && !isAuthenticated.value) {
    return { name: "login", replace: true };
  }
  if (to.meta.guestOnly && isAuthenticated.value) {
    return { name: "home", replace: true };
  }
});

export default router;
