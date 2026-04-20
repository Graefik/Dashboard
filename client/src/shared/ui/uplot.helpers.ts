import uPlot from "uplot";

export interface UPlotSeries {
  label: string;
  color: string;
  width?: number;
  fill?: string | null;
  dash?: number[];
}

export const UPLOT_CSS = {
  grid: "rgba(255,255,255,0.05)",
  axis: "#5a7d8e",
  font: "11px 'Geist Mono', 'JetBrains Mono', monospace",
} as const;

// Transforme des séries brutes en séries cumulatives (pour stacked area)
export const stackData = (raw: number[][], stacked: boolean): number[][] => {
  if (!stacked || raw.length < 3) return raw;
  const out: number[][] = [raw[0]];
  const acc = new Array(raw[0].length).fill(0);
  for (let i = 1; i < raw.length; i++) {
    const serie: number[] = [];
    for (let j = 0; j < raw[i].length; j++) {
      acc[j] += raw[i][j] ?? 0;
      serie.push(acc[j]);
    }
    out.push(serie);
  }
  return out;
};

// Bandes entre séries successives pour le fill area empilé
export const getBands = (count: number, stacked: boolean): uPlot.Band[] => {
  if (!stacked) return [];
  const bands: uPlot.Band[] = [];
  for (let i = count; i > 1; i--) {
    bands.push({ series: [i, i - 1] });
  }
  return bands;
};

// Formateur d'heure pour le header du tooltip
export const formatTime = (sec: number): string =>
  new Date(sec * 1000).toLocaleTimeString("fr-FR", {
    hour: "2-digit",
    minute: "2-digit",
    second: "2-digit",
  });
