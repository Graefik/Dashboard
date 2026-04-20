// Mock helpers pour générer des séries réalistes (à remplacer par l'API Traefik)

export const spark = (base: number, variance: number, n = 24): number[] =>
  Array.from({ length: n }, (_, i) =>
    Math.max(
      0,
      Math.round(
        base + Math.sin(i / 2.2) * variance + (Math.random() - 0.5) * variance * 0.6,
      ),
    ),
  );

export const genTs = (n: number, stepMs: number): number[] => {
  const now = Math.floor(Date.now() / 1000);
  return Array.from({ length: n }, (_, i) =>
    now - Math.floor(((n - 1 - i) * stepMs) / 1000),
  );
};

export const genY = (n: number, base: number, variance: number): number[] =>
  Array.from({ length: n }, (_, i) =>
    Math.max(
      0,
      Math.round(
        base +
          Math.sin(i / 4) * variance * 0.6 +
          Math.cos(i / 9) * variance * 0.3 +
          (Math.random() - 0.5) * variance * 0.4,
      ),
    ),
  );
