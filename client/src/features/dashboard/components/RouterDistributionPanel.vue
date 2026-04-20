<script setup lang="ts">
import type { ApexOptions } from "apexcharts";
import { computed, onBeforeUnmount, onMounted, ref } from "vue";
import {
  DropdownMenuContent,
  DropdownMenuItem,
  DropdownMenuLabel,
  DropdownMenuPortal,
  DropdownMenuRoot,
  DropdownMenuSeparator,
  DropdownMenuTrigger,
} from "reka-ui";
import Panel from "@/shared/ui/Panel.vue";
import { baseChart, series } from "@/features/dashboard/theme";
import { useMetrics } from "@/features/dashboard/composables/useMetrics";

const { overview, selectedRouter, selectedRouterData } = useMetrics();

// ── Taille responsive du donut : suit la taille du container ─────────────
const chartEl = ref<HTMLDivElement>();
const chartSize = ref(320);
let ro: ResizeObserver | null = null;

onMounted(() => {
  if (!chartEl.value) return;
  ro = new ResizeObserver(([entry]) => {
    const { width, height } = entry.contentRect;
    // donut carré, on prend le min des deux axes (clampé)
    chartSize.value = Math.max(200, Math.min(width, height));
  });
  ro.observe(chartEl.value);
});

onBeforeUnmount(() => {
  ro?.disconnect();
  ro = null;
});

// Tous les routers ayant eu du trafic (liste complète pour le dropdown)
const routers = computed(() => overview.value?.routers ?? []);

const fmtRps = (v: number): string => {
  if (v === 0) return "0";
  if (v < 1) return v.toFixed(2);
  return v.toFixed(1);
};

// Agrège les codes individuels en 4 familles + tri par code
interface FamilyBreakdown {
  c2xx: number;
  c3xx: number;
  c4xx: number;
  c5xx: number;
}

const families = computed<FamilyBreakdown>(() => {
  const out = { c2xx: 0, c3xx: 0, c4xx: 0, c5xx: 0 };
  const codes = selectedRouterData.value?.codes ?? {};
  for (const [code, count] of Object.entries(codes)) {
    if (code.startsWith("2")) out.c2xx += count;
    else if (code.startsWith("3")) out.c3xx += count;
    else if (code.startsWith("4")) out.c4xx += count;
    else if (code.startsWith("5")) out.c5xx += count;
  }
  return out;
});

const total = computed(
  () => families.value.c2xx + families.value.c3xx + families.value.c4xx + families.value.c5xx,
);

const sortedCodes = computed(() => {
  const codes = selectedRouterData.value?.codes ?? {};
  return Object.entries(codes)
    .map(([code, count]) => ({ code, count }))
    .sort((a, b) => b.count - a.count)
    .slice(0, 8);
});

const chartSeries = computed(() => [
  families.value.c2xx,
  families.value.c3xx,
  families.value.c4xx,
  families.value.c5xx,
]);

const options = computed<ApexOptions>(() => ({
  ...baseChart,
  chart: {
    ...baseChart.chart,
    type: "donut",
    height: 240,
    width: 240,
    // élimine le padding parasite de l'SVG Apex
    offsetX: 0,
    offsetY: 0,
    parentHeightOffset: 0,
    sparkline: { enabled: false },
  },
  labels: ["2xx OK", "3xx Redirect", "4xx Client", "5xx Server"],
  colors: [series.teal, series.tealDeep, series.amber, series.red],
  stroke: { width: 2, colors: ["#0a2236"] },
  plotOptions: {
    pie: {
      expandOnClick: false,
      offsetX: 0,
      offsetY: 0,
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
            showAlways: true,
            label: "Total",
            color: "#5a7d8e",
            fontSize: "1.1rem",
            fontFamily: "Geist Mono, monospace",
            formatter: () => total.value.toLocaleString("fr-FR"),
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

const codeTone = (code: string): string => {
  if (code.startsWith("2")) return "ok";
  if (code.startsWith("3")) return "ok";
  if (code.startsWith("4")) return "warn";
  if (code.startsWith("5")) return "err";
  return "mute";
};
</script>

<template>
  <Panel
    :title="selectedRouterData ? selectedRouterData.name : 'Router distribution'"
    subtitle="Répartition HTTP cumulée depuis le dernier reload Traefik"
    eyebrow="Distribution"
    :span="7"
    pad="tight"
  >
    <template #actions>
      <DropdownMenuRoot>
        <DropdownMenuTrigger as-child>
          <button
            class="router-select"
            type="button"
            :disabled="!routers.length"
            aria-label="Changer de router"
          >
            <span class="router-select__label">Router</span>
            <span class="router-select__name mono">
              {{ selectedRouter ?? "—" }}
            </span>
            <svg
              class="router-select__chevron"
              width="12"
              height="12"
              viewBox="0 0 12 12"
              fill="none"
              aria-hidden="true"
            >
              <path
                d="M3 4.5l3 3 3-3"
                stroke="currentColor"
                stroke-width="1.4"
                stroke-linecap="round"
                stroke-linejoin="round"
              />
            </svg>
          </button>
        </DropdownMenuTrigger>
        <DropdownMenuPortal>
          <DropdownMenuContent
            class="dd dd--router"
            :side-offset="8"
            align="end"
          >
            <DropdownMenuLabel class="dd__label dd__label--router">
              Routers actifs ({{ routers.length }})
            </DropdownMenuLabel>
            <DropdownMenuSeparator class="dd__sep" />
            <DropdownMenuItem
              v-for="r in routers"
              :key="r.name"
              class="dd__item dd__item--router"
              :data-active="r.name === selectedRouter ? 'true' : undefined"
              @select="selectedRouter = r.name"
            >
              <span class="dd__router-name mono">{{ r.name }}</span>
              <span class="dd__router-rps mono">{{ fmtRps(r.rps) }} r/s</span>
              <svg
                v-if="r.name === selectedRouter"
                class="dd__router-check"
                width="14"
                height="14"
                viewBox="0 0 14 14"
                fill="none"
                aria-hidden="true"
              >
                <path
                  d="m3 7.5 2.5 2.5L11 4.5"
                  stroke="currentColor"
                  stroke-width="1.8"
                  stroke-linecap="round"
                  stroke-linejoin="round"
                />
              </svg>
            </DropdownMenuItem>
          </DropdownMenuContent>
        </DropdownMenuPortal>
      </DropdownMenuRoot>
    </template>

    <div v-if="!selectedRouterData || total === 0" class="dist__empty">
      <svg width="32" height="32" viewBox="0 0 32 32" fill="none" aria-hidden="true">
        <circle cx="16" cy="16" r="11" stroke="currentColor" stroke-width="1.4" />
        <path d="M16 11v5l3 3" stroke="currentColor" stroke-width="1.4" stroke-linecap="round" />
      </svg>
      <p>
        {{
          selectedRouterData
            ? "Aucune requête enregistrée sur ce router."
            : "Sélectionne un router dans le Top à gauche."
        }}
      </p>
    </div>

    <div v-else class="dist">
      <div ref="chartEl" class="dist__chart">
        <apexchart
          type="donut"
          :height="chartSize"
          :width="chartSize"
          :options="options"
          :series="chartSeries"
        />
      </div>

      <div class="dist__codes">
        <div class="dist__codes-title">Détail codes</div>
        <div class="dist__rows">
          <div
            v-for="row in sortedCodes"
            :key="row.code"
            class="dist__row"
            :class="`dist__row--${codeTone(row.code)}`"
          >
            <span class="dist__code mono">{{ row.code }}</span>
            <span class="dist__bar">
              <span
                class="dist__bar-fill"
                :style="{ width: `${(row.count / total) * 100}%` }"
              ></span>
            </span>
            <span class="dist__count mono">
              {{ row.count.toLocaleString("fr-FR") }}
            </span>
            <span class="dist__pct mono">
              {{ ((row.count / total) * 100).toFixed(1) }}%
            </span>
          </div>
        </div>
      </div>
    </div>
  </Panel>
</template>

<style scoped lang="scss">
// ── Select (trigger) du router actuel, dans les actions du Panel ──
.router-select {
  display: inline-flex;
  align-items: center;
  gap: 0.7rem;
  height: 3.2rem;
  padding: 0 1rem;
  background: $bg-inset;
  border: 1px solid $border-subtle;
  border-radius: $radius-md;
  color: $text-primary;
  cursor: pointer;
  font-family: $font-sans;
  transition:
    border-color 0.15s $ease-out,
    background 0.15s $ease-out,
    color 0.15s $ease-out;

  &:hover:not(:disabled) {
    border-color: $border-accent;
    background: $bg-elevated;
  }

  &:disabled {
    opacity: 0.55;
    cursor: not-allowed;
  }

  &[data-state="open"] {
    border-color: $border-accent;
    background: $bg-elevated;
  }

  &__label {
    font-family: $font-mono;
    font-size: 0.95rem;
    text-transform: uppercase;
    letter-spacing: 0.12em;
    color: $text-muted;
  }

  &__name {
    max-width: 22rem;
    overflow: hidden;
    text-overflow: ellipsis;
    white-space: nowrap;
    font-size: 1.2rem;
    color: $text-primary;
  }

  &__chevron {
    color: $text-muted;
    flex-shrink: 0;
  }
}

.dist {
  display: grid;
  grid-template-columns: minmax(0, 1fr) minmax(0, 1fr);
  gap: 2.4rem;
  align-items: stretch;
  min-height: 34rem;

  @media (max-width: 1100px) {
    grid-template-columns: 1fr;
    gap: 2rem;
    min-height: 0;
  }

  &__chart {
    position: relative;
    min-width: 0;
    width: 100%;
    height: 100%;
    min-height: 32rem;
    display: flex;
    align-items: center;
    justify-content: center;

    // Le wrapper que vue3-apexcharts génère est un div block avec
    // line-height par défaut, qui ajoute une baseline fantôme.
    :deep(.vue-apexcharts) {
      line-height: 0;
      display: flex;
      align-items: center;
      justify-content: center;
    }

    :deep(.apexcharts-canvas) {
      display: block;
      margin: 0 auto;
    }
  }

  &__codes {
    min-width: 0;
    align-self: center;
    display: flex;
    flex-direction: column;
    gap: 0.6rem;
    padding-right: 0.4rem;
  }

  &__codes-title {
    font-family: $font-mono;
    font-size: 1rem;
    color: $text-muted;
    text-transform: uppercase;
    letter-spacing: 0.12em;
    margin-bottom: 0.4rem;
  }

  &__rows {
    display: flex;
    flex-direction: column;
    gap: 0.3rem;
  }

  &__row {
    display: grid;
    grid-template-columns: 3.2rem 1fr auto auto;
    align-items: center;
    gap: 0.8rem;
    padding: 0.35rem 0;
    font-size: 1.2rem;
  }

  &__code {
    font-weight: 600;
    text-align: right;
  }

  &--ok &__code {
    color: $severity-success;
  }
  &--warn &__code {
    color: $severity-warning;
  }
  &--err &__code {
    color: $severity-danger;
  }

  &__bar {
    position: relative;
    height: 6px;
    background: rgba(255, 255, 255, 0.04);
    border-radius: 2px;
    overflow: hidden;
  }

  &__bar-fill {
    position: absolute;
    inset: 0 auto 0 0;
    border-radius: 2px;
    background: linear-gradient(90deg, $accent, $accent-glow);
  }

  &__row--warn &__bar-fill {
    background: $severity-warning;
  }
  &__row--err &__bar-fill {
    background: $severity-danger;
  }

  &__count {
    color: $text-primary;
    font-variant-numeric: tabular-nums;
    min-width: 5rem;
    text-align: right;
  }

  &__pct {
    color: $text-muted;
    font-variant-numeric: tabular-nums;
    min-width: 4rem;
    text-align: right;
    font-size: 1.1rem;
  }

  &__empty {
    display: flex;
    flex-direction: column;
    align-items: center;
    justify-content: center;
    gap: 0.8rem;
    height: 26rem;
    color: $text-muted;
    text-align: center;

    p {
      font-size: 1.3rem;
      margin: 0;
    }
  }
}
</style>

<style lang="scss">
// Dropdown téléporté par Reka, styles globaux comme pour .dd--env
.dd--router {
  min-width: 32rem;
  max-height: 40rem;
  overflow-y: auto;

  .dd__label--router {
    padding: 0.7rem 1rem 0.5rem;
    font-family: $font-mono;
    font-size: 1rem;
    color: $text-muted;
    text-transform: uppercase;
    letter-spacing: 0.1em;
  }

  .dd__item--router {
    display: grid;
    grid-template-columns: 1fr auto auto;
    align-items: center;
    gap: 1rem;
    padding: 0.6rem 1rem;
    font-family: $font-sans;
    font-size: 1.25rem;
  }

  .dd__item--router[data-active="true"] {
    background: $accent-soft;
    color: $text-primary;
  }

  .dd__router-name {
    color: $text-primary;
    font-size: 1.2rem;
    overflow: hidden;
    text-overflow: ellipsis;
    white-space: nowrap;
  }

  .dd__router-rps {
    color: $text-muted;
    font-size: 1.1rem;
    font-variant-numeric: tabular-nums;
    white-space: nowrap;
  }

  .dd__router-check {
    color: $accent-glow;
    flex-shrink: 0;
  }
}
</style>
