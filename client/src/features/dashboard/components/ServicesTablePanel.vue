<script setup lang="ts">
import { ref } from "vue";
import Panel from "@/shared/ui/Panel.vue";
import StatusPill from "@/shared/ui/StatusPill.vue";
import Button from "@/shared/ui/Button.vue";

interface ServiceRow {
  name: string;
  provider: string;
  rps: number;
  p95: number;
  errorRate: number;
  status: "success" | "warning" | "danger";
}

const services = ref<ServiceRow[]>([
  { name: "api-gateway@docker", provider: "docker", rps: 412, p95: 38, errorRate: 0.12, status: "success" },
  { name: "users-svc@kubernetes", provider: "k8s", rps: 318, p95: 52, errorRate: 0.34, status: "success" },
  { name: "billing-svc@docker", provider: "docker", rps: 126, p95: 78, errorRate: 0.89, status: "warning" },
  { name: "search-svc@consul", provider: "consul", rps: 184, p95: 112, errorRate: 1.42, status: "warning" },
  { name: "legacy-api@file", provider: "file", rps: 48, p95: 340, errorRate: 4.8, status: "danger" },
  { name: "static-cdn@docker", provider: "docker", rps: 68, p95: 12, errorRate: 0.01, status: "success" },
]);

const providerColor: Record<string, "info" | "accent" | "muted"> = {
  docker: "info",
  k8s: "accent",
  consul: "info",
  file: "muted",
};

const statusLabel = (s: ServiceRow["status"]) =>
  s === "success" ? "healthy" : s === "warning" ? "degraded" : "failing";
</script>

<template>
  <Panel
    title="Services health"
    subtitle="Détection automatique des providers"
    eyebrow="Infrastructure"
    :span="12"
    pad="none"
  >
    <template #actions>
      <Button variant="ghost" size="sm">Tout voir</Button>
    </template>

    <div class="svc" role="table">
      <div class="svc__head" role="row">
        <span role="columnheader">Service</span>
        <span role="columnheader">Provider</span>
        <span role="columnheader" class="svc__num">req/s</span>
        <span role="columnheader" class="svc__num">p95</span>
        <span role="columnheader" class="svc__num">error %</span>
        <span role="columnheader">Status</span>
      </div>

      <div v-for="s in services" :key="s.name" class="svc__row" role="row">
        <span class="svc__name" role="cell">
          <span class="svc__dot" :class="`svc__dot--${s.status}`"></span>
          <span class="mono">{{ s.name }}</span>
        </span>
        <span role="cell">
          <StatusPill :tone="providerColor[s.provider]" :label="s.provider" :dot="false" />
        </span>
        <span role="cell" class="svc__num mono">
          {{ s.rps.toLocaleString("fr-FR") }}
        </span>
        <span role="cell" class="svc__num mono">
          {{ s.p95 }}<span class="svc__unit">ms</span>
        </span>
        <span role="cell" class="svc__num mono" :class="`svc__err--${s.status}`">
          {{ s.errorRate.toFixed(2) }}<span class="svc__unit">%</span>
        </span>
        <span role="cell">
          <StatusPill :tone="s.status" :label="statusLabel(s.status)" />
        </span>
      </div>
    </div>

    <template #footer>
      <div class="svc__foot">
        <span>6 services · 2 providers</span>
        <span>Dernier scan <span class="mono">2s</span> ago</span>
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
    grid-template-columns: 2.2fr 1fr 1fr 1fr 1fr 1.2fr;
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

  &__err--warning {
    color: $severity-warning;
  }
  &__err--danger {
    color: $severity-danger;
  }

  &__foot {
    display: flex;
    justify-content: space-between;
    gap: 1.6rem;
    flex-wrap: wrap;
  }

  @media (max-width: 900px) {
    &__head,
    &__row {
      grid-template-columns: 2fr 1fr 1fr;
      padding: 1rem;
    }
    &__head span:nth-child(n + 4),
    &__row span:nth-child(n + 4) {
      display: none;
    }
  }
}
</style>
