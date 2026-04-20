<script setup lang="ts">
import { ref } from "vue";
import KpiCard from "@/shared/ui/KpiCard.vue";
import { spark } from "@/features/dashboard/utils/mock";

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
      :delta-label="kpi.deltaLabel"
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
