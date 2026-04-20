<script setup lang="ts">
import type { ApexOptions } from "apexcharts";
import { computed, ref } from "vue";
import KpiCard from "@/shared/ui/KpiCard.vue";
import Panel from "@/shared/ui/Panel.vue";
import StatusPill from "@/shared/ui/StatusPill.vue";
import Button from "@/shared/ui/Button.vue";

// ── Palette de séries : 1 accent teal + échelle neutre, statuts sémantiques ──
const series = {
  teal: "#00c2b2",
  tealGlow: "#2ce8d8",
  tealDeep: "#0b6f66",
  amber: "#f5a524",
  red: "#f16969",
  green: "#2dd4a7",
  slate: "#94a3b8",
  slateDeep: "#556978",
};

// ── Base config partagé pour tous les charts ─────────────────────────────
const baseChart: Partial<ApexOptions> = {
  chart: {
    toolbar: { show: false },
    zoom: { enabled: false },
    background: "transparent",
    fontFamily: "Geist Mono, JetBrains Mono, monospace",
    foreColor: "#5a7d8e",
    animations: {
      enabled: true,
      speed: 420,
      animateGradually: { enabled: true, delay: 120 },
    },
  },
  dataLabels: { enabled: false },
  legend: { show: false },
  grid: {
    borderColor: "rgba(255, 255, 255, 0.05)",
    strokeDashArray: 3,
    padding: { left: 10, right: 10, top: 0, bottom: 0 },
    xaxis: { lines: { show: false } },
  },
  tooltip: {
    theme: "dark",
    style: { fontSize: "1.15rem", fontFamily: "Geist Mono, monospace" },
  },
};

const axisCommon = {
  labels: {
    style: {
      colors: "#5a7d8e",
      fontSize: "1.05rem",
      fontFamily: "Geist Mono, monospace",
    },
  },
  axisBorder: { show: false },
  axisTicks: { show: false },
};

// ── KPIs (mock data, branchés à l'API Traefik plus tard) ────────────────
const spark = (base: number, variance: number, n = 24) =>
  Array.from({ length: n }, (_, i) =>
    Math.max(0, Math.round(base + Math.sin(i / 2.2) * variance + (Math.random() - 0.5) * variance * 0.6)),
  );

const kpis = ref([
  {
    label: "Requests / sec",
    value: "1 284",
    unit: "req/s",
    delta: 8.4,
    deltaLabel: "vs. 6h",
    tone: "accent" as const,
    icon: "requests" as const,
    trend: spark(1200, 180),
  },
  {
    label: "Latency p95",
    value: "42",
    unit: "ms",
    delta: -3.1,
    deltaLabel: "vs. 6h",
    tone: "info" as const,
    icon: "latency" as const,
    trend: spark(45, 10),
  },
  {
    label: "Active connections",
    value: "3 412",
    unit: "conns",
    delta: 2.6,
    deltaLabel: "vs. 6h",
    tone: "success" as const,
    icon: "connections" as const,
    trend: spark(3300, 250),
  },
  {
    label: "Error rate",
    value: "0.42",
    unit: "%",
    delta: 12.8,
    deltaLabel: "vs. 6h",
    tone: "danger" as const,
    icon: "errors" as const,
    trend: spark(18, 8),
  },
]);

// ── Chart 1 : Requests timeline (area, double série) ─────────────────────
const buildTimePoints = (n = 60, base = 900, variance = 300) => {
  const now = Date.now();
  return Array.from({ length: n }, (_, i) => ({
    x: now - (n - i) * 60_000,
    y: Math.max(
      0,
      Math.round(
        base +
          Math.sin(i / 4) * variance * 0.6 +
          Math.cos(i / 9) * variance * 0.3 +
          (Math.random() - 0.5) * variance * 0.4,
      ),
    ),
  }));
};

const requestsSeries = ref([
  { name: "2xx/3xx", data: buildTimePoints(60, 1100, 240) },
  { name: "4xx", data: buildTimePoints(60, 90, 40) },
  { name: "5xx", data: buildTimePoints(60, 14, 10) },
]);

const requestsOptions = computed<ApexOptions>(() => ({
  ...baseChart,
  chart: { ...baseChart.chart, type: "area", stacked: true, height: 320 },
  colors: [series.teal, series.amber, series.red],
  stroke: { curve: "smooth", width: 2, lineCap: "round" },
  fill: {
    type: "gradient",
    gradient: {
      shadeIntensity: 1,
      opacityFrom: 0.35,
      opacityTo: 0.02,
      stops: [0, 95],
    },
  },
  xaxis: {
    ...axisCommon,
    type: "datetime",
    labels: {
      ...axisCommon.labels,
      datetimeUTC: false,
      format: "HH:mm",
    },
    tooltip: { enabled: false },
  },
  yaxis: {
    ...axisCommon,
    labels: {
      ...axisCommon.labels,
      formatter: (v: number) => `${v.toFixed(0)}`,
    },
  },
  tooltip: {
    ...baseChart.tooltip,
    x: { format: "HH:mm:ss" },
    y: { formatter: (v: number) => `${v.toLocaleString("fr-FR")} req` },
  },
  markers: { size: 0, hover: { size: 5 } },
}));

// ── Chart 2 : Status codes (donut) ───────────────────────────────────────
const statusSeries = ref([73245, 18420, 6210, 892, 148]);
const statusOptions = computed<ApexOptions>(() => ({
  ...baseChart,
  chart: { ...baseChart.chart, type: "donut", height: 320 },
  labels: ["2xx OK", "3xx Redirect", "4xx Client", "5xx Server", "Timeouts"],
  colors: [series.teal, series.tealDeep, series.amber, series.red, series.slate],
  stroke: { width: 2, colors: ["#0a2236"] },
  plotOptions: {
    pie: {
      donut: {
        size: "72%",
        labels: {
          show: true,
          name: {
            show: true,
            fontSize: "1.1rem",
            fontFamily: "Geist Mono, monospace",
            color: "#5a7d8e",
            offsetY: -4,
          },
          value: {
            show: true,
            fontSize: "2.4rem",
            fontFamily: "Geist, sans-serif",
            color: "#e8f4f3",
            fontWeight: 600,
            offsetY: 8,
            formatter: (v: string) => Number(v).toLocaleString("fr-FR"),
          },
          total: {
            show: true,
            showAlways: false,
            label: "Total",
            color: "#5a7d8e",
            fontSize: "1.1rem",
            fontFamily: "Geist Mono, monospace",
            formatter: (w: any) =>
              w.globals.seriesTotals
                .reduce((a: number, b: number) => a + b, 0)
                .toLocaleString("fr-FR"),
          },
        },
      },
    },
  },
  tooltip: {
    ...baseChart.tooltip,
    y: { formatter: (v: number) => `${v.toLocaleString("fr-FR")} hits` },
  },
  legend: { show: false },
  dataLabels: { enabled: false },
}));

// ── Chart 3 : Latency percentiles (line) ─────────────────────────────────
const latencySeries = ref([
  { name: "p50", data: Array.from({ length: 48 }, () => 10 + Math.random() * 6) },
  { name: "p95", data: Array.from({ length: 48 }, () => 32 + Math.random() * 18) },
  { name: "p99", data: Array.from({ length: 48 }, () => 80 + Math.random() * 60) },
]);

const latencyCategories = Array.from({ length: 48 }, (_, i) => {
  const d = new Date(Date.now() - (48 - i) * 300_000);
  return d.toLocaleTimeString("fr-FR", { hour: "2-digit", minute: "2-digit" });
});

const latencyOptions = computed<ApexOptions>(() => ({
  ...baseChart,
  chart: { ...baseChart.chart, type: "line", height: 320 },
  colors: [series.slate, series.teal, series.amber],
  stroke: { curve: "smooth", width: [2, 2.5, 2], dashArray: [0, 0, 4] },
  xaxis: {
    ...axisCommon,
    categories: latencyCategories,
    labels: {
      ...axisCommon.labels,
      rotate: 0,
      hideOverlappingLabels: true,
    },
    tickAmount: 8,
  },
  yaxis: {
    ...axisCommon,
    labels: {
      ...axisCommon.labels,
      formatter: (v: number) => `${v.toFixed(0)}ms`,
    },
  },
  tooltip: {
    ...baseChart.tooltip,
    y: { formatter: (v: number) => `${v.toFixed(1)} ms` },
  },
  markers: { size: 0, hover: { size: 5 } },
  legend: {
    show: true,
    position: "top",
    horizontalAlign: "right",
    fontFamily: "Geist Mono, monospace",
    fontSize: "1.15rem",
    labels: { colors: "#93b4c2" },
    markers: {
      strokeWidth: 0,
    },
  },
}));

// ── Chart 4 : Top routers (bar horizontal) ───────────────────────────────
const topRouters = ref([
  { name: "api.myapp.com", rps: 412 },
  { name: "app.myapp.com", rps: 318 },
  { name: "grafana.internal", rps: 184 },
  { name: "auth.myapp.com", rps: 126 },
  { name: "docs.myapp.com", rps: 94 },
  { name: "webhooks.myapp", rps: 82 },
  { name: "cdn.myapp.com", rps: 68 },
]);

const routersSeries = computed(() => [
  { name: "Requests", data: topRouters.value.map((r) => r.rps) },
]);

const routersOptions = computed<ApexOptions>(() => ({
  ...baseChart,
  chart: { ...baseChart.chart, type: "bar", height: 320 },
  colors: [series.teal],
  plotOptions: {
    bar: {
      horizontal: true,
      borderRadius: 4,
      borderRadiusApplication: "end",
      barHeight: "62%",
      colors: {
        backgroundBarColors: ["rgba(255,255,255,0.02)"],
        backgroundBarOpacity: 1,
        backgroundBarRadius: 4,
      },
    },
  },
  fill: {
    type: "gradient",
    gradient: {
      shade: "dark",
      type: "horizontal",
      gradientToColors: [series.tealGlow],
      stops: [0, 100],
    },
  },
  xaxis: {
    ...axisCommon,
    categories: topRouters.value.map((r) => r.name),
    labels: {
      ...axisCommon.labels,
      formatter: (v: string) => `${v} r/s`,
    },
  },
  yaxis: {
    ...axisCommon,
    labels: {
      ...axisCommon.labels,
      style: {
        ...axisCommon.labels.style,
        fontFamily: "Geist Mono, monospace",
      },
    },
  },
  tooltip: {
    ...baseChart.tooltip,
    y: { formatter: (v: number) => `${v} req/s` },
  },
  dataLabels: {
    enabled: true,
    textAnchor: "end",
    offsetX: -8,
    style: {
      colors: ["#e8f4f3"],
      fontSize: "1.1rem",
      fontFamily: "Geist Mono, monospace",
      fontWeight: 600,
    },
    formatter: (v: number) => `${v}`,
  },
}));

// ── Services health table ────────────────────────────────────────────────
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

const providerColor: Record<string, string> = {
  docker: "info",
  k8s: "accent",
  consul: "info",
  file: "muted",
};
</script>

<template>
  <div class="graefik-dashboard">
    <!-- Page header ─────────────────────────────────────────────────── -->
    <header class="page-head">
      <div class="page-head__titles">
        <div class="eyebrow">
          <span class="page-head__live">
            <span class="page-head__live-dot" aria-hidden="true"></span>
            LIVE
          </span>
          <span>Observatory · Overview</span>
        </div>
        <h1 class="page-head__title">Traefik Telemetry</h1>
        <p class="page-head__subtitle">
          Vue d'ensemble en temps réel de vos routers, services et middlewares.
          <span class="mono page-head__last">Mis à jour {{ new Date().toLocaleTimeString("fr-FR") }}</span>
        </p>
      </div>

      <div class="page-head__actions">
        <Button variant="subtle" size="sm">
          <svg width="14" height="14" viewBox="0 0 14 14" fill="none" aria-hidden="true">
            <path d="M7 1v8m0 0L4 6m3 3 3-3" stroke="currentColor" stroke-width="1.5" stroke-linecap="round" stroke-linejoin="round" />
            <path d="M2 11h10" stroke="currentColor" stroke-width="1.5" stroke-linecap="round" />
          </svg>
          Exporter
        </Button>
        <Button variant="ghost" size="sm">
          <svg width="14" height="14" viewBox="0 0 14 14" fill="none" aria-hidden="true">
            <path
              d="M12 7A5 5 0 1 1 7 2"
              stroke="currentColor"
              stroke-width="1.5"
              stroke-linecap="round"
            />
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
        <Button size="sm">
          <svg width="14" height="14" viewBox="0 0 14 14" fill="none" aria-hidden="true">
            <path d="M7 3v8M3 7h8" stroke="currentColor" stroke-width="1.8" stroke-linecap="round" />
          </svg>
          Nouveau panneau
        </Button>
      </div>
    </header>

    <!-- KPIs row ────────────────────────────────────────────────────── -->
    <section class="kpi-row">
      <KpiCard
        v-for="(kpi, idx) in kpis"
        :key="kpi.label"
        :label="kpi.label"
        :value="kpi.value"
        :unit="kpi.unit"
        :delta="kpi.delta"
        :delta-label="kpi.deltaLabel"
        :tone="kpi.tone"
        :icon="kpi.icon"
        :trend="kpi.trend"
        :hero="idx === 0"
      />
    </section>

    <!-- Grid principal ─────────────────────────────────────────────── -->
    <section class="grid">
      <Panel
        title="Requests over time"
        subtitle="Stacked par famille de code HTTP"
        eyebrow="Timeline"
        :span="8"
        live
        pad="tight"
      >
        <template #actions>
          <div class="panel-chips">
            <span class="panel-chip panel-chip--teal">
              <span class="panel-chip__dot"></span> 2xx/3xx
            </span>
            <span class="panel-chip panel-chip--amber">
              <span class="panel-chip__dot"></span> 4xx
            </span>
            <span class="panel-chip panel-chip--red">
              <span class="panel-chip__dot"></span> 5xx
            </span>
          </div>
        </template>
        <apexchart
          type="area"
          height="320"
          width="100%"
          :options="requestsOptions"
          :series="requestsSeries"
        />
      </Panel>

      <Panel
        title="Status codes"
        subtitle="Distribution sur 6h"
        eyebrow="HTTP"
        :span="4"
        pad="tight"
      >
        <apexchart
          type="donut"
          height="320"
          width="100%"
          :options="statusOptions"
          :series="statusSeries"
        />
        <template #footer>
          <div class="legend">
            <span class="legend__item"><span class="legend__dot" style="background:#00c2b2"></span> 2xx</span>
            <span class="legend__item"><span class="legend__dot" style="background:#0b6f66"></span> 3xx</span>
            <span class="legend__item"><span class="legend__dot" style="background:#f5a524"></span> 4xx</span>
            <span class="legend__item"><span class="legend__dot" style="background:#f16969"></span> 5xx</span>
            <span class="legend__item"><span class="legend__dot" style="background:#94a3b8"></span> timeout</span>
          </div>
        </template>
      </Panel>

      <Panel
        title="Response time percentiles"
        subtitle="p50 · p95 · p99 — fenêtre glissante 4h"
        eyebrow="Latency"
        :span="7"
        pad="tight"
      >
        <apexchart
          type="line"
          height="320"
          width="100%"
          :options="latencyOptions"
          :series="latencySeries"
        />
      </Panel>

      <Panel
        title="Top routers"
        subtitle="Par requêtes / seconde"
        eyebrow="Traffic"
        :span="5"
        pad="tight"
      >
        <apexchart
          type="bar"
          height="320"
          width="100%"
          :options="routersOptions"
          :series="routersSeries"
        />
      </Panel>

      <!-- Services table ─────────────────────────────────────── -->
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
        <div class="svc-table" role="table">
          <div class="svc-table__head" role="row">
            <span role="columnheader">Service</span>
            <span role="columnheader">Provider</span>
            <span role="columnheader" class="svc-table__num">req/s</span>
            <span role="columnheader" class="svc-table__num">p95</span>
            <span role="columnheader" class="svc-table__num">error %</span>
            <span role="columnheader">Status</span>
          </div>

          <div
            v-for="s in services"
            :key="s.name"
            class="svc-table__row"
            role="row"
          >
            <span class="svc-table__name" role="cell">
              <span class="svc-table__name-dot" :class="`svc-table__name-dot--${s.status}`"></span>
              <span class="mono">{{ s.name }}</span>
            </span>
            <span role="cell">
              <StatusPill
                :tone="(providerColor[s.provider] as any)"
                :label="s.provider"
                :dot="false"
              />
            </span>
            <span role="cell" class="svc-table__num mono">
              {{ s.rps.toLocaleString("fr-FR") }}
            </span>
            <span role="cell" class="svc-table__num mono">{{ s.p95 }}<span class="svc-table__unit">ms</span></span>
            <span role="cell" class="svc-table__num mono" :class="`svc-table__err--${s.status}`">
              {{ s.errorRate.toFixed(2) }}<span class="svc-table__unit">%</span>
            </span>
            <span role="cell">
              <StatusPill
                :tone="s.status"
                :label="s.status === 'success' ? 'healthy' : s.status === 'warning' ? 'degraded' : 'failing'"
              />
            </span>
          </div>
        </div>

        <template #footer>
          <div class="svc-table__foot">
            <span>6 services · 2 providers</span>
            <span>Dernier scan <span class="mono">2s</span> ago</span>
          </div>
        </template>
      </Panel>
    </section>
  </div>
</template>

<style scoped lang="scss">
.graefik-dashboard {
  display: flex;
  flex-direction: column;
  gap: 2.4rem;
}

// ── Page head ────────────────────────────────────────────────────
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
    gap: 1rem;
    flex-wrap: wrap;
  }

  &__last {
    font-size: 1.15rem;
    color: $text-faint;
    padding: 0.2rem 0.6rem;
    border-radius: $radius-sm;
    background: $bg-inset;
    border: 1px solid $border-subtle;
  }

  &__actions {
    display: inline-flex;
    align-items: center;
    gap: 0.8rem;
  }
}

// ── KPI row ────────────────────────────────────────────────────
.kpi-row {
  display: grid;
  grid-template-columns: repeat(4, minmax(0, 1fr));
  gap: 1.6rem;

  @media (max-width: 1200px) {
    grid-template-columns: repeat(2, 1fr);
  }
  @media (max-width: 640px) {
    grid-template-columns: 1fr;
  }
}

// ── Grid principal ─────────────────────────────────────────────
.grid {
  display: grid;
  grid-template-columns: repeat(12, minmax(0, 1fr));
  gap: 1.6rem;
}

// ── Chips pour header de panel ────────────────────────────────
.panel-chips {
  display: inline-flex;
  gap: 0.4rem;
}

.panel-chip {
  display: inline-flex;
  align-items: center;
  gap: 0.4rem;
  padding: 0.25rem 0.6rem;
  background: $bg-inset;
  border: 1px solid $border-subtle;
  border-radius: $radius-sm;
  font-family: $font-mono;
  font-size: 1.05rem;
  color: $text-secondary;

  &__dot {
    width: 6px;
    height: 6px;
    border-radius: 2px;
  }

  &--teal &__dot {
    background: $series-1;
  }
  &--amber &__dot {
    background: $severity-warning;
  }
  &--red &__dot {
    background: $severity-danger;
  }
}

// ── Legend ─────────────────────────────────────────────────────
.legend {
  display: flex;
  align-items: center;
  gap: 1.6rem;
  flex-wrap: wrap;
  padding: 0.2rem 0;

  &__item {
    display: inline-flex;
    align-items: center;
    gap: 0.5rem;
    font-family: $font-mono;
    font-size: 1.1rem;
    color: $text-secondary;
    text-transform: uppercase;
    letter-spacing: 0.05em;
  }

  &__dot {
    width: 8px;
    height: 8px;
    border-radius: 2px;
  }
}

// ── Services table ────────────────────────────────────────────
.svc-table {
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

  &__name-dot {
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
