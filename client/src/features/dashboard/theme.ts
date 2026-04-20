import type { ApexOptions } from "apexcharts";

// ── Palette : 1 accent teal + échelle neutre, statuts sémantiques ──
export const series = {
  teal: "#00c2b2",
  tealGlow: "#2ce8d8",
  tealDeep: "#0b6f66",
  amber: "#f5a524",
  red: "#f16969",
  green: "#2dd4a7",
  slate: "#94a3b8",
  slateDeep: "#556978",
} as const;

// ── Base Apex (utilisée pour donut + bar horizontal + sparklines KPI) ──
export const baseChart: Partial<ApexOptions> = {
  chart: {
    toolbar: { show: false },
    zoom: { enabled: false },
    background: "transparent",
    fontFamily: "Geist Mono, JetBrains Mono, monospace",
    foreColor: "#5a7d8e",
    animations: {
      enabled: true,
      speed: 420,
      animateGradually: { enabled: true, delay: 120 },
    },
  },
  dataLabels: { enabled: false },
  legend: { show: false },
  grid: {
    borderColor: "rgba(255, 255, 255, 0.05)",
    strokeDashArray: 3,
    padding: { left: 10, right: 10, top: 0, bottom: 0 },
    xaxis: { lines: { show: false } },
  },
  tooltip: {
    theme: "dark",
    style: { fontSize: "1.15rem", fontFamily: "Geist Mono, monospace" },
  },
};

export const axisCommon = {
  labels: {
    style: {
      colors: "#5a7d8e",
      fontSize: "1.05rem",
      fontFamily: "Geist Mono, monospace",
    },
  },
  axisBorder: { show: false },
  axisTicks: { show: false },
};
