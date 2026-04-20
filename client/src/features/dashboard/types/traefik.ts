export interface Counts {
  total: number;
  warnings: number;
  errors: number;
}

export interface Section {
  routers: Counts;
  services: Counts;
  middlewares: Counts;
}

export interface Features {
  tracing: string;
  metrics: string;
  accessLog: boolean;
  hub: boolean;
}

export interface Overview {
  http: Section;
  tcp: Section;
  udp: Section;
  features: Features;
  version: string;
}

export interface Router {
  name: string;
  rule: string;
  entryPoints: string[];
  service: string;
  middlewares?: string[];
  status: string;
  using?: string[];
  provider: string;
  tls?: {
    certResolver: string;
    domains?: { main: string; sans?: string[] }[];
  };
}

export interface Service {
  name: string;
  type: string;
  provider: string;
  status: string;
  usedBy?: string[];
  serverStatus?: Record<string, string>;
}
