// Client HTTP partagé. Fetch wrapper minimaliste avec JSON + erreurs typées.

export interface ApiError {
  status: number;
  code?: string;
  message: string;
}

const BASE_URL = (import.meta.env.VITE_API_URL as string | undefined) ?? "";

const TOKEN_KEY = "graefik.token";

export const getToken = (): string | null => localStorage.getItem(TOKEN_KEY);
export const setToken = (t: string) => localStorage.setItem(TOKEN_KEY, t);
export const clearToken = () => localStorage.removeItem(TOKEN_KEY);

interface RequestOptions {
  method?: "GET" | "POST" | "PUT" | "PATCH" | "DELETE";
  body?: unknown;
  auth?: boolean;
  signal?: AbortSignal;
}

export async function request<T>(
  path: string,
  opts: RequestOptions = {},
): Promise<T> {
  const headers: Record<string, string> = {
    Accept: "application/json",
  };

  if (opts.body !== undefined) {
    headers["Content-Type"] = "application/json";
  }

  if (opts.auth) {
    const token = getToken();
    if (token) headers["Authorization"] = `Bearer ${token}`;
  }

  let res: Response;
  try {
    res = await fetch(`${BASE_URL}${path}`, {
      method: opts.method ?? "GET",
      headers,
      body: opts.body !== undefined ? JSON.stringify(opts.body) : undefined,
      signal: opts.signal,
    });
  } catch (e) {
    throw {
      status: 0,
      code: "NETWORK_ERROR",
      message: "Impossible de joindre le serveur.",
    } satisfies ApiError;
  }

  // 204 No Content
  if (res.status === 204) return undefined as T;

  const isJSON = res.headers.get("content-type")?.includes("application/json");
  const payload = isJSON ? await res.json().catch(() => null) : null;

  if (!res.ok) {
    throw {
      status: res.status,
      code: payload?.code,
      message: payload?.error ?? `Erreur ${res.status}`,
    } satisfies ApiError;
  }

  return payload as T;
}
