import { onBeforeUnmount, onMounted, ref } from "vue";
import { request, type ApiError } from "@/shared/api/http";

export interface SystemInfo {
  traefikURL: string;
}

const info = ref<SystemInfo | null>(null);
const connected = ref<boolean>(true);
let refs = 0;
let healthTimer: number | undefined;

const fetchInfo = async () => {
  try {
    info.value = await request<SystemInfo>("/api/system/info", { auth: true });
  } catch {
    // 401 tant que pas loggé : on ignore silencieusement
  }
};

const checkHealth = async () => {
  try {
    await request<unknown>("/healthz");
    connected.value = true;
  } catch (e) {
    const err = e as ApiError;
    connected.value = err.code !== "NETWORK_ERROR";
  }
};

export const useSystem = () => {
  onMounted(() => {
    refs++;
    if (refs === 1) {
      fetchInfo();
      checkHealth();
      healthTimer = window.setInterval(checkHealth, 15_000);
    }
  });
  onBeforeUnmount(() => {
    refs--;
    if (refs <= 0 && healthTimer !== undefined) {
      clearInterval(healthTimer);
      healthTimer = undefined;
    }
  });

  return {
    info,
    connected,
    refresh: fetchInfo,
  };
};
