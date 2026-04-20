import { computed, onBeforeUnmount, onMounted, ref } from "vue";
import { metricsService } from "@/features/dashboard/service/metrics.service";
import type { MetricsOverview } from "@/features/dashboard/types/metrics";
import type { ApiError } from "@/shared/api/http";

const REFRESH_INTERVAL_MS = 10_000;

// Singleton partagé entre tous les panels — 1 seul fetch par tick.
const overview = ref<MetricsOverview | null>(null);
const loading = ref(true);
const error = ref<string | null>(null);
const lastFetch = ref<Date | null>(null);

// Sélection partagée du router courant (clic dans TopRoutersPanel).
const selectedRouter = ref<string | null>(null);

const selectedRouterData = computed(() => {
  const name = selectedRouter.value;
  if (!name) return null;
  return overview.value?.routers.find((r) => r.name === name) ?? null;
});

let refs = 0;
let timer: number | undefined;

const fetchOnce = async () => {
  try {
    overview.value = await metricsService.getOverview();
    error.value = null;
    lastFetch.value = new Date();

    // auto-select du premier router si rien n'est encore sélectionné
    // ou si le router sélectionné n'est plus dans le top
    const names = overview.value.routers.map((r) => r.name);
    if (!selectedRouter.value || !names.includes(selectedRouter.value)) {
      selectedRouter.value = names[0] ?? null;
    }
  } catch (e) {
    const err = e as ApiError;
    error.value = err.message ?? "Metrics indisponibles";
  } finally {
    loading.value = false;
  }
};

const start = () => {
  if (timer !== undefined) return;
  fetchOnce();
  timer = window.setInterval(fetchOnce, REFRESH_INTERVAL_MS);
};

const stop = () => {
  if (timer !== undefined) {
    clearInterval(timer);
    timer = undefined;
  }
};

export const useMetrics = () => {
  onMounted(() => {
    refs++;
    start();
  });
  onBeforeUnmount(() => {
    refs--;
    if (refs <= 0) stop();
  });

  return {
    overview,
    loading,
    error,
    lastFetch,
    refresh: fetchOnce,
    selectedRouter,
    selectedRouterData,
  };
};
