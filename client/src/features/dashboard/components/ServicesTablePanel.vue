<script setup lang="ts">
import { computed, onMounted, ref } from "vue";
import Panel from "@/shared/ui/Panel.vue";
import StatusPill from "@/shared/ui/StatusPill.vue";
import Button from "@/shared/ui/Button.vue";
import { observeService } from "@/features/dashboard/service/observe.service";
import type { Service } from "@/features/dashboard/types/traefik";
import type { ApiError } from "@/shared/api/http";

const services = ref<Service[]>([]);
const loading = ref(true);
const error = ref<string | null>(null);
const lastScan = ref<Date | null>(null);

const providerTone = (p: string): "info" | "accent" | "muted" => {
  if (p === "docker") return "info";
  if (p === "kubernetes" || p === "kubernetescrd") return "accent";
  if (p === "consul" || p === "consulcatalog") return "info";
  return "muted";
};

const toStatus = (s: string): "success" | "warning" | "danger" => {
  if (s === "enabled") return "success";
  if (s === "warning") return "warning";
  return "danger";
};

const statusLabel = (s: string) => {
  const t = toStatus(s);
  if (t === "success") return "healthy";
  if (t === "warning") return "degraded";
  return "failing";
};

// Un service est "up" si tous ses serverStatus valent "UP"
const liveness = (svc: Service): number => {
  if (!svc.serverStatus) return 0;
  const values = Object.values(svc.serverStatus);
  if (values.length === 0) return 0;
  const up = values.filter((v) => v.toUpperCase() === "UP").length;
  return up / values.length;
};

const rows = computed(() => {
  return services.value.map((s) => {
    const lv = liveness(s);
    const status: "success" | "warning" | "danger" =
      lv === 1
        ? toStatus(s.status)
        : lv > 0
          ? "warning"
          : s.status === "enabled"
            ? toStatus(s.status)
            : "danger";
    return {
      name: s.name,
      provider: s.provider,
      servers: s.serverStatus ? Object.keys(s.serverStatus).length : 0,
      liveness: lv,
      status,
    };
  });
});

const providersCount = computed(
  () => new Set(services.value.map((s) => s.provider)).size,
);

const load = async () => {
  loading.value = true;
  error.value = null;
  try {
    services.value = await observeService.listServices();
    lastScan.value = new Date();
  } catch (e) {
    const err = e as ApiError;
    error.value = err.message ?? "Impossible de charger les services";
  } finally {
    loading.value = false;
  }
};

onMounted(load);
</script>

<template>
  <Panel
    title="Services health"
    subtitle="Détection automatique des providers Traefik"
    eyebrow="Infrastructure"
    :span="12"
    pad="none"
  >
    <template #actions>
      <Button variant="ghost" size="sm" @click="load" :loading="loading">
        Actualiser
      </Button>
    </template>

    <div v-if="error" class="svc__error">{{ error }}</div>

    <div v-else class="svc" role="table">
      <div class="svc__head" role="row">
        <span role="columnheader">Service</span>
        <span role="columnheader">Provider</span>
        <span role="columnheader" class="svc__num">servers</span>
        <span role="columnheader" class="svc__num">live</span>
        <span role="columnheader">Status</span>
      </div>

      <div
        v-if="loading && !rows.length"
        class="svc__skeleton"
      >
        Chargement…
      </div>

      <div v-else-if="!rows.length" class="svc__empty">
        Aucun service détecté.
      </div>

      <div v-for="s in rows" :key="s.name" class="svc__row" role="row">
        <span class="svc__name" role="cell">
          <span class="svc__dot" :class="`svc__dot--${s.status}`"></span>
          <span class="mono">{{ s.name }}</span>
        </span>
        <span role="cell">
          <StatusPill :tone="providerTone(s.provider)" :label="s.provider" :dot="false" />
        </span>
        <span role="cell" class="svc__num mono">
          {{ s.servers || "—" }}
        </span>
        <span role="cell" class="svc__num mono" :class="{ 'svc__warn': s.liveness < 1 && s.servers > 0 }">
          <template v-if="s.servers > 0">{{ Math.round(s.liveness * 100) }}<span class="svc__unit">%</span></template>
          <template v-else>—</template>
        </span>
        <span role="cell">
          <StatusPill :tone="s.status" :label="statusLabel(s.status)" />
        </span>
      </div>
    </div>

    <template #footer>
      <div class="svc__foot">
        <span>{{ services.length }} services · {{ providersCount }} providers</span>
        <span v-if="lastScan">
          Dernier scan
          <span class="mono">
            {{ lastScan.toLocaleTimeString("fr-FR") }}
          </span>
        </span>
      </div>
    </template>
  </Panel>
</template>

<style scoped lang="scss">
.svc {
  display: flex;
  flex-direction: column;

  &__head,
  &__row {
    display: grid;
    grid-template-columns: 2.4fr 1fr 1fr 1fr 1.2fr;
    align-items: center;
    gap: 1.2rem;
    padding: 1.1rem 1.8rem;
  }

  &__head {
    background: $bg-inset;
    border-bottom: 1px solid $border-subtle;
    font-family: $font-mono;
    font-size: 1rem;
    color: $text-muted;
    text-transform: uppercase;
    letter-spacing: 0.1em;
  }

  &__row {
    border-bottom: 1px solid $border-subtle;
    transition: background 0.15s $ease-out;

    &:last-child {
      border-bottom: 0;
    }
    &:hover {
      background: rgba(0, 194, 178, 0.025);
    }
  }

  &__name {
    display: inline-flex;
    align-items: center;
    gap: 0.8rem;
    min-width: 0;
    color: $text-primary;
    font-size: 1.3rem;
    overflow: hidden;

    .mono {
      overflow: hidden;
      text-overflow: ellipsis;
      white-space: nowrap;
    }
  }

  &__dot {
    width: 6px;
    height: 6px;
    border-radius: 50%;
    flex-shrink: 0;

    &--success {
      background: $severity-success;
      box-shadow: 0 0 6px $severity-success;
    }
    &--warning {
      background: $severity-warning;
      box-shadow: 0 0 6px $severity-warning;
    }
    &--danger {
      background: $severity-danger;
      box-shadow: 0 0 6px $severity-danger;
      animation: graefik-pulse 1.4s $ease-in-out infinite;
    }
  }

  &__num {
    text-align: right;
    font-size: 1.3rem;
    color: $text-primary;
    font-variant-numeric: tabular-nums;
  }

  &__unit {
    color: $text-muted;
    margin-left: 0.2rem;
    font-size: 1.1rem;
  }

  &__warn {
    color: $severity-warning;
  }

  &__foot {
    display: flex;
    justify-content: space-between;
    gap: 1.6rem;
    flex-wrap: wrap;
  }

  &__skeleton,
  &__empty,
  &__error {
    padding: 2.4rem 1.8rem;
    text-align: center;
    font-family: $font-mono;
    font-size: 1.3rem;
    color: $text-muted;
  }

  &__error {
    color: $severity-danger;
  }
}
</style>
