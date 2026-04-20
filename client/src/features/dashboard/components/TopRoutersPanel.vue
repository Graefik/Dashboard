<script setup lang="ts">
import type { ApexOptions } from "apexcharts";
import { computed, ref } from "vue";
import Panel from "@/shared/ui/Panel.vue";
import { axisCommon, baseChart, series } from "@/features/dashboard/theme";

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

const options = computed<ApexOptions>(() => ({
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
    labels: { ...axisCommon.labels, formatter: (v: string) => `${v} r/s` },
  },
  yaxis: { ...axisCommon },
  tooltip: { ...baseChart.tooltip, y: { formatter: (v: number) => `${v} req/s` } },
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
</script>

<template>
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
      :options="options"
      :series="routersSeries"
    />
  </Panel>
</template>
