<script setup lang="ts">
import { onMounted, ref } from "vue";
import Button from "@/shared/ui/Button.vue";
import { observeService } from "@/features/dashboard/service/observe.service";
import type { Overview } from "@/features/dashboard/types/traefik";
import type { ApiError } from "@/shared/api/http";

const overview = ref<Overview | null>(null);
const error = ref<string | null>(null);
const loading = ref(true);

const load = async () => {
  loading.value = true;
  error.value = null;
  try {
    overview.value = await observeService.getOverview();
  } catch (e) {
    const err = e as ApiError;
    error.value = err.message ?? "Impossible de charger l'overview";
  } finally {
    loading.value = false;
  }
};

onMounted(load);
</script>

<template>
  <header class="page-head">
    <div class="page-head__titles">
      <div class="eyebrow">
        <span class="page-head__live">
          <span class="page-head__live-dot" aria-hidden="true"></span>
          LIVE
        </span>
        <span>Overview</span>
        <span
          v-if="overview?.version"
          class="page-head__version mono"
          :title="`Traefik ${overview.version}`"
        >
          traefik {{ overview.version }}
        </span>
      </div>
      <h1 class="page-head__title">Traefik Telemetry</h1>
      <p class="page-head__subtitle">
        <span v-if="loading" class="mono page-head__placeholder">chargement…</span>
        <span v-else-if="error" class="page-head__error">{{ error }}</span>
        <template v-else-if="overview">
          <span class="page-head__stat mono">
            <b>{{ overview.http.routers.total }}</b> routers
          </span>
          <span class="page-head__stat mono">
            <b>{{ overview.http.services.total }}</b> services
          </span>
          <span class="page-head__stat mono">
            <b>{{ overview.http.middlewares.total }}</b> middlewares
          </span>
          <span v-if="overview.http.routers.errors > 0" class="page-head__stat page-head__stat--err mono">
            <b>{{ overview.http.routers.errors }}</b> erreurs
          </span>
        </template>
      </p>
    </div>

    <div class="page-head__actions">
      <Button variant="ghost" size="sm" @click="load" :loading="loading">
        <svg width="14" height="14" viewBox="0 0 14 14" fill="none" aria-hidden="true">
          <path d="M12 7A5 5 0 1 1 7 2" stroke="currentColor" stroke-width="1.5" stroke-linecap="round" />
          <path
            d="M9 1 12 2l-1 3"
            stroke="currentColor"
            stroke-width="1.5"
            stroke-linecap="round"
            stroke-linejoin="round"
          />
        </svg>
        Actualiser
      </Button>
    </div>
  </header>
</template>

<style scoped lang="scss">
.page-head {
  display: flex;
  align-items: flex-start;
  justify-content: space-between;
  gap: 2rem;
  flex-wrap: wrap;

  &__titles {
    min-width: 0;
  }

  &__live {
    display: inline-flex;
    align-items: center;
    gap: 0.5rem;
    padding: 0.2rem 0.55rem;
    background: $severity-danger-soft;
    color: $severity-danger;
    border-radius: $radius-sm;
    font-size: 0.95rem;
    font-weight: 600;
    letter-spacing: 0.15em;
    margin-right: 0.4rem;
  }

  &__live-dot {
    width: 5px;
    height: 5px;
    border-radius: 50%;
    background: $severity-danger;
    box-shadow: 0 0 6px $severity-danger;
    animation: graefik-pulse 1.2s $ease-in-out infinite;
  }

  &__version {
    margin-left: 0.8rem;
    padding: 0.15rem 0.55rem;
    background: $accent-soft;
    color: $accent-glow;
    border: 1px solid $border-accent;
    border-radius: $radius-sm;
    font-size: 1rem;
    font-weight: 500;
    letter-spacing: 0.06em;
    text-transform: none;
  }

  &__title {
    margin-top: 0.8rem;
    font-size: 2.6rem;
    font-weight: 600;
    letter-spacing: -0.025em;
    line-height: 1.15;
    color: $text-primary;
  }

  &__subtitle {
    margin-top: 0.6rem;
    color: $text-muted;
    font-size: 1.4rem;
    display: flex;
    align-items: center;
    gap: 1.6rem;
    flex-wrap: wrap;
  }

  &__stat {
    display: inline-flex;
    align-items: baseline;
    gap: 0.4rem;
    color: $text-secondary;
    font-size: 1.25rem;
    font-variant-numeric: tabular-nums;

    b {
      color: $text-primary;
      font-weight: 600;
      font-size: 1.35rem;
    }

    &--err {
      color: $severity-danger;

      b {
        color: $severity-danger;
      }
    }
  }

  &__placeholder {
    font-size: 1.2rem;
    color: $text-faint;
  }

  &__error {
    font-size: 1.25rem;
    color: $severity-danger;
  }

  &__actions {
    display: inline-flex;
    align-items: center;
    gap: 0.8rem;
  }
}
</style>
