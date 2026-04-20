<template>
  <div class="brand">
    <img src="@/assets/icon.png" alt="Graefik" class="brand__logo" />
    <div class="brand__text">
      <span class="brand__name">Graefik</span>
      <span class="brand__meta">v0.1 · on-prem</span>
    </div>
  </div>
</template>

<style scoped lang="scss">
.brand {
  display: flex;
  align-items: center;
  gap: 1.1rem;
  padding: 0 2rem 1.8rem;

  &__logo {
    width: 4rem;
    height: 4rem;
    flex-shrink: 0;
    object-fit: contain;
    border-radius: $radius-sm;
    filter: drop-shadow(0 2px 6px rgba(0, 0, 0, 0.5));
  }

  &__text {
    display: flex;
    flex-direction: column;
    line-height: 1.1;
  }

  &__name {
    font-size: 1.7rem;
    font-weight: 600;
    color: $text-primary;
    letter-spacing: -0.02em;
  }

  &__meta {
    font-family: $font-mono;
    font-size: 1rem;
    color: $text-muted;
    margin-top: 0.3rem;
    letter-spacing: 0.04em;
  }
}
</style>
