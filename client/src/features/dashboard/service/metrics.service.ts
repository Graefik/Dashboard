import { request } from "@/shared/api/http";
import type { MetricsOverview } from "@/features/dashboard/types/metrics";

export const metricsService = {
  getOverview(): Promise<MetricsOverview> {
    return request<MetricsOverview>("/api/metrics/overview", { auth: true });
  },
};
