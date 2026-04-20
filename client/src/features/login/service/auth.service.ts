import { request } from "@/shared/api/http";
import type { LoginPayload, LoginResponse } from "@/features/login/types/auth";

export const authService = {
  login(payload: LoginPayload): Promise<LoginResponse> {
    return request<LoginResponse>("/api/auth/login", {
      method: "POST",
      body: payload,
    });
  },
};
