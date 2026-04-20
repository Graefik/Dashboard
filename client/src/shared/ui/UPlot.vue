<script setup lang="ts">
import uPlot from "uplot";
import "uplot/dist/uPlot.min.css";
import { onBeforeUnmount, onMounted, ref, watch } from "vue";
import {
  formatTime,
  getBands,
  stackData,
  UPLOT_CSS,
  type UPlotSeries,
} from "@/shared/ui/uplot.helpers";

export type { UPlotSeries };

interface Props {
  data: number[][];
  series: UPlotSeries[];
  height?: number;
  time?: boolean;
  stacked?: boolean;
  yFormatter?: (v: number) => string;
  xFormatter?: (v: number) => string;
}

const props = withDefaults(defineProps<Props>(), {
  height: 320,
  time: true,
  stacked: false,
});

const root = ref<HTMLDivElement>();
const host = ref<HTMLDivElement>();
const tooltip = ref<HTMLDivElement>();
let chart: uPlot | null = null;
let resizeObs: ResizeObserver | null = null;

const buildSeries = (): uPlot.Series[] => [
  { label: "x" },
  ...props.series.map<uPlot.Series>((s) => ({
    label: s.label,
    stroke: s.color,
    width: s.width ?? 2,
    fill:
      s.fill === null
        ? undefined
        : s.fill ??
          (((_u: any) => {
            const ctx = (chart as any)?.ctx as CanvasRenderingContext2D | undefined;
            if (!ctx) return s.color + "22";
            const g = ctx.createLinearGradient(0, 0, 0, props.height);
            g.addColorStop(0, s.color + "4d");
            g.addColorStop(1, s.color + "00");
            return g;
          }) as any),
    dash: s.dash,
    points: { show: false },
    paths: (uPlot.paths as any)?.spline?.({}),
  })),
];

const renderTooltip = (u: uPlot) => {
  if (!tooltip.value || !root.value) return;
  const { left, idx } = u.cursor;
  if (idx == null || left == null || left < 0) {
    tooltip.value.style.opacity = "0";
    return;
  }
  const xVal = u.data[0][idx];
  const rows = props.series
    .map((s, i) => {
      const raw = (props.data[i + 1] ?? [])[idx];
      if (raw == null) return "";
      const v = props.yFormatter ? props.yFormatter(raw) : raw.toLocaleString("fr-FR");
      return `<div class="uplot-tt__row">
        <span class="uplot-tt__dot" style="background:${s.color}"></span>
        <span class="uplot-tt__lbl">${s.label}</span>
        <span class="uplot-tt__val">${v}</span>
      </div>`;
    })
    .join("");
  const head = props.time
    ? formatTime(xVal as number)
    : props.xFormatter
      ? props.xFormatter(xVal as number)
      : String(xVal);
  tooltip.value.innerHTML = `<div class="uplot-tt__head">${head}</div>${rows}`;
  tooltip.value.style.opacity = "1";
  const ttW = tooltip.value.offsetWidth;
  const rootW = root.value.clientWidth;
  let x = left + 14;
  if (x + ttW > rootW) x = left - ttW - 14;
  tooltip.value.style.transform = `translate(${x}px, 8px)`;
};

const buildOptions = (): uPlot.Options => ({
  width: host.value!.clientWidth,
  height: props.height,
  padding: [10, 12, 0, 6],
  series: buildSeries(),
  bands: getBands(props.series.length, props.stacked),
  cursor: {
    drag: { x: false, y: false },
    points: {
      size: 8,
      width: 2,
      stroke: (_u: any, sIdx: number) =>
        (props.series[sIdx - 1]?.color ?? "#00c2b2") + "ff",
      fill: "#0a2236",
    },
    x: true,
    y: false,
  },
  scales: {
    x: { time: props.time },
    y: { auto: true },
  },
  axes: [
    {
      stroke: UPLOT_CSS.axis,
      grid: { stroke: UPLOT_CSS.grid, width: 1, dash: [3, 4] },
      ticks: { show: false },
      font: UPLOT_CSS.font,
      size: 30,
      values: props.xFormatter
        ? (_u, ticks) => ticks.map((t) => props.xFormatter!(t))
        : undefined,
    },
    {
      stroke: UPLOT_CSS.axis,
      grid: { stroke: UPLOT_CSS.grid, width: 1, dash: [3, 4] },
      ticks: { show: false },
      font: UPLOT_CSS.font,
      size: 44,
      values: props.yFormatter
        ? (_u, ticks) => ticks.map((t) => props.yFormatter!(t))
        : (_u, ticks) => ticks.map((t) => t.toLocaleString("fr-FR")),
    },
  ],
  legend: { show: false },
  hooks: {
    setCursor: [renderTooltip],
    setSize: [
      () => {
        if (tooltip.value) tooltip.value.style.opacity = "0";
      },
    ],
  },
});

const build = () => {
  if (!host.value) return;
  const data = stackData(props.data, props.stacked) as uPlot.AlignedData;
  chart = new uPlot(buildOptions(), data, host.value);
};

const destroy = () => {
  chart?.destroy();
  chart = null;
};

onMounted(() => {
  build();
  resizeObs = new ResizeObserver(() => {
    if (!chart || !host.value) return;
    chart.setSize({ width: host.value.clientWidth, height: props.height });
  });
  resizeObs.observe(host.value!);
});

onBeforeUnmount(() => {
  resizeObs?.disconnect();
  destroy();
});

watch(
  () => [props.data, props.series, props.stacked],
  () => {
    destroy();
    build();
  },
  { deep: true },
);
</script>

<template>
  <div ref="root" class="uplot-wrap">
    <div ref="host" class="uplot-host"></div>
    <div ref="tooltip" class="uplot-tt" aria-hidden="true"></div>
  </div>
</template>

<style lang="scss">
.uplot-wrap {
  position: relative;
  width: 100%;
}

.uplot-host {
  width: 100%;

  .u-over,
  .u-under {
    background: transparent !important;
  }

  .u-legend {
    display: none !important;
  }

  .u-cursor-x,
  .u-cursor-y {
    background: rgba(0, 194, 178, 0.35) !important;
  }
}

.uplot-tt {
  position: absolute;
  top: 0;
  left: 0;
  pointer-events: none;
  opacity: 0;
  min-width: 15rem;
  padding: 0;
  background: $bg-surface;
  border: 1px solid $border-default;
  border-radius: $radius-md;
  box-shadow: $shadow-lg;
  font-family: $font-mono;
  font-size: 1.15rem;
  color: $text-primary;
  transition: opacity 0.1s $ease-out;
  z-index: 10;
  overflow: hidden;

  &__head {
    padding: 0.5rem 1rem;
    background: $bg-inset;
    border-bottom: 1px solid $border-subtle;
    font-size: 1.05rem;
    color: $text-muted;
    text-transform: uppercase;
    letter-spacing: 0.08em;
  }

  &__row {
    display: grid;
    grid-template-columns: auto 1fr auto;
    align-items: center;
    gap: 0.7rem;
    padding: 0.4rem 1rem;
    font-variant-numeric: tabular-nums;

    & + & {
      border-top: 1px solid $border-subtle;
    }
  }

  &__dot {
    width: 7px;
    height: 7px;
    border-radius: 2px;
  }

  &__lbl {
    color: $text-secondary;
  }

  &__val {
    color: $text-primary;
    font-weight: 600;
  }
}
</style>
