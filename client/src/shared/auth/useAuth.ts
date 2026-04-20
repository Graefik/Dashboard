import { computed, ref } from "vue";
import { clearToken, getToken, setToken } from "@/shared/api/http";

const USER_KEY = "graefik.username";
const EXP_KEY = "graefik.expiresAt";

const tokenRef = ref<string | null>(getToken());
const usernameRef = ref<string | null>(localStorage.getItem(USER_KEY));
const expiresAtRef = ref<number | null>(
  Number(localStorage.getItem(EXP_KEY)) || null,
);

const isExpired = () => {
  if (!expiresAtRef.value) return false;
  return Date.now() / 1000 >= expiresAtRef.value;
};

const isAuthenticated = computed(
  () => !!tokenRef.value && !isExpired(),
);

const login = (token: string, username: string, expiresAt: number) => {
  setToken(token);
  localStorage.setItem(USER_KEY, username);
  localStorage.setItem(EXP_KEY, String(expiresAt));
  tokenRef.value = token;
  usernameRef.value = username;
  expiresAtRef.value = expiresAt;
};

const logout = () => {
  clearToken();
  localStorage.removeItem(USER_KEY);
  localStorage.removeItem(EXP_KEY);
  tokenRef.value = null;
  usernameRef.value = null;
  expiresAtRef.value = null;
};

export const useAuth = () => ({
  token: tokenRef,
  username: usernameRef,
  isAuthenticated,
  login,
  logout,
});
