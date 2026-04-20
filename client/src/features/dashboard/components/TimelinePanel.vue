<script setup lang="ts">
import { ref } from "vue";
import Panel from "@/shared/ui/Panel.vue";
import PanelChips from "@/shared/ui/PanelChips.vue";
import UPlot, { type UPlotSeries } from "@/shared/ui/UPlot.vue";
import { series } from "@/features/dashboard/theme";
import { genTs, genY } from "@/features/dashboard/utils/mock";

const xs = genTs(60, 60_000);
const data = ref<number[][]>([
  xs,
  genY(60, 1100, 240),
  genY(60, 90, 40),
  genY(60, 14, 10),
]);

const uplotSeries: UPlotSeries[] = [
  { label: "2xx/3xx", color: series.teal, width: 2 },
  { label: "4xx", color: series.amber, width: 2 },
  { label: "5xx", color: series.red, width: 2 },
];

const chips = [
  { label: "2xx/3xx", color: series.teal },
  { label: "4xx", color: series.amber },
  { label: "5xx", color: series.red },
];
</script>

<template>
  <Panel
    title="Requests over time"
    subtitle="Stacked par famille de code HTTP · rendu uPlot"
    eyebrow="Timeline"
    :span="8"
    live
    pad="tight"
  >
    <template #actions>
      <PanelChips :chips="chips" />
    </template>
    <UPlot
      :data="data"
      :series="uplotSeries"
      :height="300"
      stacked
      :y-formatter="(v: number) => v.toLocaleString('fr-FR') + ' req'"
    />
  </Panel>
</template>
