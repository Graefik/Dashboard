<script setup lang="ts">
import type { ApexOptions } from "apexcharts";
import { computed, ref } from "vue";
import Panel from "@/shared/ui/Panel.vue";
import { baseChart, series } from "@/features/dashboard/theme";

const data = ref([73245, 18420, 6210, 892, 148]);

const options = computed<ApexOptions>(() => ({
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

const legend = [
  { label: "2xx", color: series.teal },
  { label: "3xx", color: series.tealDeep },
  { label: "4xx", color: series.amber },
  { label: "5xx", color: series.red },
  { label: "timeout", color: series.slate },
];
</script>

<template>
  <Panel
    title="Status codes"
    subtitle="Distribution sur 6h"
    eyebrow="HTTP"
    :span="4"
    pad="tight"
  >
    <apexchart type="donut" height="320" width="100%" :options="options" :series="data" />
    <template #footer>
      <div class="legend">
        <span v-for="l in legend" :key="l.label" class="legend__item">
          <span class="legend__dot" :style="{ background: l.color }"></span>
          {{ l.label }}
        </span>
      </div>
    </template>
  </Panel>
</template>

<style scoped lang="scss">
.legend {
  display: flex;
  align-items: center;
  gap: 1.6rem;
  flex-wrap: wrap;

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
</style>
