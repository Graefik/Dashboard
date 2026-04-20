<script setup lang="ts">
import { ref } from "vue";
import Panel from "@/shared/ui/Panel.vue";
import PanelChips from "@/shared/ui/PanelChips.vue";
import UPlot, { type UPlotSeries } from "@/shared/ui/UPlot.vue";
import { series } from "@/features/dashboard/theme";
import { genTs } from "@/features/dashboard/utils/mock";

const xs = genTs(48, 300_000);
const data = ref<number[][]>([
  xs,
  Array.from({ length: 48 }, () => 10 + Math.random() * 6),
  Array.from({ length: 48 }, () => 32 + Math.random() * 18),
  Array.from({ length: 48 }, () => 80 + Math.random() * 60),
]);

const uplotSeries: UPlotSeries[] = [
  { label: "p50", color: series.slate, width: 2, fill: null },
  { label: "p95", color: series.teal, width: 2.5, fill: null },
  { label: "p99", color: series.amber, width: 2, fill: null, dash: [4, 4] },
];

const chips = [
  { label: "p50", color: series.slate },
  { label: "p95", color: series.teal },
  { label: "p99", color: series.amber, dash: true },
];
</script>

<template>
  <Panel
    title="Response time percentiles"
    subtitle="p50 · p95 · p99 — fenêtre glissante 4h · rendu uPlot"
    eyebrow="Latency"
    :span="7"
    pad="tight"
  >
    <template #actions>
      <PanelChips :chips="chips" />
    </template>
    <UPlot
      :data="data"
      :series="uplotSeries"
      :height="300"
      :y-formatter="(v: number) => v.toFixed(0) + 'ms'"
    />
  </Panel>
</template>
