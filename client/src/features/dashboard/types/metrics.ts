export interface Kpis {
  rps: number;
  rpsDeltaPct: number;
  latencyP95Ms: number;
  latencyP95DeltaPct: number;
  connections: number;
  connectionsDeltaPct: number;
  errorRatePct: number;
  errorRateDeltaPct: number;
}

export interface RequestsSeries {
  timestamps: number[]; // unix seconds
  series: Record<string, number[]>; // "2xx","3xx","4xx","5xx"
}

export interface LatencySeries {
  timestamps: number[];
  p50: number[];
  p95: number[];
  p99: number[];
}

export interface StatusCodes {
  c2xx: number;
  c3xx: number;
  c4xx: number;
  c5xx: number;
  timeouts: number;
}

export interface RouterRate {
  name: string;
  rps: number;
  codes?: Record<string, number>;
}

export interface MetricsOverview {
  at: number;
  empty: boolean;
  kpis: Kpis;
  requests: RequestsSeries;
  latency: LatencySeries;
  statusCodes: StatusCodes;
  routers: RouterRate[];
}
