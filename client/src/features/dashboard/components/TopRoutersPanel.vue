<script setup lang="ts">
import type { ApexOptions } from "apexcharts";
import { computed } from "vue";
import Panel from "@/shared/ui/Panel.vue";
import { axisCommon, baseChart, series } from "@/features/dashboard/theme";
import { useMetrics } from "@/features/dashboard/composables/useMetrics";

const { overview, selectedRouter } = useMetrics();

// Top = 7 premiers routers actifs (liste globale déjà triée par RPS côté back)
const routers = computed(() => (overview.value?.routers ?? []).slice(0, 7));

const routersSeries = computed(() => [
  { name: "RPS", data: routers.value.map((r) => +r.rps.toFixed(2)) },
]);

const fmtRps = (v: number): string => {
  if (v === 0) return "0";
  if (v < 0.01) return v.toFixed(3);
  if (v < 1) return v.toFixed(2);
  if (v < 10) return v.toFixed(2);
  return v.toFixed(1);
};

const options = computed<ApexOptions>(() => {
  const selectedIdx = routers.value.findIndex(
    (r) => r.name === selectedRouter.value,
  );
  return {
    ...baseChart,
    chart: {
      ...baseChart.chart,
      type: "bar",
      height: 320,
      events: {
        dataPointSelection: (_e, _ctx, cfg: { dataPointIndex: number }) => {
          const r = routers.value[cfg.dataPointIndex];
          if (r) selectedRouter.value = r.name;
        },
        click: (_e, _ctx, cfg: { dataPointIndex?: number }) => {
          // fallback pour les clics hors plotArea (sur la catégorie y)
          if (cfg.dataPointIndex != null && cfg.dataPointIndex >= 0) {
            const r = routers.value[cfg.dataPointIndex];
            if (r) selectedRouter.value = r.name;
          }
        },
      },
    },
    colors: [series.teal],
    plotOptions: {
      bar: {
        horizontal: true,
        borderRadius: 4,
        borderRadiusApplication: "end",
        barHeight: "62%",
        distributed: false,
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
    stroke: {
      show: true,
      width: routers.value.map((_, i) => (i === selectedIdx ? 2 : 0)),
      colors: ["#2ce8d8"],
    },
    states: {
      hover: { filter: { type: "lighten", value: 0.08 } },
      active: { filter: { type: "none", value: 0 } },
    },
    xaxis: {
      ...axisCommon,
      categories: routers.value.map((r) => r.name),
      labels: {
        ...axisCommon.labels,
        formatter: (v: string) => fmtRps(Number(v)),
      },
    },
    yaxis: { ...axisCommon },
    dataLabels: {
      enabled: true,
      textAnchor: "end",
      offsetX: -8,
      offsetY: 4,
      style: {
        colors: ["#e8f4f3"],
        fontSize: "1.1rem",
        fontFamily: "Geist Mono, monospace",
        fontWeight: 600,
      },
      formatter: (v: number) => fmtRps(Number(v)),
    },
    tooltip: {
      enabled: true,
      custom: ({ series: s, seriesIndex, dataPointIndex, w }) => {
        const val = s[seriesIndex][dataPointIndex] as number;
        const name = w.globals.labels[dataPointIndex] as string;
        return `
          <div class="tr-tip">
            <div class="tr-tip__head">${name}</div>
            <div class="tr-tip__body">
              <span class="tr-tip__dot"></span>
              <span class="tr-tip__label">req/s</span>
              <span class="tr-tip__value">${fmtRps(val)}</span>
            </div>
            <div class="tr-tip__foot">clic pour voir la distribution</div>
          </div>
        `;
      },
    },
  };
});
</script>

<template>
  <Panel
    title="Top routers"
    subtitle="Clic sur une barre pour voir la distribution HTTP"
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

<style lang="scss">
// Tooltip custom global (Apex teleporte hors du composant scoped)
.tr-tip {
  background: $bg-surface;
  border: 1px solid $border-default;
  border-radius: $radius-md;
  box-shadow: $shadow-lg;
  overflow: hidden;
  min-width: 18rem;
  font-family: $font-mono;

  &__head {
    padding: 0.5rem 1rem;
    background: $bg-inset;
    border-bottom: 1px solid $border-subtle;
    font-size: 1.1rem;
    color: $text-primary;
    font-weight: 500;
    letter-spacing: -0.01em;
    text-transform: none;
    white-space: nowrap;
    overflow: hidden;
    text-overflow: ellipsis;
  }

  &__body {
    display: grid;
    grid-template-columns: auto 1fr auto;
    align-items: center;
    gap: 0.7rem;
    padding: 0.6rem 1rem;
    font-variant-numeric: tabular-nums;
  }

  &__dot {
    width: 7px;
    height: 7px;
    border-radius: 2px;
    background: $accent;
  }

  &__label {
    color: $text-muted;
    font-size: 1.05rem;
    text-transform: uppercase;
    letter-spacing: 0.08em;
  }

  &__value {
    color: $text-primary;
    font-size: 1.3rem;
    font-weight: 600;
  }

  &__foot {
    padding: 0.45rem 1rem;
    border-top: 1px solid $border-subtle;
    font-family: $font-sans;
    font-size: 1.05rem;
    color: $text-faint;
    text-transform: none;
    letter-spacing: 0;
  }
}
</style>
