<script setup lang="ts">
import { computed } from "vue";
import Panel from "@/shared/ui/Panel.vue";
import PanelChips from "@/shared/ui/PanelChips.vue";
import UPlot, { type UPlotSeries } from "@/shared/ui/UPlot.vue";
import { series } from "@/features/dashboard/theme";
import { useMetrics } from "@/features/dashboard/composables/useMetrics";

const { overview } = useMetrics();

// ── Séries candidates (ordre d'affichage + couleurs) ────────────────────
const allSeries = [
  {
    label: "2xx/3xx",
    color: series.teal,
    pick: (s: Record<string, number[]>) =>
      (s["2xx"] ?? []).map((v, i) => v + ((s["3xx"] ?? [])[i] ?? 0)),
  },
  {
    label: "4xx",
    color: series.amber,
    pick: (s: Record<string, number[]>) => s["4xx"] ?? [],
  },
  {
    label: "5xx",
    color: series.red,
    pick: (s: Record<string, number[]>) => s["5xx"] ?? [],
  },
];

// Ne garde que les séries avec AU MOINS une valeur > 0. Évite que uPlot
// dessine un stroke "fantôme" au sommet du stack (la ligne rouge qui
// monte alors que 5xx = 0 partout).
const active = computed(() => {
  const req = overview.value?.requests;
  if (!req || !req.timestamps || req.timestamps.length < 2) return [];
  return allSeries
    .map((s) => ({ ...s, data: s.pick(req.series ?? {}) }))
    .filter((s) => s.data.some((v) => v > 0));
});

const data = computed<number[][]>(() => {
  const xs = overview.value?.requests?.timestamps ?? [];
  if (!active.value.length || !xs.length) return [[], []];
  return [xs, ...active.value.map((s) => s.data)];
});

// Seule la série du bas (index 0) a un stroke. Les supérieures sont
// uniquement des fills (via bands) : ça évite que le stroke de 4xx/5xx
// ne soit dessiné sur le cumul et donne l'illusion d'une "ligne qui
// monte" alors que la série réelle est à 0 sur la période.
const uplotSeries = computed<UPlotSeries[]>(() =>
  active.value.map((s, idx) => ({
    label: s.label,
    color: s.color,
    width: idx === 0 ? 2 : 0,
  })),
);

// Chips : toujours les 3 familles (légende de référence), avec un flag
// `muted` quand la série n'a pas de données sur la période — ça indique
// visuellement qu'elle existe en théorie mais n'est pas active.
const chips = computed(() => {
  const activeLabels = new Set(active.value.map((s) => s.label));
  return allSeries.map((s) => ({
    label: s.label,
    color: s.color,
    muted: !activeLabels.has(s.label),
  }));
});
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
      :key="(data[0]?.length ?? 0) + '-' + uplotSeries.length"
      :data="data"
      :series="uplotSeries"
      :height="300"
      stacked
      :y-formatter="(v: number) => v.toFixed(1) + ' r/s'"
    />
  </Panel>
</template>
