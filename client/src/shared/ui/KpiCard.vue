<script setup lang="ts">
import { computed } from "vue";
import type { ApexOptions } from "apexcharts";

interface Props {
  label: string;
  value: string | number;
  unit?: string;
  delta?: number;
  deltaLabel?: string;
  trend?: number[];
  tone?: "accent" | "success" | "warning" | "danger" | "info";
  icon?: "requests" | "latency" | "connections" | "errors" | "throughput";
  hero?: boolean;
}

const props = withDefaults(defineProps<Props>(), {
  tone: "accent",
  delta: 0,
  hero: false,
});

const toneColor: Record<NonNullable<Props["tone"]>, string> = {
  accent: "#00c2b2",
  success: "#2dd4a7",
  warning: "#f5a524",
  danger: "#f16969",
  info: "#58a6ff",
};

const deltaPositive = computed(() => (props.delta ?? 0) >= 0);
const isErrorMetric = computed(() => props.tone === "danger" || props.tone === "warning");
const deltaGood = computed(() =>
  isErrorMetric.value ? !deltaPositive.value : deltaPositive.value,
);

const sparkSeries = computed(() => [
  {
    name: props.label,
    data: props.trend ?? [],
  },
]);

const sparkOptions = computed<ApexOptions>(() => {
  const c = toneColor[props.tone];
  return {
    chart: {
      type: "area",
      sparkline: { enabled: true },
      animations: { enabled: true, speed: 400 },
      background: "transparent",
    },
    colors: [c],
    stroke: {
      curve: "smooth",
      width: 2,
      lineCap: "round",
    },
    fill: {
      type: "gradient",
      gradient: {
        shadeIntensity: 1,
        opacityFrom: 0.35,
        opacityTo: 0,
        stops: [0, 100],
      },
    },
    tooltip: { enabled: false },
    xaxis: { crosshairs: { show: false } },
  };
});
</script>

<template>
  <article class="kpi" :class="[`kpi--${tone}`, { 'kpi--hero': hero }]">
    <div class="kpi__head">
      <div class="kpi__icon" aria-hidden="true">
        <svg v-if="icon === 'requests'" viewBox="0 0 20 20" fill="none">
          <path
            d="M3 10h4l2-5 2 10 2-5h4"
            stroke="currentColor"
            stroke-width="1.6"
            stroke-linecap="round"
            stroke-linejoin="round"
          />
        </svg>
        <svg v-else-if="icon === 'latency'" viewBox="0 0 20 20" fill="none">
          <circle cx="10" cy="10" r="7" stroke="currentColor" stroke-width="1.5" />
          <path d="M10 6v4l2.5 2.5" stroke="currentColor" stroke-width="1.5" stroke-linecap="round" />
        </svg>
        <svg v-else-if="icon === 'connections'" viewBox="0 0 20 20" fill="none">
          <circle cx="5" cy="10" r="2.2" stroke="currentColor" stroke-width="1.5" />
          <circle cx="15" cy="10" r="2.2" stroke="currentColor" stroke-width="1.5" />
          <path d="M7.2 10h5.6" stroke="currentColor" stroke-width="1.5" />
        </svg>
        <svg v-else-if="icon === 'errors'" viewBox="0 0 20 20" fill="none">
          <path
            d="M10 2.5 2 16.5h16L10 2.5z"
            stroke="currentColor"
            stroke-width="1.5"
            stroke-linejoin="round"
          />
          <path d="M10 8v3.5" stroke="currentColor" stroke-width="1.5" stroke-linecap="round" />
          <circle cx="10" cy="13.5" r="0.8" fill="currentColor" />
        </svg>
        <svg v-else-if="icon === 'throughput'" viewBox="0 0 20 20" fill="none">
          <path
            d="M3 14l4-4 3 3 4-6 3 4"
            stroke="currentColor"
            stroke-width="1.6"
            stroke-linecap="round"
            stroke-linejoin="round"
          />
        </svg>
      </div>
      <span class="kpi__label">{{ label }}</span>
    </div>

    <div class="kpi__body">
      <div class="kpi__value-row">
        <span class="kpi__value mono">{{ value }}</span>
        <span v-if="unit" class="kpi__unit mono">{{ unit }}</span>
      </div>
      <div
        v-if="delta !== undefined"
        class="kpi__delta"
        :class="{ 'kpi__delta--good': deltaGood, 'kpi__delta--bad': !deltaGood }"
      >
        <svg viewBox="0 0 10 10" width="10" height="10" aria-hidden="true">
          <path
            :d="deltaPositive ? 'M2 7l3-4 3 4' : 'M2 3l3 4 3-4'"
            stroke="currentColor"
            stroke-width="1.6"
            fill="none"
            stroke-linecap="round"
            stroke-linejoin="round"
          />
        </svg>
        <span class="mono">{{ deltaPositive ? "+" : "" }}{{ delta }}%</span>
        <span v-if="deltaLabel" class="kpi__delta-label">{{ deltaLabel }}</span>
      </div>
    </div>

    <div v-if="trend && trend.length" class="kpi__spark">
      <apexchart
        type="area"
        height="56"
        width="100%"
        :options="sparkOptions"
        :series="sparkSeries"
      />
    </div>
  </article>
</template>

<style lang="scss" scoped>
.kpi {
  --tone: #{$accent};
  position: relative;
  display: flex;
  flex-direction: column;
  gap: 1rem;
  padding: 1.6rem 1.8rem 0.8rem;
  background: linear-gradient(
    180deg,
    $bg-surface 0%,
    $bg-surface-2 100%
  );
  border: 1px solid $border-subtle;
  border-radius: $radius-lg;
  overflow: hidden;
  transition:
    border-color 0.2s $ease-out,
    transform 0.2s $ease-out;

  &::before {
    content: "";
    position: absolute;
    top: 0;
    left: 0;
    bottom: 0;
    width: 2px;
    background: var(--tone);
    opacity: 0;
    transition: opacity 0.2s $ease-out;
  }

  &--hero {
    border-color: $border-accent;

    &::before {
      opacity: 0.9;
    }
  }

  &:hover {
    border-color: rgba(255, 255, 255, 0.12);
    transform: translateY(-1px);
  }

  &--hero:hover {
    border-color: $border-accent;
  }

  &--accent {
    --tone: #{$accent};
  }
  &--success {
    --tone: #{$severity-success};
  }
  &--warning {
    --tone: #{$severity-warning};
  }
  &--danger {
    --tone: #{$severity-danger};
  }
  &--info {
    --tone: #{$severity-info};
  }

  &__head {
    display: flex;
    align-items: center;
    gap: 0.8rem;
    color: $text-muted;
  }

  &__icon {
    display: inline-flex;
    align-items: center;
    justify-content: center;
    width: 2.4rem;
    height: 2.4rem;
    background: rgba(255, 255, 255, 0.03);
    border: 1px solid $border-subtle;
    border-radius: $radius-sm;
    color: var(--tone);

    svg {
      width: 1.4rem;
      height: 1.4rem;
    }
  }

  &__label {
    font-family: $font-mono;
    font-size: 1.1rem;
    text-transform: uppercase;
    letter-spacing: 0.1em;
    color: $text-muted;
    font-weight: 500;
  }

  &__body {
    display: flex;
    align-items: flex-end;
    justify-content: space-between;
    gap: 1rem;
    position: relative;
    z-index: 1;
  }

  &__value-row {
    display: inline-flex;
    align-items: baseline;
    gap: 0.4rem;
    min-width: 0;
  }

  &__value {
    font-size: 2.8rem;
    font-weight: 600;
    color: $text-primary;
    line-height: 1;
    letter-spacing: -0.02em;
    font-variant-numeric: tabular-nums;
  }

  &__unit {
    font-size: 1.3rem;
    font-weight: 500;
    color: $text-muted;
  }

  &__delta {
    display: inline-flex;
    align-items: center;
    gap: 0.4rem;
    padding: 0.25rem 0.6rem;
    border-radius: $radius-sm;
    font-size: 1.15rem;
    font-weight: 600;
    white-space: nowrap;

    &--good {
      color: $severity-success;
      background: $severity-success-soft;
    }

    &--bad {
      color: $severity-danger;
      background: $severity-danger-soft;
    }
  }

  &__delta-label {
    color: $text-muted;
    font-weight: 500;
    font-size: 1rem;
    margin-left: 0.3rem;
  }

  &__spark {
    margin: 0 -1.8rem -0.4rem;
    pointer-events: none;
  }
}
</style>
