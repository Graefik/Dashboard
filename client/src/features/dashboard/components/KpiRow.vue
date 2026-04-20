<script setup lang="ts">
import { computed } from "vue";
import KpiCard from "@/shared/ui/KpiCard.vue";
import { useMetrics } from "@/features/dashboard/composables/useMetrics";

const { overview } = useMetrics();

const fmtNumber = (n: number) => {
  if (n >= 1000) return n.toLocaleString("fr-FR", { maximumFractionDigits: 0 });
  return n.toLocaleString("fr-FR", { maximumFractionDigits: 1 });
};

const kpis = computed(() => {
  const k = overview.value?.kpis;
  const trend = (): number[] => [];
  if (!k) {
    return [
      { label: "Requests / sec", value: "—", unit: "req/s", delta: 0, tone: "accent" as const, icon: "requests" as const, trend: trend() },
      { label: "Latency p95", value: "—", unit: "ms", delta: 0, tone: "info" as const, icon: "latency" as const, trend: trend() },
      { label: "Active connections", value: "—", unit: "conns", delta: 0, tone: "success" as const, icon: "connections" as const, trend: trend() },
      { label: "Error rate", value: "—", unit: "%", delta: 0, tone: "danger" as const, icon: "errors" as const, trend: trend() },
    ];
  }
  const reqSeries = overview.value?.requests.series ?? {};
  const latencySeries = overview.value?.latency.p95 ?? [];

  const all2xx = reqSeries["2xx"] ?? [];
  const all4xx = reqSeries["4xx"] ?? [];
  const all5xx = reqSeries["5xx"] ?? [];
  const total = all2xx.map((v, i) => v + (all4xx[i] ?? 0) + (all5xx[i] ?? 0));

  return [
    {
      label: "Requests / sec",
      value: fmtNumber(k.rps),
      unit: "req/s",
      delta: +k.rpsDeltaPct.toFixed(1),
      deltaLabel: "vs. 5m",
      tone: "accent" as const,
      icon: "requests" as const,
      trend: total.slice(-24),
    },
    {
      label: "Latency p95",
      value: fmtNumber(k.latencyP95Ms),
      unit: "ms",
      delta: +k.latencyP95DeltaPct.toFixed(1),
      deltaLabel: "vs. 5m",
      tone: "info" as const,
      icon: "latency" as const,
      trend: latencySeries.slice(-24),
    },
    {
      label: "Active connections",
      value: fmtNumber(k.connections),
      unit: "conns",
      delta: +k.connectionsDeltaPct.toFixed(1),
      deltaLabel: "vs. 5m",
      tone: "success" as const,
      icon: "connections" as const,
      trend: [],
    },
    {
      label: "Error rate",
      value: k.errorRatePct.toFixed(2),
      unit: "%",
      delta: +k.errorRateDeltaPct.toFixed(1),
      deltaLabel: "vs. 5m",
      tone: "danger" as const,
      icon: "errors" as const,
      trend: (reqSeries["5xx"] ?? []).slice(-24),
    },
  ];
});
</script>

<template>
  <section class="kpi-row">
    <KpiCard
      v-for="(kpi, idx) in kpis"
      :key="kpi.label"
      :label="kpi.label"
      :value="kpi.value"
      :unit="kpi.unit"
      :delta="kpi.delta"
      :delta-label="(kpi as any).deltaLabel"
      :tone="kpi.tone"
      :icon="kpi.icon"
      :trend="kpi.trend"
      :hero="idx === 0"
    />
  </section>
</template>

<style scoped lang="scss">
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
</style>
