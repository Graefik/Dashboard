import { request } from "@/shared/api/http";
import type {
  Overview,
  Router,
  Service,
} from "@/features/dashboard/types/traefik";

export const observeService = {
  getOverview(): Promise<Overview> {
    return request<Overview>("/api/overview", { auth: true });
  },
  listRouters(): Promise<Router[]> {
    return request<Router[]>("/api/http/routers", { auth: true });
  },
  listServices(): Promise<Service[]> {
    return request<Service[]>("/api/http/services", { auth: true });
  },
};
